package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 启用日志滚动
	logger := &lumberjack.Logger{
		Filename:  "./logs/app.log",
		MaxSize:   10,   // 最大文件大小，MB
		MaxAge:    7,    // 日志保留天数
		Compress:  true, // 是否压缩归档文件
		LocalTime: true, // 是否使用本地时间
	}

	// 将输出写入Log输出中
	log.Logger = zerolog.New(logger).With().
		Caller().
		Timestamp().
		Logger()
	for i := 0; i < 4096000; i++ {
		// 记录info level日志
		log.Info().Msg("info")

		// 记录error level日志
		log.Error().Msg("error")
	}

	// 手动刷新日志记录,确保所有日志都被写入磁盘中
	logger.Rotate()
}
