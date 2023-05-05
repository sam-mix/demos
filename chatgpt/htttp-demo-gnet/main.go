package main

import (
	"fmt"
	"time"

	"github.com/panjf2000/gnet"
)

type echoServer struct {
	gnet.EventHandler
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	out = frame
	return
}

func main() {
	server := &echoServer{}

	// 创建一个 gnet 服务器并监听 8080 端口
	serverOpts := []gnet.Option{
		gnet.WithMulticore(true),
	}
	err := gnet.Serve(server, "tcp://0.0.0.0:8080", serverOpts...)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务器启动后，等待 1 分钟
	time.Sleep(1 * time.Minute)
}
