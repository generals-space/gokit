package main

import (
	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/usermanager"
)

var uManagerService *usermanager.UserManager

func init() {
	// 真实场景中, 业务服务就不会有这么简单了.
	// 可能涉及到读取配置, 数据库连接等操作.
	// 这里可以简单地把UserManager看成一张表, ta的成员方法就是传统的CURD的模拟操作.
	uManagerService = &usermanager.UserManager{
		Users: []*common.User{
			&common.User{
				Name:    "李彦宏",
				Company: "北京",
			},
			&common.User{
				Name:    "马云",
				Company: "杭州",
			},
			&common.User{
				Name:    "马化腾",
				Company: "广州",
			},
		},
	}
}
