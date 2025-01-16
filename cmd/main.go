package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/yumin00/go-hexagonal/core/config"
	pb "github.com/yumin00/go-hexagonal/go-proto/go-api/core/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config.Init()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	grpcServer := grpc.NewServer()
	config.RegisterDataServer(grpcServer)

	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		log.Printf("Starting gRPC server on port %s", cfg.Server.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Create a new context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a new mux for the HTTP server
	gwmux := runtime.NewServeMux()

	// Register the gateway handler
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = pb.RegisterUserDataHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("localhost:%s", cfg.Server.GRPCPort), opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	// Start HTTP server
	log.Printf("Starting HTTP server on port %s", cfg.Server.HTTPPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.HTTPPort), gwmux); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
