参考文章

1. [go-kit offical example](https://github.com/go-kit/kit/tree/master/examples)
    - 官方示例
    - DDD(domain drive design) 领域驱动设计
    - ...过度设计, 不宜用作入门
2. [一、go 语言编写 grpc 微服务实例](https://hacpai.com/article/1524816248447)
    - 系列文章
    - 看看概念和思路就好, 代码示例比较生硬, 而且代码质量并不高.
3. [Micro-services Using Go-kit: REST Endpoint](http://www.ru-rocker.com/2017/02/17/micro-services-using-go-kit-http-endpoint/)
    - 系列文章, 层层递进, 代码详尽, 推荐
    - 限流, 断路器, 监控, 日志等详细的解决方案.
4. [opencensus-gokit-example](https://github.com/basvanbeek/opencensus-gokit-example)
    - 一个完整的项目, 涉及到go-kit的多种组件
5. [go-kit微服务系列目录](https://juejin.im/post/5c861c93f265da2de7138615)
    - 系列文章
    - 条理分明, 概念解释得很清晰
    - 代码质量很高, 结构与风格与参考文章3相似

本仓库中示例06.[gokit-playground-example](https://github.com/generals-space/gokit/tree/master/06.gokit-playground-example)是最完整的系列示例, 取自示例05.[gokit系列文章(译)](https://github.com/generals-space/gokit/tree/master/05.gokit%E7%B3%BB%E5%88%97%E6%96%87%E7%AB%A0(%E8%AF%91)), 但是添加了很多注释, 比较复杂的示例也做了拆分, 理解起来会更容易.

- 06.gokit-playground-example: **05.gokit系列文章(译)** 部分的示例代码
    - [01.gokit-lorem-restful](./06.gokit-playground-example/01.gokit-lorem-restful/readme): 按照 go-kit 模式提供的 lorem 服务.
    - [02.gokit-lorem-restful-ServerErrorEncoder](./06.gokit-playground-example/02.gokit-lorem-restful-ServerErrorEncoder/readme): 在[01.gokit-lorem-restful](./06.gokit-playground-example/01.gokit-lorem-restful/readme)的基础上, 添加对错误信息json化的操作.
    - [03.gokit-lorem-restful-ServerErrorLogger](./06.gokit-playground-example/03.gokit-lorem-restful-ServerErrorLogger/readme): 在[01.gokit-lorem-restful](./06.gokit-playground-example/01.gokit-lorem-restful/readme)的基础上, 添加对错误信息日志打印的操作.
    - [04.gokit-lorem-restful-client](./06.gokit-playground-example/04.gokit-lorem-restful-client/readme): 在[01.gokit-lorem-restful](./06.gokit-playground-example/01.gokit-lorem-restful/readme)的基础上, 添加 client 服务, 作为微服务中的一环进行通信.
    - [05.gokit-lorem-restful-lb](./06.gokit-playground-example/05.gokit-lorem-restful-lb/readme): 在[04.gokit-lorem-restful-client](./06.gokit-playground-example/04.gokit-lorem-restful-client/readme)的基础上, 对 server 服务做**负载均衡**, 但是没有使用**服务发现**, 在启动 client 服务时需要把 server 服务的服务列表传入.
    - [11.gokit-lorem-grpc](./06.gokit-playground-example/11.gokit-lorem-grpc/readme): 在[04.gokit-lorem-restful-client](./06.gokit-playground-example/04.gokit-lorem-restful-client/readme)的基础上, 进行 grpc 改造, client 与 server 间通信使用 grpc 而非 restful 形式.
    - [21.gokit-lorem-logging](./06.gokit-playground-example/21.gokit-lorem-logging/readme): 在[01.gokit-lorem-restful](./06.gokit-playground-example/01.gokit-lorem-restful/readme)的基础上, 添加 go-kit 内置的日志中间件. 
    - [31.gokit-lorem-ratelimit](./06.gokit-playground-example/31.gokit-lorem-ratelimit/readme): 在[01.gokit-lorem-restful](./06.gokit-playground-example/01.gokit-lorem-restful/readme)的基础上, 添加**限流**机制, 使用`golang.org/x/time/rate`限流器.
    - [32.gokit-lorem-ratelimit-juju](./06.gokit-playground-example/32.gokit-lorem-ratelimit-juju/readme): 在[01.gokit-lorem-restful](./06.gokit-playground-example/01.gokit-lorem-restful/readme)的基础上, 添加**限流**机制, 使用`juju/ratelimit`限流器.
    - [51.gokit-lorem-monitor](./06.gokit-playground-example/51.gokit-lorem-monitor/readme): 在[01.gokit-lorem-restful](./06.gokit-playground-example/01.gokit-lorem-restful/readme)的基础上, 添加**监控**操作, 使用`prometheus`进行日志埋点, 使用`grafana`查看.
    - [61.gokit-lorem-consul](./06.gokit-playground-example/61.gokit-lorem-consul/readme): 在[04.gokit-lorem-restful-client](./06.gokit-playground-example/04.gokit-lorem-restful-client/readme)的基础上, 使用`consul`中间件做**负载均衡**与**服务发现**.
    - [62.gokit-lorem-consul-client](./06.gokit-playground-example/62.gokit-lorem-consul-client/readme): 在[04.gokit-lorem-restful-client](./06.gokit-playground-example/04.gokit-lorem-restful-client/readme)的基础上, 使用`consul`中间件做**负载均衡**与**服务发现**. 在本例中, client 的角色类似于微服务中的 gateway, 将来自用户的请求转发给后端 server 服务.
    - [61.gokit-lorem-consul](./06.gokit-playground-example/61.gokit-lorem-consul/readme): 在[04.gokit-lorem-restful-client](./06.gokit-playground-example/04.gokit-lorem-restful-client/readme)的基础上, 使用`etcd`中间件做**负载均衡**与**服务发现**.
    - [71.gokit-lorem-hystrix](./06.gokit-playground-example/71.gokit-lorem-hystrix/readme): 在[04.gokit-lorem-restful-client](./06.gokit-playground-example/04.gokit-lorem-restful-client/readme)的基础上, 集成`hystrix-go/hystrix`断路器工具.

- 07.gokit链路追踪: **05.gokit系列文章(译)** 没有涉及到链路追踪的话题, 这里单独列出.
    - [01.gokit-lorem-single-tracing](./06.gokit-playground-example/01.gokit-lorem-single-tracing/readme): 在[04.gokit-lorem-restful-client](./06.gokit-playground-example/04.gokit-lorem-restful-client/readme)的基础上, 使用`zipkin`中间件添加**链路**追踪机制.
    - [02.gokit-lorem-tracing-consul](./06.gokit-playground-example/02.gokit-lorem-tracing-consul/readme): 在上例基础上, 继续添加 consul 完成**负载均衡**与**服务发现**功能.

## 历史记录

Endpoint类似于web服务的url接口, 一个endpoint表示一个路由接口.

但是Endpoint不能直接对外提供服务, 需要通过Transport转换成http/grpc类型的接口才可以.

2018-09-06

在完成第4个demo时意识到go-kit貌似什么都没有做, 没有gateway(入口路由服务), 微服务间通信为原生grpc, http处理为原生net/http库. ta就像js里的`backbone`, 定义好一些核心概念, 组件如何定义, 组件间如何配合, 数据接口如何定义等完全不管...

只有一点, 通过`makeXXXEndpoint`及`encode`, `decode`几个函数后, 我们写的业务逻辑可以同时生成http/grpc两种接口, 只要选择不同的transport进行`NewServer`操作即可, 但这`grpc-ecosystem`包也能完成.

我已实在想不出继续用ta的理由了, 接下来尝试ta的middleware和监控什么的, 但下一步的重点会放在`go-micro`上.
