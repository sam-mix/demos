package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	for {
		logURL("baidu.com")
		time.Sleep(500 * time.Millisecond)
	}
}

func logURL(url string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

}
