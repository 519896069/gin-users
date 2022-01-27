package bootstrap

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
	"user/routes"
)

func StartRpc() {
	//rpc server
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	rpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute, //这个连接最大的空闲时间，超过就释放，解决proxy等到网络问题（不通知grpc的client和server）
		}),
	)
	routes.RpcLoadRouter(rpcServer)
	go func() {
		fmt.Printf("Listening and serving TCP on %v\n", ":8888")
		err = rpcServer.Serve(listen)
		if err != nil {
			fmt.Printf("listen: %s\n", err)
		}
	}()
}
