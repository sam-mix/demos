package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255); not null"`
	Email    string `gorm:"type:varchar(255); uniqueIndex"`
	Password string `gorm:"-"`
	IsAdmin  bool
}

func main() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/test1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	// 创建表格
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	// Assuming a user object
	user := User{Name: "John Smith", Email: "john@example.com", Password: "secret", IsAdmin: true}

	// Inserting user record
	db.Create(&user)

	// Updating user record
	db.Model(&user).Update("Name", "Jane Doe")

	// Querying user record
	db.Where("email = ?", "john@example.com").First(&user)

	// Deleting user record
	db.Delete(&user)
}
