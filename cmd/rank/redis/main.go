package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	// // 创建Redis客户端
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "127.0.0.1:6373",
	// 	Password: "123456", // Redis密码
	// 	DB:       0,        // Redis数据库
	// })

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

	ctx := context.Background()
	// 添加玩家得分
	client.ZAdd(ctx, "leaderboard",
		redis.Z{Score: 100, Member: "player1"},
		redis.Z{Score: 200, Member: "player2"},
		redis.Z{Score: 300, Member: "player3"},
		redis.Z{Score: 250, Member: "player4"},
		redis.Z{Score: 250, Member: "player5"},
	)

	// 获取排行榜前三名
	topPlayers := client.ZRevRangeWithScores(ctx, "leaderboard", 0, 5).Val()

	// 输出排行榜
	for rank, player := range topPlayers {
		fmt.Printf("%d. %s: %v points\n", rank+1, player.Member, player.Score)
	}
}
