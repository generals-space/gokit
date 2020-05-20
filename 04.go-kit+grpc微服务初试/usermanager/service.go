package usermanager

import (
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/grpc"

	"gokit/common"
)

// UserManager ...
type UserManager struct {
	Users               []*common.User
	DManagerServiceConn *grpc.ClientConn
	DManagerServiceCli  common.DepartmentManagerServiceClient
}

// ErrUserNotFound ...
var ErrUserNotFound = errors.New("目标用户不存在")

// List ...
func (m *UserManager) List() (users []*common.User, err error) {
	return m.Users, nil
}

// Dispatch ...
func (m *UserManager) Dispatch(name, company string) (err error) {
	for _, u := range m.Users {
		if u.Name == name {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel()

			_, err = m.DManagerServiceCli.PersonnelChange(ctx, &common.PersonnelChangeRequest{
				User: u,
				Company: company,
			})
			u.Company = company

			return err
		}
	}
	return ErrUserNotFound
}

// AddUser ...
func (m *UserManager) AddUser(userList *common.UserList) (err error) {
	m.Users = append(m.Users, userList.List...)
	return
}

// NewUserManagerService ...
func NewUserManagerService() *UserManager {
	// 真实场景中, 业务服务就不会有这么简单了.
	// 可能涉及到读取配置, 数据库连接等操作.
	// 这里可以简单地把UserManager看成一张表, ta的成员方法就是传统的CURD的模拟操作.
	userManager := &UserManager{
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
	go connectDManagerService(userManager)
	return userManager
}

func connectDManagerService(userManager *UserManager) {
	conn, err := grpc.Dial(common.DepartmentGrpcTransportAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	dManagerServiceCli := common.NewDepartmentManagerServiceClient(conn)
	userManager.DManagerServiceConn = conn
	userManager.DManagerServiceCli = dManagerServiceCli
	log.Println("department manager service connected")
}
