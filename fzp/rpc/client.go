package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetConn() (*grpc.ClientConn, error) {
	opt := insecure.NewCredentials()
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(opt))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
