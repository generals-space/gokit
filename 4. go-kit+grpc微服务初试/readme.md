go-kit是什么这里就不介绍了, 就算写也不如网上的文章写的正规. 这里简单记录一下go-kit中提供的3个核心概念

1. Service

2. Endpoint

3. Transport

go-kit官方示例`stringsvc1`中对这三个概念表现的比较得...一般. 

`Service`即是我们的业务逻辑, 为了得到清晰的分层架构, `Service`独立与另外两者, 且划分地很明确.(只不过`stringsvc1`用了接口表示, 感觉多此一举)

`Endpoint`是项目内部微服务之间通信的接口, 每个endpoint相当于protobuf中一个rpc字段. 在`stringsvc1`中并没有内部子服务间的交互, 而是直接把`Service`暴露为http接口了, 所以这部分的作用没有体现出来.

`Transport`是go-kit为我们提供的, 暴露给外部访问的入口. 有grpc, 也有restful类型. 示例中为restful, 使用curl访问即可.

这里说一下, `stringsvc1`中把`Service`包装为`Endpoint`类型的部分, 也就是`makeXXXEndpoint`函数, 是不是和grpc中的服务端代码很像? 不过这里没有用protobuf, 而是直接定义了`uppercaseRequest`, `uppercaseResponse`这种结构体. 从`Transport`访问到`Service`过程中, 有参数与响应的转换操作. 

在我们的示例中, 把这部分用gprc + protobuf代替, 同时把客户端操作抽象成另一个子服务, 来体现`Endpoint`的作用.
