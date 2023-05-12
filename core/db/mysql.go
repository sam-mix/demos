package db

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func init() {
	once.Do(func() {
		log.Println("mysql init()")
		// 读取配置文件
		configFile := flag.String("mysql-conf", "./conf/dev/mysql.conf", "mysql 配置文件路径")
		flag.Parse()
		viper.SetConfigFile(*configFile)
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Failed to read config file: %v", err)
		}
		// 配置数据库连接
		host := viper.GetString("db.host")
		port := viper.GetString("db.port")
		user := viper.GetString("db.user")
		password := viper.GetString("db.password")
		dbname := viper.GetString("db.name")
		timezone := viper.GetString("db.timezone")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
			user, password, host, port, dbname, timezone)

		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		// 读取配置文件中的 GORM 相关配置
		debug := viper.GetBool("gorm.debug")
		connMaxLifetime := viper.GetDuration("gorm.conn_max_lifetime")
		maxIdleConn := viper.GetInt("gorm.max_idle_conn")
		maxOpenConn := viper.GetInt("gorm.max_open_conn")

		// 进行 GORM 相关配置
		if debug {
			DB = DB.Debug() // 开启 DEBUG 模式，打印 SQL 语句
		}
		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatal(err)
		}
		sqlDB.SetConnMaxLifetime(connMaxLifetime) // 最大连接时间
		sqlDB.SetMaxIdleConns(maxIdleConn)        // 最大空闲连接数
		sqlDB.SetMaxOpenConns(maxOpenConn)        // 最大连接数
	})

}
