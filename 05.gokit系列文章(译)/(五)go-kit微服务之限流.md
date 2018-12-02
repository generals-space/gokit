# (五)go-kit微服务之限流

原文链接

1. [Micro-services Using go-kit: Rate Limiting](http://www.ru-rocker.com/2017/03/19/micro-services-using-go-kit-rate-limiting/)

## 1. 概述

在前面的文章中, 我们在原来的随机文本服务中添加了日志功能, 但是对于生产环境还是不够可靠. 我们仍然需要在服务中继续添加中间件以增强服务能力.

所以, 在这篇文章中, 我将介绍一种中间件功能, 类似于仪表控制(真别扭, 没读懂). 我将在中间件中实现限流的功能.

## 2. 限流服务

在微服务的世界里, 我们经常需要对接口进行请求次数的限制, 以保护服务不会过载. 因为有时一个请求可能耗费大量CPU(或是内存)资源. 单位时候内, 大量的访问请求可能会导致服务性能的下降.

### 2.1 Token Bucket Limiter(令牌桶)

为了实现限流的功能, 我需要添加一种基于token的限流算法. 简单来说, 我们拥有一个令牌桶(其中有一定数量的令牌), 每个请求都会从中取得一个令牌, 以便处理接下来的逻辑. 如果桶中已经没有令牌, 那么此次请求便无法完成. 当前令牌桶在一定时间内会重新充满. 关于此次算法的更多令牌, 可以参考[这里](https://en.wikipedia.org/wiki/Token_bucket)

在golang中, 有一个库[juju](https://github.com/juju/ratelimit)实现了这个算法. 此外go-kit有一个内置的中间件也实现了这个算法, 这将极大简化我们的工作.

### 2.2 go-kit中间件

在go-kit中, 应用于endpoint的中间件是一个函数的别名, 输入和输出都是`endpoint.Endpoint`类型.

```go
# Go-kit Middleware Endpoint
type Middleware func(Endpoint) Endpoint
```

...而且, `endpoint.Endpoint`本身也是函数类型.

```go
# Go-kit Endpoint Function
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
```

**一步一步来**

好了, 现在我们开始着手写代码. 我要告诉你, 其实不必在原来代码上做太大改动就可以完成, 是不是很厉害? 那在这篇文章里, 原有的代码就是前面我们讲的, 带有日志功能的工程`lorem-logging`, 我们要为ta添加增强型功能: 限流. 拷贝一份, 然后命名为`lorem-rate-limit`, 仍然叫做工作目录

首先下载依赖库juju

```
# juju library
go get github.com/juju/ratelimit
```

### 2.3 `instrument.go`

在已有的工作目录下创建`instrument.go`文件, 然后添加名为`NewTokenBucketLimiter`的函数, 这个函数接受一个**令牌桶对象**作为参数, 然后返回`endpoint.Endpoint`类型对象. 注意一点: 在调用下一个endpoint(next())之前, 我们需要用`TakeAvailable`函数检查token是否可用.

```go
var ErrLimitExceed = errors.New("Rate Limit Exceed")
func NewTokenBucketLimiter(tb *ratelimit.Bucket) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			if tb.TakeAvailable(1) == 0 {
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}
```

### 2.4 `main.go`

在这一步, 我们需要定义一个初始化**限流令牌桶**的函数`NewBucket`. 本示例中, 每秒都会重新填满令牌桶, 桶中最多有5个令牌. 注意哦, 生产环境中可不要定义这么小, 这里只是一个演示而已. 接下来为原有的endpoint添加中间件函数.

```go
func main() {
	ctx := context.Background()
	errChan := make(chan error)

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var svc lorem_rate_limit.Service
	svc = lorem_rate_limit.LoremService{}
	svc = lorem_rate_limit.LoggingMiddleware(logger)(svc)

    // 看这里
	rlbucket := ratelimit.NewBucket(1*time.Second, 5)
    e := lorem_rate_limit.MakeLoremLoggingEndpoint(svc)
    // 还有这里
	e = lorem_rate_limit.NewTokenBucketLimiter(rlbucket)(e)
	endpoint := lorem_rate_limit.Endpoints{
		LoremEndpoint: e,
	}

	r := lorem_rate_limit.MakeHttpHandler(ctx, endpoint, logger)

	// HTTP transport
	go func() {
		fmt.Println("Starting server at port 8080")
		handler := r
		errChan <- http.ListenAndServe(":8080", handler)
	}()


	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<- errChan)
}
```

### 2.5 运行, 测试

就这么几步, 活已经干完了, 是不是很简单? 现在我们来测试一下. 启动服务端, 然后发几个请求, 如果过于频繁的话, 会返回错误响应, 告诉你`Rate Limit Exceed(速度受限)`

```
# output logs
ts=2017-03-19T04:33:57.97656492Z caller=logging.go:31 function=Word min=20 max=20 result=persentiscere took=3.417µs
ts=2017-03-19T04:33:58.123130597Z caller=logging.go:31 function=Word min=20 max=20 result=exclamaverunt took=2.678µs
ts=2017-03-19T04:33:58.258280166Z caller=logging.go:31 function=Word min=20 max=20 result=difficultates took=2.934µs
ts=2017-03-19T04:33:58.84378762Z caller=logging.go:31 function=Word min=20 max=20 result=recognoscimus took=3.47µs
ts=2017-03-19T04:33:59.338875968Z caller=logging.go:31 function=Word min=20 max=20 result=cognosceremus took=4.142µs
ts=2017-03-19T04:33:59.755599747Z caller=server.go:110 err="Rate Limit Exceed"
ts=2017-03-19T04:33:59.923025144Z caller=server.go:110 err="Rate Limit Exceed"
ts=2017-03-19T04:34:00.086307562Z caller=logging.go:31 function=Word min=20 max=20 result=similitudinem took=3.108µs
ts=2017-03-19T04:34:00.224307681Z caller=server.go:110 err="Rate Limit Exceed"
```

### 2.6 (可选)

可能当请求过于频繁时, 服务端直接返回错误会让人感觉很不爽, 因为一点也不优雅. 好一点的作法是, 让请求等待一会, 直到桶中有新的token可用, 再继续完成之后的流程.(译者:...我怎么感觉直接返回错误更好呢). 为了实现这个想法, go-kit提供了`NewTokenBucketThrottler`中间件, 我们需要修改一行代码

```go
// add import: ratelimitkit "github.com/go-kit/kit/ratelimit"

// replace this line
// e = lorem_rate_limit.NewTokenBucketLimiter(rlbucket)(e)
// with 
e = ratelimitkit.NewTokenBucketThrottler(rlbucket, time.Sleep)(e)
```

## 3. 总结

创建一个服务, 无论何时都不要忘了做请求次数限制. 尤其是每次请求都会耗费大量CPU/内存资源时. 这样才能保证你的服务性能表现更为良好(译者: 我的确遇到请求过多导致服务崩溃的情况).

不过, 我仍然对go-kit的表现十分惊讶. ta把这些做得十分方便, 模块化, 让我们无论何时想要添加一个新的特性都能工作得很好.

本文所用源码在[这里](https://github.com/ru-rocker/gokit-playground/tree/master/lorem), 本系列文章的所有示例都基于这个服务.

## 4. 参考

Go Programming Blueprints – Second Edition by Mat Ryer
