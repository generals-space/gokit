此工程为[ru-rocker/gokit-playground](https://github.com/ru-rocker/gokit-playground/tree/master/lorem-rate-limit)的精简版.

在原文[Micro-services Using go-kit: Rate Limiting](http://www.ru-rocker.com/2017/03/19/micro-services-using-go-kit-rate-limiting/)中提到了两种可引用的令牌桶算法.

一种是[Juju](https://github.com/juju/ratelimit), 另一种说是go-kit内置的中间件.

在示例[31.gokit-lorem-ratelimit]()中我只使用了`golang/x/time/rate`库, 而在本例中替换为了`Juju`, 作用相同.
