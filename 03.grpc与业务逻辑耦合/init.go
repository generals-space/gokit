package main

import "errors"

// ErrUserNotFound ...
var ErrUserNotFound = errors.New("目标用户不存在")
var ServerAddr = ":7718"
var uManagerServiceServer *UManagerServiceServer

// User ...
type User struct {
	Name    string
	Title   string
	Company string
}

func init() {
	// 真实场景中, 业务服务就不会有这么简单了.
	// 可能涉及到读取配置, 数据库连接等操作.
	// 这里可以简单地把UserManager看成一张表, ta的成员方法就是传统的CURD的模拟操作.
	uManagerServiceServer = &UManagerServiceServer{
		Users: []*User{
			&User{
				Name:    "李彦宏",
				Title:   "总经理",
				Company: "百度",
			},
			&User{
				Name:    "马云",
				Title:   "市场总监",
				Company: "阿里",
			},
			&User{
				Name:    "马化腾",
				Title:   "商务专员",
				Company: "腾讯",
			},
		},
	}
}
