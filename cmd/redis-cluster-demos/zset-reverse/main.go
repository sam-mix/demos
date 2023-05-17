package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	client := redis.NewClusterClient(&redis.ClusterOptions{
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

	// 向有序集合中添加元素
	err := client.ZAdd(ctx, "myzset",
		redis.Z{
			Score:  1.0,
			Member: "member-1",
		},
		redis.Z{
			Score:  2.0,
			Member: "member-2",
		},
		redis.Z{
			Score:  5.0,
			Member: "member-3",
		}).Err()

	if err != nil {
		panic(err)
	}

	// 使用ZRevRange命令获取有序集合指定成员下标范围内的元素列表，并倒序排序
	// 降序
	// res, err := client.ZRevRangeWithScores(ctx, "myzset", 0, -1).Result()
	// 升序
	res, err := client.ZRangeWithScores(ctx, "myzset", 0, -1).Result()
	if err != nil {
		panic(err)
	}

	// 打印结果
	for _, z := range res {
		fmt.Println(z.Member, z.Score)
	}

	client.ZRem(ctx, "myzset",
		"member-1")
	res, err = client.ZRangeWithScores(ctx, "myzset", 0, -1).Result()
	if err != nil {
		panic(err)
	}

	// 打印结果
	for _, z := range res {
		fmt.Println(z.Member, z.Score)
	}

}
