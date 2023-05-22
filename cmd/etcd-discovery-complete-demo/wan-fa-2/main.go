package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
		// handle error
	}
	watcher := client.Watch(context.Background(), "/prefix/", clientv3.WithPrefix())
	for resp := range watcher {
		for _, ev := range resp.Events {
			switch ev.Type {
			case mvccpb.PUT:
			}
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
