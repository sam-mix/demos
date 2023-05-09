package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// 创建TCP服务器
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ln.Close()

	// 等待连接
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error, 1)
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				// 处理Accept错误，可能是因为服务器正在关闭
				select {
				case <-ctx.Done():
					return
				default:
					errChan <- err
				}
			}
			wg.Add(1)
			go func(conn net.Conn) {
				handleConn(conn)
				wg.Done()
			}(conn)
		}
	}()

	<-ctx.Done()

	// 关闭TCP服务器
	fmt.Println("Shutdown TCP server...")
	if err := ln.Close(); err != nil {
		fmt.Println(err)
	}

	// 等待所有协程结束
	fmt.Println("Wait for all goroutines to exit...")
	wg.Wait()

	// 停止服务
	fmt.Println("Stop server gracefully.")
	stop()
	os.Exit(0)
}

func handleConn(conn net.Conn) {
	fmt.Printf("Received connection from %s\n", conn.RemoteAddr())
	defer conn.Close()

	// 读取客户端信息
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		// 处理Read错误，可能是因为客户端关闭了连接
		if err != io.EOF {
			fmt.Println(err)
		}
		return
	}
	fmt.Printf("Received data from %s: %s\n", conn.RemoteAddr(), buf[:n])

	// 模拟长时间处理任务
	time.Sleep(2 * time.Second)
	fmt.Fprintf(conn, "Hello, World!\n")

	fmt.Printf("Connection from %s closed.\n", conn.RemoteAddr())

}
