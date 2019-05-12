在原作者的[gokit-playground]()示例中, `httptransport.ServerOption{}`包含了两个额外选项. [01.gokit-lorem-restful]()为了精简, 将这两个选项移除了. 

此工程就是为了实验其中另一个选项`ServerErrorLogger()`的用法和作用.

所有操作都与[01.gokit-lorem-restful]()完全相同.

正如其名称`ServerErrorLogger`, 这个选项只在endpoint函数返回error时有效, ta将记录出错的部分及错误信息.

我们让endpoint返回error, 访问时服务端日志将会如下

```
Starting server at port 8080
ts=2019-05-12T12:47:42.9487334Z caller=error_handler.go:27 err="test error"
```
