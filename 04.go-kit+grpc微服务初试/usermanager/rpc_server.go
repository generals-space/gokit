package usermanager

import (
	"log"
	"net"

	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Start ...
func (server *UManagerServiceServer) Start() {
	log.Println("usermanager server: 启动监听")
	lis, err := net.Listen("tcp", common.UserManagerServerAddr)
	if err != nil {
		panic(err)
	}
	rpcServer := grpc.NewServer()
	log.Println("usermanager server: 注册服务")
	common.RegisterUserManagerServiceServer(rpcServer, server)
	reflection.Register(rpcServer)
	log.Println("usermanager server: 等待连接")
	if err := rpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
