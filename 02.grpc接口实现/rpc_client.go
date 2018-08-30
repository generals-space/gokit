package main

import (
	"log"
	"time"
	"context"
	"google.golang.org/grpc"
)

// NewClient ...
func NewClient(){
	log.Println("client: 启动客户端")
	conn, err := grpc.Dial(ServerAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Println("client: 连接成功")
	uManagerServiceClient := NewUserManagerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	log.Println("client: 查询用户: 马云")
	user, err := uManagerServiceClient.GetUser(ctx, &GetUserRequest{Name: "马云"})
	if err != nil {
		log.Fatalln(err)
	}else{
		log.Printf("%+v\n", user)
		log.Printf("姓名: %s\n", user.Name)
		log.Printf("职位: %s\n", user.Title)
		log.Printf("公司: %s\n", user.Company)
	}

	log.Println("client: 李彦宏升职为CEO")
	_, err = uManagerServiceClient.SetTitle(ctx, &SetTitleRequest{
		Name: "李彦宏", Title: "CEO",
	})
	if err != nil {
		log.Fatalln(err)
	}else{
		user, err = uManagerServiceClient.GetUser(ctx, &GetUserRequest{Name: "李彦宏"})
		log.Printf("%+v\n", user)
		log.Printf("姓名: %s\n", user.Name)
		log.Printf("职位: %s\n", user.Title)
		log.Printf("公司: %s\n", user.Company)
	}

	log.Println("client: 委派马化腾到深圳")
	_, err = uManagerServiceClient.Dispatch(ctx, &DispatchRequest{
		Name: "马化腾", Company: "深圳",
	})
	if err != nil {
		log.Fatalln(err)
	}else{
		user, err = uManagerServiceClient.GetUser(ctx, &GetUserRequest{Name: "马化腾"})
		log.Printf("%+v\n", user)
		log.Printf("姓名: %s\n", user.Name)
		log.Printf("职位: %s\n", user.Title)
		log.Printf("公司: %s\n", user.Company)
	}
}