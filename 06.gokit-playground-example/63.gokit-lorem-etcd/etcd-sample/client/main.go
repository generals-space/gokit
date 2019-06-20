package main

import (
	discovery "github.com/generals-space/gokit/06.gokit-playground-example/63.gokit-lorem-etcd/etcd-sample"
)

func main() {
	worker := discovery.NewWorker("node-01", "127.0.0.1", discovery.Endpoints)
	worker.HeartBeat()
}
