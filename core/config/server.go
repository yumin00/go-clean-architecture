package config

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	userPb "github.com/yumin00/go-hexagonal/go-proto/go-api/core/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strings"
)

func RegisterDataServer(server *grpc.Server) {
	userPb.RegisterUserDataServer(server, NewUserServer())
}

type RegisterHandlerFunc func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error

func Start(ctx context.Context) (grpcServer *grpc.Server, gatewayServer *http.Server) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	endpoint := fmt.Sprintf("localhost:%s", LoadedConfig.Server.GRPCPort)

	if err := registerHandlers(ctx, mux, endpoint, opts); err != nil {
		return
	}

	grpcServer = grpc.NewServer()
	RegisterDataServer(grpcServer)

	gatewayServer = &http.Server{
		Addr: LoadedConfig.Server.HTTPPort,
		//Handler:           allowCORS(mux),
		//ReadHeaderTimeout: 60 * time.Second,
	}

	return grpcServer, gatewayServer
}

func registerHandlers(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) error {
	handlers := []RegisterHandlerFunc{
		userPb.RegisterUserDataHandlerFromEndpoint,
	}

	for _, handler := range handlers {
		if err := handler(ctx, mux, grpcServerEndpoint, opts); err != nil {
			return fmt.Errorf("failed to register handler: %v", err)
		}
	}

	return nil
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "X-Requested-With", "traceparent"}
	//headers = append(headers, customHeaders...)
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PATCH", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Print("preflight request for %s", r.URL.Path)
}
