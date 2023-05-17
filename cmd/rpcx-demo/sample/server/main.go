package main

import (
	"context"
	"fmt"
	"log"

	"github.com/smallnest/rpcx/server"
)

var (
	Count = 1
)

type Arithmetic int

func (t *Arithmetic) Mul(ctx context.Context, args *struct{ A, B int }, reply *int) error {
	*reply = args.A * args.B
	Count++
	fmt.Println("Count:", Count)
	return nil
}

func (t *Arithmetic) Add(ctx context.Context, args *struct{ A, B int }, reply *int) error {
	*reply = args.A + args.B
	Count++
	fmt.Println("Count:", Count)
	return nil
}

func main() {
	s := server.NewServer()
	s.RegisterName("Arithmetic", new(Arithmetic), "")
	err := s.Serve("tcp", "127.0.0.1:8972")
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
