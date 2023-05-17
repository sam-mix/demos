package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"go.etcd.io/etcd/clientv3"
// 	"go.etcd.io/etcd/mvcc/mvccpb"
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

// 	// 注册并启动健康检查协程
// 	if err := registerHealthCheck(cli, "example-service", "localhost:8080"); err != nil {
// 		panic(err)
// 	}

// 	// 监听 etcd 键值对变化
// 	go watchKey(cli, "/services/example-service")
// }

// func registerHealthCheck(cli *clientv3.Client, name, addr string) error {
// 	// 创建一个租约
// 	leaseGrantResp, err := cli.Grant(context.Background(), 10)
// 	if err != nil {
// 		return err
// 	}

// 	// 同时更新节点信息和健康检查
// 	_, err = cli.Put(context.Background(), fmt.Sprintf("/services/%s/%s", name, addr), addr, clientv3.WithLease(leaseGrantResp.ID))
// 	if err != nil {
// 		return err
// 	}

// 	// 自动续约
// 	ch, err := cli.KeepAlive(context.Background(), leaseGrantResp.ID)
// 	if err != nil {
// 		return err
// 	}

// 	go func() {
// 		for {
// 			_, err = cli.Put(context.Background(), fmt.Sprintf("/services/%s/%s", name, addr), addr, clientv3.WithLease(leaseGrantResp.ID))
// 			if err != nil {
// 				panic(err)
// 			}

// 			<-ch
// 		}
// 	}()

// 	return nil
// }

// func watchKey(cli *clientv3.Client, prefix string) {
// 	rch := cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
// 	for wresp := range rch {
// 		for _, ev := range wresp.Events {
// 			switch ev.Type {
// 			case mvccpb.PUT:
// 				fmt.Printf("Added key:%s value:%s\n", ev.Kv.Key, ev.Kv.Value)
// 			case mvccpb.DELETE:
// 				fmt.Printf("Deleted key:%s\n", ev.Kv.Key)
// 			}
// 		}
// 	}
// }
