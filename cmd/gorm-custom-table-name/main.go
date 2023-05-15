package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255); not null"`
	Email string `gorm:"type:varchar(255); uniqueIndex"`
}

func (User) TableName() string {
	return "my_users"
}

func main() {
	// Connect to MySQL database
	dsn := "root:123456@tcp(127.0.0.1:3306)/test1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto-migrate the "my_users" table
	db.AutoMigrate(&User{})
}
