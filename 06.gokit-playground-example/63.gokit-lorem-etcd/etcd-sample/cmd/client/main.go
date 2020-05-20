package main

import (
	"sample/pkg/discovery"
)

func main() {
	worker := discovery.NewWorker("node-01", "127.0.0.1", discovery.Endpoints)
	worker.HeartBeat()
}
