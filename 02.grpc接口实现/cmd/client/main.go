package main

import "gokit/pkg/client"

func main() {
	// 客户端操作, 执行完毕就会退出.
	// 实际场景中一般为维持此连接
	client.NewClient()
}
