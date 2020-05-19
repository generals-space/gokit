## protobuf 使用方法

```
go install github.com/golang/protobuf/protoc-gen-go
```

然后从这个页面[Protocol Buffers Release页](https://github.com/protocolbuffers/protobuf/releases)下载`protoc`可执行文件.

使用如下命令生成`pb`文件.

```
cd pkg/model
protoc --go_out=plugins=grpc:. ./user_manager.proto
```

> `protoc`的`--go_out`选项会调用`protoc-gen-go`工具, 所以`$GOPATH/bin`也需要添加到`$PATH`路径中.

## 设备思路

gprc层本身与业务逻辑服务就是分开实现的.

我们需要把grpc传入的protobuf参数转化为业务函数需要的格式, 同时也需要把业务层的响应转化为grpc的protobuf结果.

`rpc_server`中, `xxxServiceServer`部分的结构体一般都可为空, 与实际的业务服务分离. 因为`xxxService`在`protobuf`文件中只定义了`rpc`字段, 也就是要实现的方法, 而没有定义成员变量.

在业务逻辑与rpc层分离的架构中, 可以在rpc server启动前, 先启动业务服务. 然后`xxxServiceServer`的方法都去调用业务服务暴露出来的方法. 

这样做的缺点就是之前所说的, 多了请求与响应之间的转换层. 这个转换层大概就放在`xxxServiceServer`的实现的`rpc`成员方法中.

当然, 也可以将业务逻辑与rpc的`xxxServiceServer`结合, 但我觉得这样的架构未免不清晰...也可能是我过度设计了.

## 程序执行方法

先启动`server`端服务

```console
$ go run cmd/server/main.go
2020/05/19 22:50:26 server: 启动监听
2020/05/19 22:50:26 server: 注册服务
2020/05/19 22:50:26 server: 等待连接
```

再启动`client`服务

```console
$ go run cmd/client/main.go
2020/05/19 22:51:25 client: 启动客户端
2020/05/19 22:51:25 client: 连接成功
2020/05/19 22:51:25 client: 查询用户: 马云
2020/05/19 22:51:25 Name:"马云"  Title:"市场总监"  Company:"阿里"
2020/05/19 22:51:25 姓名: 马云
2020/05/19 22:51:25 职位: 市场总监
2020/05/19 22:51:25 公司: 阿里
2020/05/19 22:51:25 client: 李彦宏升职为CEO
2020/05/19 22:51:25 Name:"李彦宏"  Title:"CEO"  Company:"百度"
2020/05/19 22:51:25 姓名: 李彦宏
2020/05/19 22:51:25 职位: CEO
2020/05/19 22:51:25 公司: 百度
2020/05/19 22:51:25 client: 委派马化腾到深圳
2020/05/19 22:51:25 Name:"马化腾"  Title:"商务专员"  Company:"深圳"
2020/05/19 22:51:25 姓名: 马化腾
2020/05/19 22:51:25 职位: 商务专员
2020/05/19 22:51:25 公司: 深圳
```

> 在微服务架构中, 应该不能说 server 是服务端, client 就是客户端. ta们的角色是平等的, 只不过 client 需要调用 server 提供的服务而已. 而在实际场景中, client 应该是与 server 保持持久连接, 而不是执行完成就断开.
