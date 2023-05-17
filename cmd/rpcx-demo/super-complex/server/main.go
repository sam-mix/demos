package main

import (
	"context"
	"fmt"

	"github.com/smallnest/rpcx/server"
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

type SuperComplexService struct{}

func (s *SuperComplexService) SuperComplexFunc(ctx context.Context, args *Args, reply *Reply) error {
	// 进行超级复杂逻辑处理，并返回结果
	result := fmt.Sprintf("A=%d, B=%v", args.A, args.B)
	reply.Result = result

	data := []Data{
		{Name: "Tom", Age: 20},
		{Name: "Jerry", Age: 21},
		{Name: "Mickey", Age: 22},
	}
	reply.Data = data
	return nil
}

func main() {
	s := server.NewServer()

	s.RegisterName("supercomplex", new(SuperComplexService), "")

	err := s.Serve("tcp", "127.0.0.1:8972")
	if err != nil {
		panic(err)
	}
}
