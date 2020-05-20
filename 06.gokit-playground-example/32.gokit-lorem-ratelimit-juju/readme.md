此工程为[ru-rocker/gokit-playground](https://github.com/ru-rocker/gokit-playground/tree/master/lorem-rate-limit)的精简版.

在原文[Micro-services Using go-kit: Rate Limiting](http://www.ru-rocker.com/2017/03/19/micro-services-using-go-kit-rate-limiting/)中提到了两种可引用的令牌桶算法.

一种是[Juju](https://github.com/juju/ratelimit), 另一种说是go-kit内置的中间件.

在示例[31.gokit-lorem-ratelimit]()中我只使用了`golang/x/time/rate`库, 而在本例中替换为了`Juju`, 作用相同.

超过流速的请求需要通过同时手动发起n个请求以查看各响应来比较, 本示例的上限是每秒3个请求.

先启动 server 服务

```console
$ go run cmd/main.go
Starting server at port 8080
```

常规请求如下

```console
$ curl -XPOST localhost:8080/lorem/sentence/1/20
{"message":"Concurrunt nota re dicam fias sim aut pecco die appetitum ea mortalitatis hi."}
```

如下使用多开终端进行模拟则可以看到限流效果.

![](https://gitee.com/generals-space/gitimg/raw/master/8917e1c9dc8d0ffae4bda93cec3fa867.jpg)
