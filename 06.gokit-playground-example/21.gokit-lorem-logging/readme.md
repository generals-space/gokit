此工程为[ru-rocker/gokit-playground](https://github.com/ru-rocker/gokit-playground/tree/master/lorem-logging)的精简版, 添加部分注释.

这个工程其实就是restful示例, 添加了一个用于业务逻辑中Service对象的装饰器, 在每次执行业务逻辑前后打印一些消息. 在`cmd/main.go`中有两种输出方式. 一种是将日志输出到文件(只有接口日志, 没有服务启动等其他信息), 一种是打印到标准输出, 将随服务启动日志一起打印.

所有操作都与[01.gokit-lorem-restful]()完全相同.

启动服务

```console
$ go run cmd/main.go
Starting server at port 8080
```

对其发起 http 请求进行测试.

```console
$ curl -XPOST localhost:8080/lorem/word/1/20
{"message":"exclamaverunt"}
$ curl -XPOST localhost:8080/lorem/sentence/1/20
{"message":"Cura ob pro qui."}
$ curl -XPOST localhost:8080/lorem/paragraph/1/20
{"message":"Deo vos satietate retenta instat te en igitur aequo tibi. Contexo pro peregrinorum, heremo absconditi araneae meminerim deliciosas actionibus, facere modico dura sonuerunt psalmi contra rerum tempus. Agit cadere o mole te necessaria sonuerit nomen nam da. Se talibus re pro me audio cuius aula me se coepta se lux res inter motus. E tua qui video te, psalmi agam me mea pro da dici sentio tradidisti ipsa. Ita praeire nescio faciant notiones proceditur paucis te da fortitudinem ubique ne pro. Grex e o, recorder cor re, libidine assecutus aderat per profunda. Cura hi deo agnoscerent retractanda hi, extingueris sacramenta abundare agam ante dixi habitaculum absunt genus pulchritudo. Da rei ne cui iam. Nam acceptam sed e a corruptione re tot cui tradidisti redemit gero ore affectu. Amem genus odor pecco re quiescente occursantur semper agam curiositate ore fallere colunt sapores fidei maneas saepe corporis re oris. Rogo lacrimas ea crucis sanaturi melior via. Aut saucio odorem pro angustus fine tua e die potu recti at est ab. Me etiamne nam da vi, vide vestra eum agam loca noe ob an. Desperatione sed re ore rei reficiatur o te latet tam benedicendo partes vi vehementer maerere corones abigo ea omni timeo meos."}
```

server 服务将有如下日志输出

```console
$ go run cmd/main.go
Starting server at port 8080
ts=2020-05-20T13:12:51.976413033Z caller=middle.go:27 type=word min=1 max=20 result=exclamaverunt took=4.134µs
ts=2020-05-20T13:12:59.022481286Z caller=middle.go:27 type=sentence min=1 max=20 result="Cura ob pro qui." took=9.26µs
ts=2020-05-20T13:13:05.058214391Z caller=middle.go:27 type=paragraph min=1 max=20 result="Deo vos satietate retenta instat te en igitur aequo tibi. Contexo pro peregrinorum, heremo absconditi araneae meminerim deliciosas actionibus, facere modico dura sonuerunt psalmi contra rerum tempus. Agit cadere o mole te necessaria sonuerit nomen nam da. Se talibus re pro me audio cuius aula me se coepta se lux res inter motus. E tua qui video te, psalmi agam me mea pro da dici sentio tradidisti ipsa. Ita praeire nescio faciant notiones proceditur paucis te da fortitudinem ubique ne pro. Grex e o, recorder cor re, libidine assecutus aderat per profunda. Cura hi deo agnoscerent retractanda hi, extingueris sacramenta abundare agam ante dixi habitaculum absunt genus pulchritudo. Da rei ne cui iam. Nam acceptam sed e a corruptione re tot cui tradidisti redemit gero ore affectu. Amem genus odor pecco re quiescente occursantur semper agam curiositate ore fallere colunt sapores fidei maneas saepe corporis re oris. Rogo lacrimas ea crucis sanaturi melior via. Aut saucio odorem pro angustus fine tua e die potu recti at est ab. Me etiamne nam da vi, vide vestra eum agam loca noe ob an. Desperatione sed re ore rei reficiatur o te latet tam benedicendo partes vi vehementer maerere corones abigo ea omni timeo meos." took=450.623µs
```
