此工程为[ru-rocker/gokit-playground](https://github.com/ru-rocker/gokit-playground/tree/master/lorem-rate-limit)的精简版.

在原文[Micro-services Using go-kit: Rate Limiting](http://www.ru-rocker.com/2017/03/19/micro-services-using-go-kit-rate-limiting/)中提到了两种可引用的令牌桶算法.

一种是[Juju](https://github.com/juju/ratelimit), 另一种说是go-kit内置的中间件.

然而在研究原作的示例工程时我发现原工程中直接使用了`golang.org/x/time/rate`限流器, 原作文章中所提到的`Juju`和`instrument.go`根本没有用到.

可以说`time/rate`与`Juju`实现了同样的功能, 本例代码中移除了`instrument.go`, 其余操作与最初的restful工程完全相同.

注意在原作示例中限流提示是写入到日志文件中的, 其实就是为endpoint增加一个log中间件, 本例中移除了日志中间件. 超过流速的请求需要通过同时手动发起n个请求以查看各响应来比较, 本示例的上限是每秒3个请求.

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

![](https://gitee.com/generals-space/gitimg/raw/master/f209a5f66dc03fcd4932a27ac7b775d1.jpg)
