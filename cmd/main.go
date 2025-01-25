package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/yumin00/go-hexagonal/core/config"
	pb "github.com/yumin00/go-hexagonal/go-proto/go-api/core/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func main() {
	//config.Setup()
	//ctx := context.Background()
	//
	//grpcServer, gatewayServer := config.Start(ctx)
	//
	//go func() {
	//	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.LoadedConfig.Server.GRPCPort))
	//	if err != nil {
	//		log.Fatalf("Failed to listen: %v", err)
	//	}
	//	log.Printf("Starting gRPC server on port %s", config.LoadedConfig.Server.GRPCPort)
	//	if err := grpcServer.Serve(lis); err != nil {
	//		log.Fatalf("Failed to serve gRPC: %v", err)
	//	}
	//}()
	//
	//go func() {
	//	log.Printf("Starting HTTP server on port %s", config.LoadedConfig.Server.HTTPPort)
	//	if err := gatewayServer.ListenAndServe(); err != nil {
	//		log.Fatalf("Failed to serve HTTP: %v", err)
	//	}
	//}()

	//

	config.Setup()
	grpcServer := grpc.NewServer()

	config.RegisterDataServer(grpcServer)

	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.LoadedConfig.Server.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		log.Printf("Starting gRPC server on port %s", config.LoadedConfig.Server.GRPCPort)
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
	err = pb.RegisterUserDataHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("localhost:%s", config.LoadedConfig.Server.GRPCPort), opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	// Start HTTP server
	log.Printf("Starting HTTP server on port %s", config.LoadedConfig.Server.HTTPPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.LoadedConfig.Server.HTTPPort), gwmux); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
