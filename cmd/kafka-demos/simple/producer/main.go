package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// 创建Kafka生产者
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9098"},
		Topic:    "kafka_demo",
		Balancer: &kafka.LeastBytes{},
	})

	// 生产消息
	for i := 0; i < 1000; i++ {
		msg := kafka.Message{
			Value: []byte(fmt.Sprintf("message %d at %v", i, time.Now().Format(time.RFC3339))),
		}
		err := w.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Printf("error producing message: %s\n", err.Error())
			break
		}
	}

	// 关闭Kafka生产者
	w.Close()
}
