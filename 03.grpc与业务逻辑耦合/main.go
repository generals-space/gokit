package main

import "time"

func main(){
	// 启动grpc服务端.
	// 注意: 业务逻辑服务由于写在`init.go`中, 所以已经事先启动. 
	// 在实际场景中, 也应该是先启动业务服务, 再启动grpc接口.
	go NewServer()
	time.Sleep(time.Second * 3)
	// 客户端操作
	NewClient()
}