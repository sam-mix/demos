package main

import (
	"log"

	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Price string `gorm:"type:varchar(255);"`
}

func main() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	// 创建表格
	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatal(err)
	}
	a, _ := decimal.NewFromString("1000000000000000000000000000000000000000000000000000000000000")
	b, _ := decimal.NewFromString("1000000000000000000000000000000000000000000000000000000000000")
	c := a.Mul(b)
	// 插入数据
	p := Product{Name: "apple", Price: c.String()}
	err = db.Create(&p).Error
	if err != nil {
		log.Fatalf("Failed to create product: %v", err)
	}

	// 查询数据
	var p2 Product
	err = db.Where("name = ?", "apple").First(&p2).Error
	if err != nil {
		log.Fatalf("Failed to find product: %v", err)
	}
	log.Printf("Product: %d %s %s", p2.ID, p2.Name, p2.Price)
}
