package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error, 1)

	// 创建HTTP服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("HTTP server is listening at :8080")
		errChan <- server.ListenAndServe()
	}()

	<-ctx.Done()

	// 关闭HTTP服务器
	fmt.Println("Shutdown HTTP server...")
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctxShutDown); err != nil {
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
