参考文章

1. [三、go-kit 与 grpc 结合实现注册发现与负载均衡](https://hacpai.com/article/1524894068545)

此工程为[ru-rocker/gokit-playground](https://github.com/ru-rocker/gokit-playground/tree/master/lorem-consul)的精简版, 移除了日志, API监控部分的代码, 添加了注释.

本示例在一定程度上借鉴了参考文章1, 尤其是注释方面, 有助于对go-kit组件的理解.

## 关于概念

之前一直以为"服务发现"是一个很高级的功能, 想像中是各种服务注册到服务中心, 比如`zk`, `etcd`, `consul`等, 然后子级服务只要知道父级服务的名称或其他的一些信息, 就可以调用父级服务提供的方法了. 

类似于大家都加入了一个QQ群, 只要通过`@昵称(服务名称)`就能和人对话一样(其他人保持静默), 这样每加入一个新人(新服务), 已有的服务在调用时就多了一个选择.

...然而我还是too young. 

作为子级服务, 在调用父级服务时, 必须要先加好友. 虽然你不用再记复杂的QQ号(服务地址), 但仍然要先通过昵称从QQ群里将父级服务找出来, 建立单独的连接, 才能相互沟通.

并且, 并不是每加入一个新人(新服务), 就可以根据昵称(服务名称)随意调用, 那种功能叫做"热加载", 或者说叫"动态引用". 一个服务要依赖哪些父级服务, 在启动前就已经确定了, 启动时建立好与父级服务的连接, 之后保持通信就行了.

## 关于consul

原作中使用的consul镜像为`progrium/consul`, 但本例中使用的为官方镜像`consul`, 启动命令为

```
docker run -d --name consul -p 8500:8500 consul agent -dev -ui -client=0.0.0.0
```

启动后可通过`localhost:8500`访问.

## 关于客户端

本例对原作示例中的客户端改动是比较大的. 在原作示例中, 客户端也提供了http服务, 用户要访问的是客户端的http服务而不再是服务端. 这样就需要在客户端重复定义编解码函数, 并重新挂载路由, 并不直观.

本例中将客户端当作微服务架构中的一个独立的服务, ta依赖于服务端. 在实际场景中, 我们需要在客户端中配置ta所依赖的服务名称(服务地址可以通过consul注册中心得到), 得到其端点地址, 之后发送请求就可以了, 解码得到`LoremResponse`将会是服务内部所需的对象, 无需额外处理.

具体的需要结合代码来理解.

## 测试

构建好docker镜像后通过docker-compose启动, 最开始client应该是会启动失败的, 因为服务还没来得及注册.

当服务端注册完成后, 多次重启client, 查看其日志, 可以看到如下输出

```
$ docker-compose restart client
Restarting 61gokit-lorem-consul_client_1 ... done
gener@LAPTOP-PD3FLKC8 /d/gopath/src/github.com/generals-space/gokit/06.gokit-playground-example/61.gokit-lorem-consul (master)
$ docker logs -f 61gokit-lorem-consul_client_1
ts=2019-05-15T14:18:43.3725567Z caller=instancer.go:48 service=lorem tags="[lorem ru-rocker]" instances=0
endpoints not found
ts=2019-05-15T14:19:39.9502211Z caller=instancer.go:48 service=lorem tags="[lorem ru-rocker]" instances=2
{Tenuiter ipsos modico cui praecedentia in redire redire conprehendant eliqua os tenent iste re re quotiens ac. <nil>}
```

因为启动了2个服务端, 客户端也拥有负载均衡机制, 可以在服务端Lorem方法的代码中添加fmt日志, 看看负载均衡是否有效.

补充:

打脸了.

客户端通过endpoint直接执行方式访问服务端, 虽然可行, 但是负载均衡就不好用了. 每次重启客户端, 访问的都是server1. 只有当server1停止后(一段时间后, 要让consul判断其已经掉线), 重启客户端才会访问到server2.

对于这一点, 我想应该是由于负载均衡器是内置在客户端中, ta内部维护着访问记录, 如果客户端服务一直存在, 这个机制就有作用. 而像我们本例所示, 客户端发现->请求->中止一连串操作, 每次请求都是新的, 所以每次都访问的是第一个节点. 如果在生命周期中连续发送两(多)个请求, 才有可能平均分发到不同服务.

于是我按照原作所说, 把客户端也做成了http服务. 见[62.gokit-lorem-consul-client]().

------

目前我疑惑的是客户端内部需要维护什么对象? 就像数据库连接一样, 只保留consul连接对象应该是不行的吧? 那么, `instancer`, `endpointer`, 还是`balancer`?

...在一个实际的场景中我的确是只保存了consul连接对象, 不过那也是因为每次发来的请求要调用的服务不同, 需要动态获取不同服务的地址.
