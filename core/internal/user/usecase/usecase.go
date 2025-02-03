package usecase

import (
	"context"

	"github.com/yumin00/go-clean-architecture/core/domain"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) GetUserInfoById(ctx context.Context, id int32) (*domain.UserInfo, error) {
	userInfo, err := u.userRepo.GetUserInfoById(ctx, id)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
