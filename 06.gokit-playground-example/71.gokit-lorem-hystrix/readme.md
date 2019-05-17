此工程为[ru-rocker/gokit-playground](https://github.com/ru-rocker/gokit-playground/tree/master/lorem-hystrix)的精简版, 移除了日志, API监控部分的代码, 添加了注释.

此工程延用了lorem-consul的代码, 主要是为了在负载均衡的场景下, 能够通过断路器模式隔绝对故障服务节点的访问.

使用docker-compose启动服务后, 可以通过curl直接访问两个server端, 示例如下

```
$ curl -XPOST http://localhost:8081/lorem/sentence/5/20
{"message":"Inde abs contra scrutamur benedicendo quendam ita nam concurrunt diu passionis pax specto aut sectatur pede aer."}
```

也可以访问client端, 示例如下

```
$ curl -X POST -d '{"requestType":"sentence", "min":5, "max":20}' http://localhost:8090/sd-lorem
{"message":"Es extra se intonas tangunt corrigebat exclamaverunt horum hi immo diligi se da sequi me regina da omne."}
```

为了展示断路器工作正常, 且能够显示fallback信息, 这里以循环形式发起请求.

```
while true; do curl -XPOST -d'{"requestType":"word", "min":10, "max":10}' http://localhost:8080/sd-lorem; sleep 1; done;
while true; do curl -XPOST -d'{"requestType":"sentence", "min":10, "max":10}' http://localhost:8080/sd-lorem; sleep 1; done;
while true; do curl -XPOST -d'{"requestType":"paragraph", "min":10, "max":10}' http://localhost:8080/sd-lorem; sleep 1; done;
```

此后可以通过关闭/启动server端来查看请求的结果. 你可以看到如下类似的信息.

```
{"message":"curiosarum"}
{"message":"caecitatem"}
{"fallback":"Service currently unavailable"}
{"fallback":"Service currently unavailable"}
{"message":"nonnullius"}
```

然后访问`http://localhost:8181/hystrix`, 配置断流器监听地址, 正是在client中配置的`hystrixStreamHandler`的地址. 开始监听后, 可以展示出请求的数据, 这里就不多说了.
