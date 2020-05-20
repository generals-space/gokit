此工程为[ru-rocker/gokit-playground](https://github.com/ru-rocker/gokit-playground/tree/master/lorem-grpc)的精简版, 简化了部分代码, 添加了一些注释.

基本上与[04.gokit-lorem-restful-client]()完全相同, 只不过 client 与 server 服务间的通信用 grpc 形式而非 restful.

先启动 server 服务

```console
$ go run cmd/server/main.go
Starting server
```

然后启动 client 对 server 服务发起请求, 完成后 client 运行结束.

```console
$ go run cmd/client/main.go
Concurrunt nota re dicam fias, sim aut pecco, die appetitum.
```

------

这一示例中有一个小bug, 在`client/main.go`中调用`grpctransport.NewClient()`时, 第二个参数(即服务名称)不应为`Lorem`, 这导致执行客户端时出现如下错误

```
rpc error: code = Unimplemented desc = unknown service Lorem
exit status 1
```

按照官方issue[Bad service name](https://github.com/ru-rocker/gokit-playground/issues/1)给出的解决方案, 将`Lorem` -> `pb.Lorem`成功解决.
