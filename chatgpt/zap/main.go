package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	logger := createLogger()
	defer logger.Sync()

	for i := 0; i < 40960; i++ {
		logger.Info("test log", zap.Int("i", i))
	}
}

func createLogger() *zap.Logger {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	core := zapcore.NewCore(
		encoder,
		getLogWriter(),   // 使用getLogWriter()获取输出到文件句柄的 WriteSyncer
		getAtomicLevel(), // 获取 zap.AtomicLevel
	)

	return zap.New(core, zap.AddCaller())
}

func getLogWriter() zapcore.WriteSyncer {
	// 获取日志文件路径
	filePath := getRollingLogFileName()
	// 获取到目录路径，并创建目录
	dirPath := filepath.Dir(filePath)
	_ = os.MkdirAll(dirPath, os.ModePerm)

	// 尺寸滚动文件配置
	rollingConfig := zapcore.AddSync(&lumberjack.Logger{
		Filename:  filePath,
		MaxSize:   1,     // 每个日志文件最大10MB
		LocalTime: true,  // 文件名用本地时间表示，而非UTC
		Compress:  false, // 不启用压缩
	})

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), rollingConfig)
}

func getRollingLogFileName() string {
	// 设置日志文件名，其中{xxx}为需要被替换的字段
	filename := "tcp_server_{date}_{time}.log"
	filename = filepath.Join("logs", filename)

	// 将当前日期和时间格式化，用于替换日志文件名中的{xxx}字段
	now := time.Now()
	dateString := now.Format("2006-01-02")
	timeString := now.Format("15-04-05.000")

	filename = filepath.Clean(filename)
	filename = filepath.Join(filepath.Dir(filename), fmt.Sprintf("%s_%s", dateString, timeString)+filepath.Ext(filename))

	return filename
}

func getAtomicLevel() zap.AtomicLevel {
	return zap.NewAtomicLevelAt(zap.DebugLevel)
}
