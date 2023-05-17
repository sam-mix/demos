package main

import (
	"context"
	"fmt"

	example "demos/rpcxpb" // example.pb.go的路径

	"github.com/smallnest/rpcx/client"
)

func main() {
	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:50605", "")
	xclient := client.NewXClient("Greeter", client.Failtry, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	req := &example.Request{Name: "World"}
	rsp := &example.Response{}
	err := xclient.Call(context.Background(), "Greet", req, rsp)
	if err != nil {
		fmt.Printf("failed to call: %v\n", err)
		return
	}
	fmt.Printf("%s\n", rsp.Message)
}
