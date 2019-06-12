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

构建好docker镜像后通过docker-compose启动, 最开始client应该是会请求失败的, 因为服务还没来得及注册.

当服务端注册完成后, 查看其client日志, 可以看到如下输出

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

因为启动了2个服务端, 客户端也拥有负载均衡机制, 可以查看服务端日志, 看看负载均衡是否有效.

另外, 在服务运行期间停止server1/server2, 你会发现client可能会出现`no endpoints available`, 但是在之后的请求就会自动调整, 转发到正常的服务中去. 如果重新启动server1/server2, 请求又会被转发进入, 还是蛮智能的.

还过我还是按照原作所说, 把客户端也做成了http服务. 见[62.gokit-lorem-consul-client]().
