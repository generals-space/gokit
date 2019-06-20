package discovery

import (
	"context"
	"encoding/json"
	"log"

	"go.etcd.io/etcd/clientv3"
)

type Master struct {
	// members保存客户端地址, 需要客户端指定key的TTL,
	// 然后发送心跳, 当客户端停止运行, 其所属key过期时, 会将其从members中删除.
	members    map[string]*Member
	etcdClient *clientv3.Client
}

// Member is a client machine
type Member struct {
	IP   string
	Name string
	CPU  int
}

// NewMaster 创建Master结构对象
// @params endpoints: etcd服务地址列表([x.x.x.x:2379,])
func NewMaster(endpoints []string) *Master {
	etcdClient, err := ConnectEtcd()
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}

	master := &Master{
		members:    make(map[string]*Member),
		etcdClient: etcdClient,
	}

	return master
}

func (m *Master) AddWorker(key string, info *WorkerInfo) {
	m.members[key] = &Member{
		IP:   info.IP,
		Name: info.Name,
		CPU:  info.CPU,
	}
}

func (m *Master) UpdateWorker(key string, info *WorkerInfo) {
	member := m.members[key]
	member.Name = info.Name
}

func GetWorkerInfo(ev *clientv3.Event) *WorkerInfo {
	info := &WorkerInfo{}
	err := json.Unmarshal([]byte(ev.Kv.Value), info)
	if err != nil {
		log.Print(err)
	}
	return info
}

func (m *Master) WatchWorkers() {
	watcher := clientv3.NewWatcher(m.etcdClient)
	wch := watcher.Watch(context.Background(), "workers/", clientv3.WithPrefix(), clientv3.WithRev(0))
	for wresp := range wch {
		for _, ev := range wresp.Events {
			// fmt.Printf("watching: %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			key := string(ev.Kv.Key)
			if ev.Type.String() == "PUT" {
				info := GetWorkerInfo(ev)
				if _, ok := m.members[key]; ok {
					log.Printf("Update worker %s: %s", key, info.Name)
					m.UpdateWorker(key, info)
				} else {
					log.Printf("Add worker %s: %s", key, info.Name)
					m.AddWorker(key, info)
				}
			} else if ev.Type.String() == "DELETE" {
				// key由于过期被删除, 则value值变为空, 没有办法得到其中的信息
				log.Println("Delete worker ", key)
				delete(m.members, key)
			}
		}
	}
}
