package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	serviceName = "test_service"
)

type Registry struct {
	cli         *clientv3.Client
	leaseID     clientv3.LeaseID
	registerKey string
	stopCh      chan struct{}
}

func NewRegistry(cli *clientv3.Client, ttl int64) (*Registry, error) {
	r := &Registry{
		cli:    cli,
		stopCh: make(chan struct{}),
	}

	resp, err := cli.Grant(context.Background(), ttl)
	if err != nil {
		return nil, err
	}
	r.leaseID = resp.ID

	r.registerKey = fmt.Sprintf("/%s/%d", serviceName, r.leaseID)

	if _, err := cli.Put(context.Background(), r.registerKey, "", clientv3.WithLease(r.leaseID)); err != nil {
		return nil, err
	}

	go r.keepAlive()

	return r, nil
}

func (r *Registry) keepAlive() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if _, err := r.cli.KeepAliveOnce(context.Background(), r.leaseID); err != nil {
				log.Printf("failed to keep alive lease: %v", err)
			}
		case <-r.stopCh:
			return
		}
	}
}

func (r *Registry) Register(addr string) error {
	key := fmt.Sprintf("%s/%s", r.registerKey, addr)

	if _, err := r.cli.Put(context.Background(), key, strconv.Itoa(int(time.Now().Unix())), clientv3.WithLease(r.leaseID)); err != nil {
		return err
	}

	return nil
}

func (r *Registry) Unregister(addr string) error {
	key := fmt.Sprintf("%s/%s", r.registerKey, addr)

	if _, err := r.cli.Delete(context.Background(), key); err != nil {
		return err
	}

	return nil
}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	reg, err := NewRegistry(cli, 10)
	if err != nil {
		log.Fatal(err)
	}

	if err := reg.Register("127.0.0.1:1234"); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Minute)

	if err := reg.Unregister("127.0.0.1:1234"); err != nil {
		log.Fatal(err)
	}
}
