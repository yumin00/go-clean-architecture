package config

import (
	userDelivery "github.com/yumin00/go-hexagonal/core/internal/user/delivery"
	userRepository "github.com/yumin00/go-hexagonal/core/internal/user/repository/db"
	userUsecase "github.com/yumin00/go-hexagonal/core/internal/user/usecase"
)

func NewUserServer() *userDelivery.Server {
	s := new(userDelivery.Server)

	userRepo := userRepository.NewDBUserRepository(DB)

	s.UserUsecase = userUsecase.NewUserUsecase(userRepo)

	return s
}
