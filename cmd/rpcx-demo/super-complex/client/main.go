package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/smallnest/rpcx/client"
)

type Args struct {
	A int
	B map[string]interface{}
}

type Data struct {
	Name string
	Age  int
}

type Reply struct {
	Result string
	Data   []Data
}

func main() {
	d, err := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8972", "")
	if err != nil {
		panic(err)
	}
	xclient := client.NewXClient("supercomplex", client.Failtry, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()
	argsMap := make(map[string]interface{})
	argsMap["a"] = 123
	argsMap["b"] = "test"
	argsMapJson, _ := json.Marshal(argsMap)
	fmt.Println("argsMapJson: ", string(argsMapJson))
	args := Args{
		A: 123,
		B: argsMap,
	}
	reply := &Reply{}
	ctx := context.Background()
	err = xclient.Call(ctx, "SuperComplexFunc", &args, reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply.Result, reply.Data)
}
