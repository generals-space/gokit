# (一)使用go-kit构建rest接口服务

原文链接

1. [Micro-services Using Go-kit: REST Endpoint]([图片]http://www.ru-rocker.com/2017/02/17/micro-services-using-go-kit-http-endpoint/)

> 在这篇文章中, 我会创建一个golang版的简单的微服务, 使用go-kit作为标准的微服务工具库, 这个服务可以提供restful的API.

## 1. 微服务

在当下的编程圈, 微服务架构已经得到了普及. 我并不会专门介绍微服务这种架构, 因为网上已经有太多的人在讨论. 不过我还是推荐两个关于微服务的好网站.

1. [Martin Fowler]([图片]https://martinfowler.com/articles/microservices.html)

2. [[图片]microservices.io]([图片]http://microservices.io/)

在里面有大量的优秀文章在讨论模式(pattern)和参考示例.

### 1.1 golang

略过...

### 1.2 go-kit

go-kit在简化微服务架构的具体实现上有很大帮助, 因为ta有相当多的功能组件比如服务连接(service connectivity)、度量(metrics)和日志(logging). 所以我非常感谢`[Peter Bourgon](@peterbourgon)`和go-kit的所有贡献者, 开发出了这样一个优秀的库.

## 2. 示例场景

在我们的示例场景中, 将会以一个服务的形式提供一个restful的接口, 可以通过POST的方式访问, 然后服务将会返回一条json格式的`lorem ipsum`消息.

请求url格式为`/lorem/{type}/{min}/{max}`, 各参数涵义为:

1. `type`为`lorem`类型, 可选值为`word(单词)`, `sentence(句子)`和`paragraph(段落)`.

2. `min`和`max`表示生成的最少或最多的字母限制.

**一步一步来**

在开始之前, 需要先安装几个依赖库.

1. [go-kit](https://github.com/go-kit/kit)

2. [golorem](https://github.com/drhodes/golorem)

3. [gorilla mux](https://github.com/gorilla/mux)

```
go get github.com/go-kit/kit
go get github.com/drhodes/golorem
go get github.com/gorilla/mux
```

> 译者注: 没搞明白`lorem ipsum`有什么靠谱点的翻译, 可以看作是一段随机生成的文本, 之后也以**随机文本**代称了, 由`golorem`包提供实现. 这里定义的接口比如`/lorem/word/1/50`会随机生成一个最少1个, 最多50个字母的单词, 而`/lorem/sentence/1/50`会随机生成一句最少1个, 最多50个单词的句子, 同理`/lorem/paragraph/1/50`则会随便生成一段最少1个, 最多50个句子的段落.

### 2.1 构建服务(业务层)

不管你用了多么高大上的工具库, 我们都需要先创建业务逻辑. 正如上面描述过的使用场景, 我们的业务逻辑就是, 基于`word(单词)`, `sentence(句子)`或`paragraph(段落)`生成随机文本.

ok, 我们先创建我们的工程目录, 在这个示例中为`$GOPATH/github.com/ru-rocker/gokit-playground/lorem`, 然后在此目录下创建`service.go`, 内容如下

```go
// 定义服务接口...对于有相似类型, 不同实现的服务会方便一些. 
type Service interface {
	Word(min, max int) string

	Sentence(min, max int) string

	Paragraph(min, max int) string
}

// Implement service with empty struct
type LoremService struct {

}
```

定义了接口以后, 再添加如下实现代码.

```go
// Implement service functions
func (LoremService) Word(min, max int) string {
	return golorem.Word(min, max)
}

func (LoremService) Sentence(min, max int) string {
	return golorem.Sentence(min, max)
}

func (LoremService) Paragraph(min, max int) string {
	return golorem.Paragraph(min, max)
}
```

> 咳, 其实就是直接调用的`golorem`库的方法, 可以看一下`golorem`的`readme`文档.

> 译者注: 实际场景中, `XXXService`结构中可包含数据库连接`db`对象, 日志`logger`对象, 其他服务的`grpc`连接对象等成员属性, 这样就可以在成员方法中使用这些对象做些实质性的操作了.

### 2.2 规划请求与响应结构

因为此服务最终要提供一个http接口, 因此下一步就是要规划http请求与响应的数据结构. 在我们的场景中, url包含了3个属性: `type`, `min`和`max`.

对于响应, 只需要两个字段, 一个表示错误信息(无错误时隐藏), 另一个表示实际生成的`lorem ipsum`文本.

我们再创建一个`endpoints.go`文件, 内容如下

```go
//request
type LoremRequest struct {
	RequestType string
	Min int
	Max int
}

//response
type LoremResponse struct {
	Message string `json:"message"`
	Err     error `json:"err,omitempty"` //omitempty means, if the value is nil then this field won't be displayed
}
```

### 2.3 创建接口(endpoint)

在go-kit中, `Endpoint`是一种特殊的函数, 你可以把ta包装成`http.Handler`以处理http请求(包装的操作由transport完成). 

为了先让我们的服务函数转换成`endpoint.Endpoint`函数, 我们要创建一个工厂函数(在go-kit概念中为`MakeXXXEndpoint()`类型的函数). 在工厂函数中创建的`Endpoint`函数接受`LoremRequest`对象, 并根据其中的`type`成员决定调用业务服务中的哪个方法(译者注: 在此之前业务对象已经实例化, `MakeXXXEndpoint()`接受的就是业务对象). 然后将执行结果包装成`LoremResponse`再返回.

在`endpoints.go`中追加如下内容

```go
var (
	ErrRequestTypeNotFound = errors.New("Request type only valid for word, sentence and paragraph")
)

// endpoints wrapper
type Endpoints struct {
	LoremEndpoint endpoint.Endpoint
}

// creating Lorem Ipsum Endpoint
func MakeLoremEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoremRequest)

		var (
			txt string
			min, max int
		)

		min = req.Min
		max = req.Max

		if strings.EqualFold(req.RequestType, "Word") {
			txt = svc.Word(min, max)
		} else if strings.EqualFold(req.RequestType, "Sentence"){
			txt = svc.Sentence(min, max)
		} else if strings.EqualFold(req.RequestType, "Paragraph") {
			txt = svc.Paragraph(min, max)
		} else {
			return nil, ErrRequestTypeNotFound
		}

		return LoremResponse{Message: txt}, nil
	}

}
```

### 2.4 Transport

在处理http请求与响应前, 我们需要创建一个`encoder`和`decoder`, 以便将`LoremResponse`响应转换成json, 或者将请求体中的json数据转换成`LoremRequest`对象.

为了实现这样的功能, 我们需要再创建一个独立的文件`transport.go`, 相关代码如下

```go
// decode url path variables into request
func decodeLoremRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrBadRouting
	}

	vmin, ok := vars["min"]
	if !ok {
		return nil, ErrBadRouting
	}

	vmax, ok := vars["max"]
	if !ok {
		return nil, ErrBadRouting
	}

	min, _ := strconv.Atoi(vmin)
	max, _ := strconv.Atoi(vmax)
	return LoremRequest{
		RequestType: requestType,
		Min: min,
		Max: max,
	}, nil
}

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error.
type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encode error
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
```

在创建了`encoder`和`decoder`函数后, 就需要创建`http handler`了. 在MVC架构应该被称为`路由-控制器`映射). 

在`transport.go`中追加如下内容

```go
// 创建路由映射, 控制器为go-kit提供的httptransport的实例, 注意其参数的用法.
func MakeHttpHandler(ctx context.Context, endpoint Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	//POST /lorem/{type}/{min}/{max}
	r.Methods("POST").Path("/lorem/{type}/{min}/{max}").Handler(httptransport.NewServer(
		ctx,
		endpoint.LoremEndpoint,
		decodeLoremRequest,
		encodeResponse,
		options...,
	))
	return r
}
```

### 2.5 Main入口

写到这里, 我们的服务已经有了服务层(业务层), endpoint和transport. 所有准备工作都做好后, 我们开始创建主函数.

在`lorem`目录下, 创建一个子目录`lorem.d`, `.d`意味`daemon(守护进程, 后台服务)`, 当然, 其实名称随你喜欢. 然后在其下创建`main.go`, 主要内容如下

```go
func main() {
	ctx := context.Background()
	errChan := make(chan error)

    // 实例化业务服务, 实际场景中该是为我们的服务读取配置文件, 连接数据库, 创建日志对象的操作
	var svc lorem.Service
    svc = lorem.LoremService{}
    
    // 实例化Endpoints集合
	endpoint := lorem.Endpoints{
		LoremEndpoint: lorem.MakeLoremEndpoint(svc),
	}

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}
    // 创建handler, 其实就像是注册路由.
	r := lorem.MakeHttpHandler(ctx, endpoint, logger)

	// HTTP transport
	go func() {
		fmt.Println("Starting server at port 8080")
		handler := r
		errChan <- http.ListenAndServe(":8080", handler)
	}()

    // 主进程阻塞并监听用户输入, `ctrl+C`退出
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<- errChan)
}
```

### 2.6 运行示例

```go
cd $GOPATH
go run src/github.com/ru-rocker/gokit-playground/lorem/lorem.d/main.go
```

使用curl或postman按照`http://localhost:8080/lorem/word/1/50`的格式发送请求.

```
http://localhost:8080/lorem/word/1/50
{
"message": "suggestionum"
}
http://localhost:8080/lorem/sentence/1/50
{
"message": "Desperatione sed re ore rei reficiatur o."
}
http://localhost:8080/lorem/paragraph/1/50
{
"message": "Dicturus ita mediator ita mundum lux partes miseros percepta seu dicant avaritiam nares contra deseri, securus te. Sobrios tale rogo sanctis rerum multis teneam hi vos languor me victor suis ea si asperum ore ob audi. Contrario hi cogo ea rogo cor convinci o. Deum solebat ipsa olefac ridentem nam debeo frigidique memoriae atque nuntii. Re te da, actionibus os. Consuevit antiqua tot sed, petimus fugam solus te eodem adparet coniunctam aut. Hanc tibi freni fulgeat re fit ipsa contexo introrsus a cogit videtur tot exciderunt sat valerem. Praeditum valetudinis aer videant vi rei muta. Ita ea diligi caro psalmi e inaequaliter dicis ab. Mei da o lucem trahunt quae. Os da me sapientiae pro."
}
```

### 2.7 源码

本文所用源码在[这里](https://github.com/ru-rocker/gokit-playground/tree/master/lorem), 本系列文章的所有示例都基于这个服务.