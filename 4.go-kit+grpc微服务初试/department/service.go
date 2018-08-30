package service

import (
	"github.com/generals-space/gokit/4.go-kit+grpc微服务初试/common"
)

// Department ...
type Department struct{
	Name string
	City string
	Users []*common.User
}