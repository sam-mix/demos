package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	brokers = []string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9098"}
)

const (
	topic    = "kafka_demo"
	groupId  = "test-group"
	clientId = "test-client"
)

func main() {
	// // 创建Kafka Admin Client
	// admin, err := kafka.NewAdminClient(context.Background(), kafka.ConfigMap{"bootstrap.servers": brokers})
	// if err != nil {
	// 	fmt.Printf("failed to create admin client: %s\n", err.Error())
	// 	return
	// }

	// // 创建主题需要提供主题名称，分区数，复制因子等参数
	// topicConfigs := []kafka.TopicConfig{
	// 	{
	// 		Topic:             topic,
	// 		NumPartitions:     3, // 分区数
	// 		ReplicationFactor: 2, // 复制因子
	// 	},
	// }

	// // 调用CreateTopics函数创建主题
	// results, err := admin.CreateTopics(context.Background(), topicConfigs...)
	// if err != nil {
	// 	fmt.Printf("failed to create topic: %s\n", err.Error())
	// 	return
	// }

	// // 打印每个分区的副本分配情况
	// for _, result := range results {
	// 	fmt.Printf("topic: %s, partition: %d, replicas: %s\n", result.Topic, result.Partition, result.Replicas)
	// }

	// // 关闭Kafka Admin Client
	// admin.Close()

	// 创建Kafka消费者
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupId,
		MinBytes: 10e3,             // 最小消息字节数
		MaxBytes: 10e6,             // 最大消息字节数
		MaxWait:  10 * time.Second, // 最大等待时间
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

	// 创建Kafka生产者
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		Async:    true, // 异步发送消息
	})

	// 生产消息
	for i := 0; i < 10; i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("key%d", i)),
			Value: []byte(fmt.Sprintf("message %d", i)),
		}
		err := w.WriteMessages(ctx, msg)
		if err != nil {
			fmt.Printf("error producing message: %s\n", err.Error())
			break
		}
	}

	// 关闭Kafka生产者
	w.Close()

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
