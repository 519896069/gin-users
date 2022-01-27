package routes

import (
	"google.golang.org/grpc"
	"user/app/http/controllers/rpc"
	"user/app/services"
)

func RpcLoadRouter(sev *grpc.Server) {
	services.RegisterUserServicesServer(sev, &rpc.UserServer{})
}
