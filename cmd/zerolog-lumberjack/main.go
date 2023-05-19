package main

import (
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	logFile := &lumberjack.Logger{
		Filename:   "logs.log",
		MaxSize:    5,
		MaxBackups: 3,
		MaxAge:     180,
		Compress:   true,
	}
	defer logFile.Close()

	// 设置时间戳名字
	zerolog.TimestampFieldName = "timestamp"                  //设置时间戳字段名
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000Z08:00" //设置时间戳格式
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	multi := zerolog.MultiLevelWriter(logFile, os.Stdout) //同时输出到文件和控制台

	log.Logger = log.Output(multi).With().Caller().Logger() //设置logger输出方式，添加Caller信息
	log.Debug().Msg("This is a log message.")
}
