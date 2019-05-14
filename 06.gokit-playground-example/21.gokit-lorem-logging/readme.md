此工程为[ru-rocker/gokit-playground](https://github.com/ru-rocker/gokit-playground/tree/master/lorem-logging)的精简版, 添加部分注释.

这个工程其实就是restful示例, 添加了一个用于业务逻辑中Service对象的装饰器, 在每次执行业务逻辑前后打印一些消息. 在`cmd/main.go`中有两种输出方式. 一种是将日志输出到文件(只有接口日志, 没有服务启动等其他信息), 一种是打印到标准输出, 将随服务启动日志一起打印.

所有操作都与[01.gokit-lorem-restful]()完全相同.
