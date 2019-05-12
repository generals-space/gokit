
1. [go-kit offical example](https://github.com/go-kit/kit/tree/master/examples)

2. [一、go 语言编写 grpc 微服务实例](https://hacpai.com/article/1524816248447)
    - 系列文章

3. [Micro-services Using Go-kit: REST Endpoint](http://www.ru-rocker.com/2017/02/17/micro-services-using-go-kit-http-endpoint/)
    - 系列文章

go-kit的官方示例有点过度设计的感觉, 尤其是`领域驱动设计`的思想让人无法理解.

参考文章2中的示例简单一些, 而且条理更清晰, 但在设计上与我目前遇到的场景不匹配.

参考文章3是国外的一篇系列文章, 个人感觉写的比参考文章2要好, 且更详细规范(参考文章2的示例有很多都没有使用go-kit组件).

在这个示例库, 从一个简单的示例, 逐渐加深, 层层拆分, 希望能够更清晰的阐述微服务的架构和思想.

## 历史记录

Endpoint类似于web服务的url接口, 一个endpoint表示一个路由接口.

但是Endpoint不能直接对外提供服务, 需要通过Transport转换成http/grpc类型的接口才可以.

2018-09-06

在完成第4个demo时意识到go-kit貌似什么都没有做, 没有gateway(入口路由服务), 微服务间通信为原生grpc, http处理为原生net/http库. ta就像js里的`backbone`, 定义好一些核心概念, 组件如何定义, 组件间如何配合, 数据接口如何定义等完全不管...

只有一点, 通过`makeXXXEndpoint`及`encode`, `decode`几个函数后, 我们写的业务逻辑可以同时生成http/grpc两种接口, 只要选择不同的transport进行`NewServer`操作即可, 但这`grpc-ecosystem`包也能完成.

我已实在想不出继续用ta的理由了, 接下来尝试ta的middleware和监控什么的, 但下一步的重点会放在`go-micro`上.