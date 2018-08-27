gprc层本身与业务逻辑服务就是分开实现的.

我们需要把grpc传入的protobuf参数转化为业务函数需要的格式, 同时也需要把业务层的响应转化为grpc的protobuf结果.

`rpc_server`中, `xxxServiceServer`部分的结构体一般都可为空, 与实际的业务服务分离. 因为`xxxService`在`protobuf`文件中只定义了`rpc`字段, 也就是要实现的方法, 而没有定义成员变量.

在业务逻辑与rpc层分离的架构中, 可以在rpc server启动前, 先启动业务服务. 然后`xxxServiceServer`的方法都去调用业务服务暴露出来的方法. 这样做的缺点就是之前所说的, 多了请求与响应之间的转换层. 这个转换层大概就放在`xxxServiceServer`的实现的`rpc`成员方法中.

当然, 也可以将业务逻辑与rpc的`xxxServiceServer`结合, 但我觉得这样的架构未免不清晰...也可能是我过度设计了.
