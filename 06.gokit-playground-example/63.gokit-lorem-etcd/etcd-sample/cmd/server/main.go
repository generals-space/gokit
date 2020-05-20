package main

import (
	"sample/pkg/discovery"
)

func main() {
	master := discovery.NewMaster(discovery.Endpoints)
	master.WatchWorkers()
}
