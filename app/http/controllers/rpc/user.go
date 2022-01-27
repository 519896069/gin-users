package rpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/app/business"
	"user/app/models"
	"user/app/services"
)

type UserServer struct {
	services.UnimplementedUserServicesServer
}

func (UserServer) GetUserByToken(ctx context.Context, token *services.Token) (*services.User, error) {
	if !business.CheckToken(token.Token) {
		return nil, status.Errorf(codes.Unimplemented, "token is empty")
	}
	if models.AuthUser == nil {
		return nil, status.Errorf(codes.Unimplemented, "no user")
	}
	return &services.User{
		Id:       int64(models.AuthUser.ID),
		Username: models.AuthUser.Username,
		Avatar:   models.AuthUser.Avatar,
		Email:    models.AuthUser.Email,
		Mobile:   models.AuthUser.Mobile,
		Status:   int64(models.AuthUser.Status),
	}, nil
}
