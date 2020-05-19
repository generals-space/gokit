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

## 历史记录

Endpoint类似于web服务的url接口, 一个endpoint表示一个路由接口.

但是Endpoint不能直接对外提供服务, 需要通过Transport转换成http/grpc类型的接口才可以.

2018-09-06

在完成第4个demo时意识到go-kit貌似什么都没有做, 没有gateway(入口路由服务), 微服务间通信为原生grpc, http处理为原生net/http库. ta就像js里的`backbone`, 定义好一些核心概念, 组件如何定义, 组件间如何配合, 数据接口如何定义等完全不管...

只有一点, 通过`makeXXXEndpoint`及`encode`, `decode`几个函数后, 我们写的业务逻辑可以同时生成http/grpc两种接口, 只要选择不同的transport进行`NewServer`操作即可, 但这`grpc-ecosystem`包也能完成.

我已实在想不出继续用ta的理由了, 接下来尝试ta的middleware和监控什么的, 但下一步的重点会放在`go-micro`上.
