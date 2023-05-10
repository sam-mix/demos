package main

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

// 玩家
type Player struct {
	gorm.Model
	Name string // 名字
	Gold uint64 // 金币
	Lev  uint16 // 等级
}

func main() {
	// // 配置日志
	// logFileName := "log/application.log"
	// fileWriter := zerolog.NewFileWriter(logFileName)

	// // 设置日志文件最大尺寸为 100 MB，最大保留文件数为 10
	// maxFileSize := int64(100 * 1024 * 1024) // 100 MB
	// maxRetainedFiles := 10
	// fileWriter = fileWriter.SizeLimit(maxFileSize).RollingFile(
	// 	&zerolog.RollingFileConfig{
	// 		MaxSize:    maxFileSize,
	// 		MaxAge:     7 * 24 * time.Hour, // 文件最大保存时间为 7 天
	// 		MaxBackups: maxRetainedFiles,
	// 	},
	// )

	// // 设置 zerolog 全局日志级别为 Debug 级别
	// zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// // 将 zerolog 错误处理器设置为 pkgerrors.HandlerWithStacktrace，以记录堆栈信息
	// zerolog.ErrorStackMarshaler = pkgerrors.HandlerWithStacktrace("stacktrace")

	// // 创建 logger
	// logger := zerolog.New(fileWriter).With().Timestamp().Logger()

	// 配置数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&User{}, &Player{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate database")
	}

	// 创建用户
	user := &User{Name: "Alice", Email: "alice@example.com"}
	result := db.Create(user)
	if result.Error != nil {
		log.Fatal().Err(result.Error).Msg("Failed to create user")
	}
	log.Info().Msg("Created user")

	// 查询用户
	var users []User
	result = db.Find(&users)
	if result.Error != nil {
		log.Fatal().Err(result.Error).Msg("Failed to query users")
	}
	log.Info().Msgf("Found %d users", len(users))
	for _, user := range users {
		log.Info().Msgf("User: %+v", user)
	}

	// 更新用户
	user.Name = "Bob"
	result = db.Save(user)
	if result.Error != nil {
		log.Fatal().Err(result.Error).Msg("Failed to update user")
	}
	log.Info().Msg("Updated user")

	// 删除用户
	result = db.Delete(user)
	if result.Error != nil {
		log.Fatal().Err(result.Error).Msg("Failed to delete user")
	}
	log.Info().Msg("Deleted user")
}
