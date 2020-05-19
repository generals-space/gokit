package model

import (
	"errors"
)

// User ...
type User struct {
	Name    string
	Title   string
	Company string
}

// UserManager ...
type UserManager struct {
	Users []*User
}

// ErrUserNotFound ...
var ErrUserNotFound = errors.New("目标用户不存在")
var ServerAddr = ":7718"

// GetUser ...
func (m *UserManager) GetUser(name string) (user *User, err error) {
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

