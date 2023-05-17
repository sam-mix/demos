package main

// import (
// 	"context"
// 	"demos/etcd-demos/conn"
// 	"fmt"

// 	"go.etcd.io/etcd/clientv3"
// )

// func main() {
// 	fmt.Println("conn.Client:", conn.Client)
// 	// 实例化一个用于操作ETCD的KV
// 	kv := clientv3.NewKV(conn.Client)
// 	var putResp *clientv3.PutResponse
// 	var err error
// 	if putResp, err = kv.Put(context.TODO(), "/school/class/students", "helios0", clientv3.WithPrevKV()); err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("putResp.Header.Revisio:", putResp.Header.Revision)
// 	if putResp.PrevKv != nil {
// 		fmt.Printf("prev Value: %s \n CreateRevision : %d \n ModRevision: %d \n Version: %d \n",
// 			string(putResp.PrevKv.Value), putResp.PrevKv.CreateRevision, putResp.PrevKv.ModRevision, putResp.PrevKv.Version)
// 	}
// 	var getResp *clientv3.GetResponse
// 	if getResp, err = kv.Get(context.TODO(), "/school/class/students"); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// 输出本次的Revision
// 	fmt.Printf("Key is: %s \nValue is: %s \n", getResp.Kvs[0].Key, getResp.Kvs[0].Value)
// 	_, err = kv.Put(context.TODO(), "/school/class/students", "helios1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	var delResp *clientv3.DeleteResponse
// 	if delResp, err = kv.Delete(context.TODO(), "/school/class/students", clientv3.WithPrevKV()); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if len(delResp.PrevKvs) != 0 {
// 		for _, kvpair := range delResp.PrevKvs {
// 			fmt.Printf("delete key is: %s \nValue is: %s \n", string(kvpair.Key), string(kvpair.Value))
// 		}
// 	}
// 	// 申请一个租约
// 	lease := clientv3.NewLease(conn.Client)
// 	var leaseGrantRsp *clientv3.LeaseGrantResponse
// 	if leaseGrantRsp, err = lease.Grant(context.TODO(), 10); err != nil {
// 		panic(err)
// 	}
// 	leaseId := leaseGrantRsp.ID

// 	if _, err = kv.Put(context.TODO(), "/school/class/students", "h", clientv3.WithLease(leaseId)); err != nil {
// 		panic(err)
// 	}

// 	// for {
// 	// 	if getResp, err = kv.Get(context.TODO(), "/school/class/students"); err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// 	if getResp.Count == 0 {
// 	// 		fmt.Println("kv过期了")
// 	// 		break
// 	// 	}
// 	// 	fmt.Println("还没有过期:", getResp.Kvs)
// 	// 	time.Sleep(2 * time.Second)
// 	// }

// 	// 自动续约
// 	if keepRespChan, err := lease.KeepAlive(context.TODO(), leaseId); err != nil {
// 		panic(err)
// 	} else {
// 		go func() {
// 			for keepResp := range keepRespChan {
// 				if keepRespChan == nil {
// 					fmt.Println("租约已经失效了")
// 					goto END
// 				} else { // 每秒续约一次， 所以就会应答一次
// 					fmt.Println("收到自动续约的应答:", keepResp.ID)
// 				}
// 			}
// 		END:
// 		}()
// 	}

// }
