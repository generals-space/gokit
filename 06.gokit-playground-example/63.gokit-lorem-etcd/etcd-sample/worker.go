package discovery

import (
	"context"
	"encoding/json"
	"log"
	"runtime"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Worker struct {
	Name       string
	IP         string
	etcdClient *clientv3.Client
}

// workerInfo is the service register information to etcd
type WorkerInfo struct {
	Name string
	IP   string
	CPU  int
}

func NewWorker(name, IP string, endpoints []string) *Worker {
	etcdClient, err := ConnectEtcd()
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}

	w := &Worker{
		Name:       name,
		IP:         IP,
		etcdClient: etcdClient,
	}
	return w
}

func (w *Worker) HeartBeat() {
	for {
		info := &WorkerInfo{
			Name: w.Name,
			IP:   w.IP,
			CPU:  runtime.NumCPU(),
		}

		key := "workers/" + w.Name
		value, _ := json.Marshal(info)

		resp, err := w.etcdClient.Grant(context.TODO(), TTL)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.etcdClient.Put(context.Background(), key, string(value), clientv3.WithLease(resp.ID))
		if err != nil {
			log.Println("Error update workerInfo:", err)
		}
		time.Sleep(time.Second * 3)
	}
}
