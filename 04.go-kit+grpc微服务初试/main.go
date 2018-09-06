package main

import (
	"log"
	"net/http"

	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/department"
	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/usermanager"
)

func main() {
	uManagerService := usermanager.NewUserManagerService()
	go usermanager.StartGrpcTransport(uManagerService)

	dManagerService := department.NewDepartmentManagerService()
	go department.StartGrpcTransport(dManagerService)

	// 启动http服务, 面向用户的单一入口.
	// 通过transport添加路由及ta们各自的处理函数
	go usermanager.StartHTTPTransport(uManagerService)
	go department.StartHTTPTransport(dManagerService)

	log.Fatal(http.ListenAndServe(common.GatewayAPIServerAddr, nil))
	log.Println("exit")
}
