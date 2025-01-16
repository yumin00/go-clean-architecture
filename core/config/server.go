package config

import (
	userPb "github.com/yumin00/go-hexagonal/go-hexagonal-proto/go-api/core/user"
	"google.golang.org/grpc"
)

func RegisterDataServer(server *grpc.Server) {
	userPb.RegisterUserDataServer(server, NewUserServer())
}
