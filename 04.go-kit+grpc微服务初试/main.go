package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/usermanager"
	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/department"
)

var uManagerService *usermanager.UserManager
var dManagerService *department.DepartmentManager

func init() {
	// 真实场景中, 业务服务就不会有这么简单了.
	// 可能涉及到读取配置, 数据库连接等操作.
	// 这里可以简单地把UserManager看成一张表, ta的成员方法就是传统的CURD的模拟操作.
	uManagerService = &usermanager.UserManager{
		Users: []*common.User{
			&common.User{
				Name:    "李彦宏",
				Company: "百度",
			},
			&common.User{
				Name:    "马云",
				Company: "阿里",
			},
			&common.User{
				Name:    "马化腾",
				Company: "腾讯",
			},
		},
	}
	dManagerService = &department.DepartmentManager{
		Departments: []*common.Department{
			&common.Department{
				Name: "百度",
				Users: []*common.User{
					&common.User{
						Name: "李彦宏",
						Company: "百度",
					},
				},
			},
			&common.Department{
				Name: "阿里",
				Users: []*common.User{
					&common.User{
						Name: "马云",
						Company: "阿里",
					},
				},
			},
			&common.Department{
				Name: "腾讯",
				Users: []*common.User{
					&common.User{
						Name: "马化腾",
						Company: "腾讯",
					},
				},
			},
		},
	}
}

func main() {
	sigChannel := make(chan os.Signal)
	exitChannel := make(chan bool)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)
	go usermanager.StartGrpcTransport(uManagerService)
	go usermanager.StartHTTPTransport(uManagerService)

	go department.StartGrpcTransport(dManagerService)
	go department.StartHTTPTransport(dManagerService)
	go func() {
		sig := <-sigChannel
		log.Println(sig)
		exitChannel <- true
	}()

	<-exitChannel
	close(exitChannel)
	log.Println("exit")
}
