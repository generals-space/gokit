此工程为[ru-rocker/gokit-playground](https://github.com/ru-rocker/gokit-playground/tree/master/lorem-grpc)的精简版, 简化了部分代码, 添加了一些注释.

首先构建项目的依赖镜像(只是下载项目依赖包)

```
docker build --no-cache=true -f dep.dockerfile -t gokit-lorem-grpc .
```

然后通过`docker-compose`启动.

```
docker-componse up -d
```

由于是开发环境, 所以将项目路径挂载到容器中, 修改代码后重启服务就可以看到效果, 不用重新构建.

`client`容器启动后会立即停止, 代码的执行结果要通过日志来查看.

```
$ dk logs -f 11gokit-lorem-grpc_client_1
Aula me se coepta se lux res, inter motus.
```

如果对要执行的命令有修改, 可以直接修改源代码, 然后通过`docker-compose`重启`client`即可.

------

这一示例中有一个小bug, 在`client/main.go`中调用`grpctransport.NewClient()`时, 第二个参数(即服务名称)不应为`Lorem`, 这导致执行客户端时出现如下错误

```
rpc error: code = Unimplemented desc = unknown service Lorem
exit status 1
```

按照官方issue[Bad service name](https://github.com/ru-rocker/gokit-playground/issues/1)给出的解决方案, 将`Lorem` -> `pb.Lorem`成功解决.
