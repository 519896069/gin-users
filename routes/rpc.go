package routes

import (
	"google.golang.org/grpc"
	"user/app/http/controllers/rpc"
	"user/app/services/pb"
)

func RpcLoadRouter(sev *grpc.Server) {
	pb.RegisterUserServicesServer(sev, &rpc.UserServer{})
}
