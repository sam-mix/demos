package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/segmentio/kafka-go"
)

func main() {
	// 创建Kafka消费者
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"127.0.0.1:9092"},
		Topic:   "test-topic",
	})

	// 创建一个context用于取消消费者的读取
	ctx, cancel := context.WithCancel(context.Background())

	// 使用goroutine异步读取Kafka消息
	go func() {
		for {
			msg, err := r.ReadMessage(ctx)
			if err != nil {
				fmt.Printf("error reading message: %s\n", err.Error())
				break
			}
			fmt.Printf("message received: %s\n", string(msg.Value))
		}
	}()

	// 监听信号以取消消费者读取
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	select {
	case <-sigchan:
		fmt.Println("interrupt signal received, cancelling")
	case <-ctx.Done():
		fmt.Println("context cancelled, exiting")
	}

	// 关闭Kafka消费者
	r.Close()
	cancel()
}
