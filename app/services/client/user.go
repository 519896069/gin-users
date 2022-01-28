package client

import (
	"context"
	"fmt"
	"user/app/services/pb"
	"user/fzp/rpc"
)

func GetUserByToken(token string) *pb.User {
	conn, err := rpc.GetConn()
	if err != nil {
		return &pb.User{}
	}
	defer conn.Close()
	client := pb.NewUserServicesClient(conn)
	user, err := client.GetUserByToken(context.Background(), &pb.Token{
		Token: token,
	})
	if err != nil {
		fmt.Println()
		return &pb.User{}
	}
	return user
}
