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

// 	// 注册服务
// 	if err := registerService(cli, "example-service", "localhost:8080"); err != nil {
// 		panic(err)
// 	}
// }

// func registerService(cli *clientv3.Client, name, addr string) error {
// 	// 创建一个租约
// 	leaseGrantResp, err := cli.Grant(context.Background(), 10)
// 	if err != nil {
// 		return err
// 	}

// 	// 向 etcd 中写入节点信息
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
// 			<-ch
// 		}
// 	}()

// 	return nil
// }
