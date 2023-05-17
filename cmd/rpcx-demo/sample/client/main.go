package main

import (
	"context"
	"fmt"

	"github.com/smallnest/rpcx/client"
)

type Args struct {
	A int
	B int
}

func main() {
	d, err := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8972", "")
	if err != nil {
		panic(err)
	}
	xclient := client.NewXClient("Arithmetic", client.Failtry, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	args := Args{
		A: 10,
		B: 20,
	}
	reply := 0
	err = xclient.Call(context.Background(), "Mul", &args, &reply)
	if err != nil {
		fmt.Printf("failed to call: %v", err)
	}
	fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)

	args = Args{
		A: 10,
		B: 20,
	}
	reply = 0
	err = xclient.Call(context.Background(), "Add", &args, &reply)
	if err != nil {
		fmt.Printf("failed to call: %v", err)
	}
	fmt.Printf("%d + %d = %d\n", args.A, args.B, reply)
}
