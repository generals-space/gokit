参考文章

1. [三、go-kit 与 grpc 结合实现注册发现与负载均衡](https://hacpai.com/article/1524894068545)

本文使用etcd作为服务注册中间件来代替consul, 借鉴了参考文章1. 

使用docker-compose启动后, 可以通过curl直接访问两个server端, 示例如下

```
$ curl -XPOST http://localhost:8081/lorem/sentence/5/20
{"message":"Inde abs contra scrutamur benedicendo quendam ita nam concurrunt diu passionis pax specto aut sectatur pede aer."}
```

也可以访问client端, 示例如下

```
$ curl -X POST -d '{"requestType":"sentence", "min":5, "max":20}' http://localhost:8090/sd-lorem
{"message":"Es extra se intonas tangunt corrigebat exclamaverunt horum hi immo diligi se da sequi me regina da omne."}
```

------

与consul不同, etcd没有health健康检查机制, 所以启动两个server一段时间后停止server1或server2, 继续请求client, 会有一半的请求失败.

```
$ curl -X POST -d '{"requestType":"sentence", "min":5, "max":20}' http://localhost:8090/sd-lorem
{"message":"Ita me repente formaeque agam nosti da rei ne cui iam ea se."}
$ curl -X POST -d '{"requestType":"sentence", "min":5, "max":20}' http://localhost:8090/sd-lorem
Post http://server1:8080/lorem/sentence/5/20: dial tcp: lookup server1 on 127.0.0.11:53: no such host
```

当然本例中的代码其实是可以实现服务节点的异常检测的, 关键在于`sd/etcdv3`包中`NewTTLOption`的使用. 

其基本原理就是go-kit在为server注册服务时, 设置其key在指定时间m秒内过期, 并且按照每隔n秒刷新这个key的值, 使之保持可用状态(当然m > n).

客户端一直watch这个key的值, 将接收到的请求转发给可用server列表的其中一个.

当server异常停止, 在m秒内其所属key就会因过期而被删除, 客户端也会被通知这一事件, 将该server中列表中移除.

`etcd-sample`简单的实现了go-kit所做的事件, 具体可以见ta的readme文档.
