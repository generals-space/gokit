# 07.gokit链路追踪

参考文章

1. [go-kit微服务：服务链路追踪](https://juejin.im/post/5c77bb8b6fb9a04a027b0b99)
2. [六、go-kit 微服务请求跟踪介绍](https://hacpai.com/article/1525401758789)

本节的内容是我在阅读完示例05[gokit系列文章(译)](https://github.com/generals-space/gokit/tree/master/05.gokit%E7%B3%BB%E5%88%97%E6%96%87%E7%AB%A0(%E8%AF%91))并编写完示例06[gokit-playground-example](https://github.com/generals-space/gokit/tree/master/06.gokit-playground-example)后开始着手写的, 主要目的是有一个实际的项目在压测时性能不给力, 我需要知道在各子服务之间哪一步消耗时间较多, 由于找到了微服务架构中的"服务链路追踪"专题.

业务逻辑依然延用示例05中的lorem服务. 从项目架构上来说, 与示例06中的工程同级, 只是由于不再属于系列文章中的模块, 才单独拎出来的.

