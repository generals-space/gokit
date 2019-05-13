此工程为[ru-rocker/gokit-playground](https://github.com/ru-rocker/gokit-playground/tree/master/lorem)的精简版, 移除了日志, 及错误处理部分的代码, 将原来业务逻辑的3个方法合并成了1个, 与后面用到的示例保持统一, 并添加了详细注释.

首先构建项目的依赖镜像(只是下载项目依赖包)

```
docker build --no-cache=true -f dep.dockerfile -t gokittestdep .
```

然后通过`docker-compose`启动.

```
docker-componse up -d
```

由于是开发环境, 所以将项目路径挂载到容器中, 修改代码后重启服务就可以看到效果, 不用重新构建.

项目启动后, 正常的访问结果如下

```
$ curl -XPOST localhost:8080/lorem/word/1/20
{"message":"difficultates"}

$ curl -XPOST localhost:8080/lorem/sentence/1/20
{"message":"Concurrunt nota re dicam fias sim aut pecco die appetitum ea mortalitatis hi."}

$ curl -XPOST localhost:8080/lorem/paragraph/1/20
{"message":"Tibi ita recedimus an aut eum tenacius quae mortalitatis eram aut rapit montium inaequaliter dulcedo. Contra rerum tempus mala, anima volebant dura quae o. Sonuerit nomen nam da nuntii. Talibus re pro me audio. Deum temptatione imperas da vi, an da cuius facere valeam e tua qui video te psalmi agam. Me indicabo te tuetur audi. Mirabilia amor primus aboleatur, te, meque mundatior deserens da contexo e suaveolentiam. Aut ita sensarum nuda eripietur superbam isto ab, sana tu ita ore siderum. Lux horum an ore nam, dicens ore curiosarum filiorum eruuntur, munerum displicens. Temptationem cor plena modi agito, inlusio, deo fama propterea ab persentiscere nam acceptam sed e a corruptione re. Ea nascendo qui fuit ceterarumque me odorem amem genus odor, pecco re quiescente occursantur semper."}
```
