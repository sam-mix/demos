package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Config struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Article struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	// 初始化Redis客户端
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })

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

	ctx := context.Background()

	// 缓存用户信息
	user := User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}
	userJSON, _ := json.Marshal(user)
	err := rdb.Set(ctx, "user:1", userJSON, 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取缓存的用户信息
	var cachedUser User
	cachedUserJSON, err := rdb.Get(ctx, "user:1").Bytes()
	if err == redis.Nil {
		fmt.Println("User not found in cache")
	} else if err != nil {
		panic(err)
	} else {
		json.Unmarshal(cachedUserJSON, &cachedUser)
		fmt.Printf("Cached user: %+v\n", cachedUser)
	}

	// 缓存配置信息
	config := Config{
		Key:   "database_url",
		Value: "postgres://user:password@localhost/mydb",
	}
	configJSON, _ := json.Marshal(config)
	err = rdb.Set(ctx, "config:database_url", configJSON, 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取缓存的配置信息
	var cachedConfig Config
	cachedConfigJSON, err := rdb.Get(ctx, "config:database_url").Bytes()
	if err == redis.Nil {
		fmt.Println("Config not found in cache")
	} else if err != nil {
		panic(err)
	} else {
		json.Unmarshal(cachedConfigJSON, &cachedConfig)
		fmt.Printf("Cached config: %+v\n", cachedConfig)
	}

	// 缓存文章列表
	articles := []Article{
		{ID: 1, Title: "Article 1", Content: "Content of article 1"},
		{ID: 2, Title: "Article 2", Content: "Content of article 2"},
		{ID: 3, Title: "Article 3", Content: "Content of article 3"},
	}
	articlesJSON, _ := json.Marshal(articles)
	err = rdb.Set(ctx, "articles:list", articlesJSON, 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取缓存的文章列表
	var cachedArticles []Article
	cachedArticlesJSON, err := rdb.Get(ctx, "articles:list").Bytes()
	if err == redis.Nil {
		fmt.Println("Articles not found in cache")
	} else if err != nil {
		panic(err)
	} else {
		json.Unmarshal(cachedArticlesJSON, &cachedArticles)
		fmt.Printf("Cached articles: %+v\n", cachedArticles)
	}

	// 初始化计数器
	err = rdb.Set(ctx, "counter", 0, 0).Err()
	if err != nil {
		panic(err)
	}

	// 模拟计数
	for i := 0; i < 10; i++ {
		_, err = rdb.Incr(ctx, "counter").Result()
		if err != nil {
			panic(err)
		}
	}

	// 获取计数器的值
	count, err := rdb.Get(ctx, "counter").Int64()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Counter value: %d\n", count)
}
