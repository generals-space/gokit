参考文章

1. [三、go-kit 与 grpc 结合实现注册发现与负载均衡](https://hacpai.com/article/1524894068545)

此工程为[ru-rocker/gokit-playground](https://github.com/ru-rocker/gokit-playground/tree/master/lorem-consul)的精简版, 移除了日志, API监控部分的代码, 添加了注释.

本示例在一定程度上借鉴了参考文章1, 尤其是注释方面, 有助于对go-kit组件的理解.

...上一个示例[61.gokit-lorem-consul]()没有使用原作中把client当作一个http服务的模式, 而是直接从客户端请求服务端接口然后打印结果. 

但是因为下一个示例, 原作中关于断路器的文章示例, 需要client作为一个服务来运行. 所以本例要延用原作的示例了...真香.

使用docker-compose启动后可以通过curl直接访问两个server端, 示例如下

```
$ curl -XPOST http://localhost:8081/lorem/sentence/5/20
{"message":"Inde abs contra scrutamur benedicendo quendam ita nam concurrunt diu passionis pax specto aut sectatur pede aer."}
```

也可以访问client端, 示例如下

```
$ curl -X POST -d '{"requestType":"sentence", "min":5, "max":20}' http://localhost:8090/sd-lorem
{"message":"Es extra se intonas tangunt corrigebat exclamaverunt horum hi immo diligi se da sequi me regina da omne."}
```
