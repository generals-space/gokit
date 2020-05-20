在原作者的[gokit-playground]()示例中, `httptransport.ServerOption{}`包含了两个额外选项. [01.gokit-lorem-restful]()为了精简, 将这两个选项移除了. 此工程就是为了实验其中一个选项`ServerErrorEncoder()`的用法和作用. 

常规操作下, endpoint函数可以返回的正确的结果. 如下在某些场景下, 你想返回一个error(要看你的具体业务而定), 如果不经处理, 会得到如下输出

```
$ curl -XPOST localhost:8080/lorem/sentence/1/20
test error
```

你会发现返回的错误并没有经过`encodeResponse()`函数, 返回的是纯文本, 与正常的响应格式也不匹配. 

我们希望得到一个json响应, 通过`error`字段展示错误信息.

启动本工程, 再通过curl访问, 就可以得到这样的输出(所有操作都与[01.gokit-lorem-restful]()完全相同).

```
$ curl -XPOST localhost:8080/lorem/sentence/1/20
{"error":"test error"} 
```

> 本示例中只能返回`error`, 无法正确生成`word`, `sentence`, `paragraph`的正确文字.
