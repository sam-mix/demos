package conn

import (
	"log"
	"os"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	Client *clientv3.Client
	once   sync.Once
)

func init() {
	once.Do(func() {
		// 客户端配置
		config := clientv3.Config{
			Endpoints:   []string{"127.0.0.1:2379"},
			DialTimeout: 5 * time.Second,
		}
		var err error
		// 建立连接
		if Client, err = clientv3.New(config); err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
	})
}
