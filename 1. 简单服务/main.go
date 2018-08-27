package main

import (
	"errors"
	"log"
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

var userManager *UserManager

func init() {
	userManager = &UserManager{
		Users: []*User{
			&User{
				Name:    "李彦宏",
				Title:   "总经理",
				Company: "北京",
			},
			&User{
				Name:    "马云",
				Title:   "市场总监",
				Company: "杭州",
			},
			&User{
				Name:    "马化腾",
				Title:   "商务专员",
				Company: "广州",
			},
		},
	}
}

func main() {
	for _, u := range userManager.Users {
		log.Printf("%+v\n", u)
	}
	var err error
	err = userManager.Dispatch("马化腾", "北京")
	if err != nil {
		log.Fatalln(err)
	}
	err = userManager.SetTitle("马云", "CFO")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("=====================================")
	for _, u := range userManager.Users {
		log.Printf("%+v\n", u)
	}
	
	log.Println("搜索马云...")
	mayun, err := userManager.GetUser("马云")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", mayun)
}
