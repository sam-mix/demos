package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接到 MySQL 数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 检查数据库是否存在
	rows, err := db.Query("SHOW DATABASES LIKE 'testdb'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// 如果数据库不存在，则创建新的数据库
	if !rows.Next() {
		_, err = db.Exec("CREATE DATABASE testdb")
		if err != nil {
			panic(err)
		}
		fmt.Println("Created database 'testdb'")
	} else {
		fmt.Println("Database 'testdb' already exists")
	}
}
