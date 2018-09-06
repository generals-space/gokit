
1. [go-kit offical example](https://github.com/go-kit/kit/tree/master/examples)

2. [一、go 语言编写 grpc 微服务实例](http://www.articlechain.cn/articles/2018/04/27/1524816230413)
    - 系列文章

3. [Micro-services Using Go-kit: REST Endpoint](http://www.ru-rocker.com/2017/02/17/micro-services-using-go-kit-http-endpoint/)

go-kit的官方示例有点过度设计的感觉, 尤其是`领域驱动设计`的思想让人无法理解.

参考文章2中的示例简单一些, 而且条理更清晰, 但在设计上与我目前遇到的场景不匹配.

参考文章3只有一个示例, 太简单了.

在这个示例库, 从一个简单的示例, 逐渐加深, 层层拆分, 希望能够更清晰的阐述微服务的架构和思想.

## 历史记录

2018-09-06

在完成第4个demo时意识到go-kit貌似什么都没有做, 没有gateway(入口路由服务), 微服务间通信为原生grpc, http处理为原生net/http库. ta就像js里的`backbone`, 定义好一些核心概念, 组件如何定义, 组件间如何配合, 数据接口如何定义等完全不管...

只有一点, 通过`makeXXXEndpoint`及`encode`, `decode`几个函数后, 我们写的业务逻辑可以同时生成http/grpc两种接口, 只要选择不同的transport进行`NewServer`操作即可, 但这`grpc-ecosystem`包也能完成.

我已实在想不出继续用ta的理由了, 接下来尝试ta的middleware和监控什么的, 但下一步的重点会放在`go-micro`上.