package main

import (
	"context"
	"fmt"

	example "demos/rpcxpb" // example.pb.go的路径

	"github.com/smallnest/rpcx/server"
)

type Greeter struct{}

func (t *Greeter) Greet(ctx context.Context, req *example.Request, rsp *example.Response) error {
	rsp.Message = fmt.Sprintf("Hello, %s!", req.Name)
	return nil
}

func main() {
	s := server.NewServer()
	s.RegisterName("Greeter", new(Greeter), "")
	s.Serve("tcp", ":50605")
}
