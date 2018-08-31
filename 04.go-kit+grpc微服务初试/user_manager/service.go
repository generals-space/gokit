package user_manager

import (
	"errors"
	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
)

// UserManager ...
type UserManager struct {
	Users []*common.User
}

// ErrUserNotFound ...
var ErrUserNotFound = errors.New("目标用户不存在")
// GetUser ...
func (m *UserManager) GetUser(name string) (user *common.User, err error) {
	for _, u := range m.Users {
		if u.Name == name {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}

// Dispatch ...
func (m *UserManager) Dispatch(name, company string) (err error) {
	for _, u := range m.Users {
		if u.Name == name {
			u.Company = company
			return nil
		}
	}
	return ErrUserNotFound
}

// SetTitle ...
func (m *UserManager) SetTitle(name, title string) (err error) {
	for _, u := range m.Users {
		if u.Name == name {
			u.Title = title
			return nil
		}
	}
	return ErrUserNotFound
}

