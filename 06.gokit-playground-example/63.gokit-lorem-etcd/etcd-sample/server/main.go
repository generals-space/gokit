package main

import (
	discovery "github.com/generals-space/gokit/06.gokit-playground-example/63.gokit-lorem-etcd/etcd-sample"
)

func main() {
	master := discovery.NewMaster(discovery.Endpoints)
	master.WatchWorkers()
}
