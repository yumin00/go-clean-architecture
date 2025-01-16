package delivery

import (
	"context"
	pb "github.com/yumin00/go-hexagonal/go-hexagonal-proto/go-api/core/user"
	"github.com/yumin00/go-hexagonal/core/domain"
)

type Server struct {
	pb.UserDataServer
	UserUsecase domain.UserUsecase
}

func (s *Server) GetUserInfoById(ctx context.Context, req *pb.GetUserInfoByIdRequest) (_ *pb.GetUserInfoByIdResponse, err error) {
	userId := req.GetId()

	userInfo, err := s.UserUsecase.GetUserInfoById(ctx, userId)
	if err != nil {
		return nil, err
	}

	var userInfoPb *pb.UserInfo
	if userInfo != nil {
		userInfoPb = &pb.UserInfo{
			Id:              userInfo.Id,
			Name:            userInfo.Name,
			Email:           userInfo.Email,
			ProfileImageUrl: userInfo.ProfileImageUrl,
		}
	}

	return &pb.GetUserInfoByIdResponse{
		UserInfo: userInfoPb,
	}, nil
}
