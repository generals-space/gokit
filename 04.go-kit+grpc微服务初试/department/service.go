package department

import (
	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
)

// DepartmentManager ...
type DepartmentManager struct {
	Departments []*common.Department
}

// List ...
func (m *DepartmentManager) List() (departmentList []*common.Department, err error) {
	return m.Departments, nil
}

// Create ...
func (m *DepartmentManager) Create(department *common.Department) (err error) {
	m.Departments = append(m.Departments, department)
	return
}
