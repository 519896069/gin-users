package rpc

import (
	"context"
	"github.com/gin-gonic/gin"
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
	ginCtx := &gin.Context{}
	if !business.CheckToken(ginCtx, token.Token) {
		return nil, status.Errorf(codes.Unimplemented, "token is empty")
	}
	user := models.Auth(ginCtx)
	if user == nil {
		return nil, status.Errorf(codes.Unimplemented, "no user")
	}
	return &pb.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Mobile:   user.Mobile,
		Status:   int64(user.Status),
	}, nil
}
