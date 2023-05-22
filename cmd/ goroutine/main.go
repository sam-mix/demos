package main

import (
	"context"
	"time"
)

func myGoroutine(ctx context.Context) {
	for {
		// 执行正常的业务逻辑
		select {
		case <-ctx.Done():
			return // 监听到停止信号，结束goroutine
		default:
			time.Sleep(1 * time.Second)
			// 继续执行正常的业务逻辑
			print("hello")
		}
	}
}

func startGoroutine(ctx context.Context) {
	go myGoroutine(ctx)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	startGoroutine(ctx)
	time.Sleep(10 * time.Second)
	// 在需要停止goroutine时，调用cancel方法
	cancel()
}
