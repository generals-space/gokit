在原作者的[gokit-playground]()示例中, `httptransport.ServerOption{}`包含了两个额外选项. [01.gokit-lorem-restful]()为了精简, 将这两个选项移除了. 

此工程就是为了实验其中一个选项`ServerErrorEncoder()`的用法和作用. 

所有操作都与[01.gokit-lorem-restful]()完全相同.

常规操作下, endpoint函数返回的正确的结果, 如果你尝试返回一个error, 会得到如下输出

```
$ curl -XPOST localhost:8080/lorem/sentence/1/20
test error
```

你会发现返回的错误并没有经过`encodeResponse()`函数, 而且这是纯文本, 与正常的响应格式也不匹配. 

我们希望得到一个json响应, 通过`error`字段展示错误信息.

启动本工程, 再通过curl访问, 就可以得到这样的输出.

```
$ curl -XPOST localhost:8080/lorem/sentence/1/20
{"error":"test error"} 
```
