package config

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	userPb "github.com/yumin00/go-clean-architecture/go-proto/go-api/core/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
)

func RegisterDataServer(server *grpc.Server) {
	userPb.RegisterUserDataServer(server, NewUserServer())
}

type RegisterGRPCHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

type GRPCRegistration struct {
	RegisterHandler RegisterGRPCHandler
}

func getGRPCRegistrations() []GRPCRegistration {
	return []GRPCRegistration{
		{RegisterHandler: userPb.RegisterUserDataHandlerFromEndpoint},
	}
}

func Start(ctx context.Context) (*grpc.Server, *http.Server) {
	grpcServer := grpc.NewServer()
	RegisterDataServer(grpcServer)

	gwmux := runtime.NewServeMux(
		runtime.WithErrorHandler(runtime.DefaultHTTPErrorHandler),
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			return metadata.New(map[string]string{
				"content-type": "application/json",
			})
		}),
	)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	endpoint := fmt.Sprintf("localhost:%s", LoadedConfig.Server.GRPCPort)

	for _, registration := range getGRPCRegistrations() {
		if err := registration.RegisterHandler(ctx, gwmux, endpoint, opts); err != nil {
			log.Fatalf("Failed to register gateway: %v", err)
		}
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", LoadedConfig.Server.HTTPPort),
		Handler: gwmux,
	}

	return grpcServer, gwServer
}
