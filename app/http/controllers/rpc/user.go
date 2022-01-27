package rpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/app/business"
	"user/app/models"
	"user/app/services/pb"
)

type UserServer struct {
	pb.UnimplementedUserServicesServer
}

func (UserServer) GetUserByToken(ctx context.Context, token *pb.Token) (*pb.User, error) {
	if !business.CheckToken(token.Token) {
		return nil, status.Errorf(codes.Unimplemented, "token is empty")
	}
	if models.AuthUser == nil {
		return nil, status.Errorf(codes.Unimplemented, "no user")
	}
	return &pb.User{
		Id:       int64(models.AuthUser.ID),
		Username: models.AuthUser.Username,
		Avatar:   models.AuthUser.Avatar,
		Email:    models.AuthUser.Email,
		Mobile:   models.AuthUser.Mobile,
		Status:   int64(models.AuthUser.Status),
	}, nil
}
