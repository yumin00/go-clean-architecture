package domain

import "context"

type UserUsecase interface {
	GetUserInfoById(ctx context.Context, id int32) (*UserInfo, error)
}

type UserRepository interface {
	GetUserInfoById(ctx context.Context, id int32) (*UserInfo, error)
}

type UserInfo struct {
	Id              int32
	Name            string
	Email           string
	ProfileImageUrl string
}
