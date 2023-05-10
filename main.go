package main

import (
	"context"
	"demos/core/cache"
	"fmt"

	"github.com/qiniu/x/log"
)

func main() {
	context := context.Background()

	pong, err := cache.RedisClient.Ping(context).Result()
	if err != nil {
		log.Error("connect to redis error: ", err.Error())
	}
	fmt.Println("redis pong: ", pong)
	key := "hello"
	val := "world"
	err = cache.RedisClient.Set(context, key, val, 0).Err()
	if err != nil {
		log.Error("set redis error: ", err.Error())
	}

	res, err := cache.RedisClient.Get(context, key).Result()
	if err != nil {
		log.Error("get redis error: ", err.Error())
	}
	fmt.Printf("%s : %s\n", key, res)
}
