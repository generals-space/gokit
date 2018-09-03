package service

import (
	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
)

// DepartmentManager ...
type DepartmentManager struct{
	Departments []*common.Department
}

func (m *DepartmentManager)Create