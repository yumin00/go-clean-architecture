package main

import (
	"context"
	"fmt"
	"github.com/yumin00/go-clean-architecture/core/config"
	"log"
	"net"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	config.Setup()

	grpcServer, gatewayServer := config.Start(ctx)

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

	log.Printf("Starting HTTP server on port %s", config.LoadedConfig.Server.HTTPPort)
	if err := gatewayServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
