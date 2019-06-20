本示例基于[daizuozhuo/etcd-service-discovery](https://github.com/daizuozhuo/etcd-service-discovery), 但ta依赖的etcd客户端是[coreos/etcd](https://github.com/coreos/etcd/client), 这个项目已经迁移至[etcd/clientv3](https://github.com/etcd-io/etcd), 本例对ta做了基本的改造.

安装好依赖后, 先启动服务端, 然后运行客户端.

在客户端运行期间, 服务端会输出

```
$ go run server/main.go
2019-06-21 01:21:58.720398 I | Add worker workers/node-01: node-01
2019-06-21 01:22:01.733234 I | Update worker workers/node-01: node-01
2019-06-21 01:22:04.741934 I | Update worker workers/node-01: localhost
```

停止客户端, 过TTL秒后(TTL是代码中的一个变量), 服务端会输出

```
2019-06-21 01:23:28.970627 I | Delete worker  workers/node-01
```

在这个示例中, server实现了服务注册与服务发现的基本机制, client循环执行心跳操作, 其实就是不断刷新etcd中的key, server将client注册到etcd的key保存在自己的members列表中.

由于client端在注册时附带了TTL, 存储在etcd中的key也是有过期时间的, server监听这个key的变化, 当client终止, key由于过期而被删除, server就将该client从members列表中移除.
