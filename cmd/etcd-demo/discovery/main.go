package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"go.etcd.io/etcd/clientv3"
// )

// func main() {
// 	// 创建一个 etcd 客户端实例
// 	cli, err := clientv3.New(clientv3.Config{
// 		Endpoints:   []string{"localhost:2379"}, // etcd 服务器地址
// 		DialTimeout: 5 * time.Second,            // 连接超时时间
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer cli.Close()

// 	// 发现服务
// 	addrs, err := discoverService(cli, "example-service")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("service nodes:", addrs)
// }

// func discoverService(cli *clientv3.Client, name string) ([]string, error) {
// 	var addrs []string

// 	// 获取 name 服务的所有节点
// 	resp, err := cli.Get(context.Background(), fmt.Sprintf("/services/%s/", name), clientv3.WithPrefix())
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, kv := range resp.Kvs {
// 		addrs = append(addrs, string(kv.Value))
// 	}

// 	return addrs, nil
// }
