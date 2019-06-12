参考文章

1. [go-kit官方issue REST style URLs with HTTP Client](https://github.com/go-kit/kit/issues/741)

2. [opencensus-gokit-example](https://github.com/basvanbeek/opencensus-gokit-example)

在原本的restful示例中, 我们是通过curl来访问的. 但是在微服务间, restful接口相互调用应该通过什么工具呢? 本示例是受到grpc示例的影响, 编写的restful接口客户端示例代码.

在go-kit中, 其实也是有对应的http微服务客户端的. 与grpc接口一样, 面对restful接口, 我们也需要提前知道API地址. 不过grpc客户端一般需要将链接对象在服务内部加以维护, 而restful客户端不需要这样做.

http transport也有`NewClient()`函数, 但是面对restful接口, 如何将我们想要的参数添加到路径中呢? 

我查阅了官方issue, 参考文章1给出了解决方案. 与grpc客户端相同, 我们需要为restful客户端定义`EncodeLoremRequest`和`DecodeLoremResponse`, 前者将`LoremRequest`请求对象转换成restful形式的`/lorem/{type}/{min}/{max}`格式的url.

这一做法借鉴了参考文章2工程的方式, ta将所有http接口集合成了一张"路由表", 服务端挂载handler与客户端创建endpoint时用同一张表就可以将`mux.NewRouter()`操作抽离出来, 正是本例中新增的`routes.go`文件.

另外提一下, 参考文章2(参考文章1中提到的示例)是一个不错的微服务工程, 代码量不多, 但确是一个完整的项目, 以后可以参考.

使用docker-compose启动后, 可以通过curl直接访问server端, 示例如下

```
$ curl -XPOST http://localhost:8080/lorem/sentence/5/20
{"message":"Inde abs contra scrutamur benedicendo quendam ita nam concurrunt diu passionis pax specto aut sectatur pede aer."}
```

也可以多次重启client服务, 查看日志会发现请求的结果.
