package main

import (
	"context"
	"demos/core/cache"
	"demos/core/db"
	"fmt"

	"github.com/qiniu/x/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Price       string `gorm:"type:varchar(255);"`
	Description string `gorm:"size:255"`
}

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
	// 创建表格
	err = db.DB.AutoMigrate(&Product{})
	if err != nil {
		log.Fatal(err)
	}
	a, _ := decimal.NewFromString("1000000000000000000000000000000000000000000000000000000000000")
	b, _ := decimal.NewFromString("1000000000000000000000000000000000000000000000000000000000000")
	c := a.Mul(b)
	// 插入数据
	p := Product{Name: "apple", Price: c.String()}
	err = db.DB.Create(&p).Error
	if err != nil {
		log.Fatalf("Failed to create product: %v", err)
	}

	// 查询数据
	var p2 Product
	err = db.DB.Where("name = ?", "apple").First(&p2).Error
	if err != nil {
		log.Fatalf("Failed to find product: %v", err)
	}
	log.Printf("Product: %d %s %s", p2.ID, p2.Name, p2.Price)
}
