package cache

import (
	"flag"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	RedisClient *redis.Client
	once        sync.Once
)

func init() {
	once.Do(func() {
		log.Println("redis init()")
		configFile := flag.String("redis-conf", "./conf/dev/redis.conf", "redis 配置文件路径")
		flag.Parse()
		if RedisClient == nil {
			viper.SetConfigFile(*configFile)
			viper.ReadInConfig()
			redisURL := viper.GetString("url")
			redisPassword := viper.GetString("password")
			redisDB := viper.GetInt("db")
			redisPoolSize := viper.GetInt("pool_size")

			redisClient := redis.NewClient(&redis.Options{
				Addr:     redisURL,
				Password: redisPassword,
				DB:       redisDB,
				PoolSize: redisPoolSize,
			})
			RedisClient = redisClient
		}
	})
}
