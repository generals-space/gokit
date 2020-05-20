package main

import (
	"log"
	"net/http"

	"gokit/common"
	"gokit/department"
	"gokit/usermanager"
)
/*
	`main.go`入口程序其实是微服务架构中的网关API服务, 实现客户端请求的路由服务.
	后端存在 user manager 和 department manager 两个服务.

	当然实际场景中, 网关与后端微服务的连接应该是通过 restful 或是 grpc 的形式, 
	而不是这里"内嵌"的形式.
*/
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
