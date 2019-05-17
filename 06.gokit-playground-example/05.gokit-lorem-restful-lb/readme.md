参考文章

1. [go-kit 上手之example stringsvc3 通过代理实现分布式处理](https://blog.csdn.net/wdy_yx/article/details/78389736)

注意本例使用的代码基于[62.gokit-lorem-consul-client](), 使用了其中的C/S架构, 但是没有使用consul. 因为本例所要展现的不是服务发现, 而是负载均衡机制.

由于单独的负载均衡基本上没有实例意义(你可以选择使用kong/nginx来做), 对实际场景应该没有太大用处, 所以这里只是加深对go-kit的理解, 仅供参考.

本文最关键操作就是如果通过已知服务地址构造`endpointer`对象 - gokit的`sd.FixedEndpointer`.

本例列举出了两种方法, 都有效.

使用docker-compose启动后, 连续访问客户端, 可得到如下输出.

```
$ curl -XPOST -d '{"requestType":"sentence", "min":5, "max":20}' http://localhost:8090/sd-lorem
{"message":"Re cavernis tot ipsa qui at os me indicabo, te tuetur audi pati sim da ita."}
```

而在服务端, 两个节点将分别处理请求, 可见负载均衡器有效.
