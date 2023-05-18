package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	// 创建Redis客户端实例
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:6371",
			"127.0.0.1:6372",
			"127.0.0.1:6373",
			"127.0.0.1:6374",
			"127.0.0.1:6375",
			"127.0.0.1:6376",
		},
		Password: "123456",
	})

	// 定义消息体
	msg := map[string]interface{}{
		"id":      "001",
		"title":   "test message",
		"content": "this is a test message",
	}

	// 发送消息到Stream
	_, err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "mystream", // Stream名称
		ID:     "*",        // 自动生成ID
		Values: msg,        // 消息体
	}).Result()
	if err != nil {
		panic(err)
	}

	// 读取Stream中的消息
	msgs, err := rdb.XRead(ctx, &redis.XReadArgs{
		Streams: []string{"mystream", "0"}, // Stream名称和偏移量
		Count:   10,                        // 一次最多读取10条消息
		Block:   0,                         // 阻塞等待新消息
	}).Result()
	if err != nil {
		panic(err)
	}

	// 打印消息
	for _, msg := range msgs[0].Messages {
		fmt.Printf("Message ID: %s\n", msg.ID)
		for k, v := range msg.Values {
			fmt.Printf(" %s: %v\n", k, v)
		}
	}
}
