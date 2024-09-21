package main

import (
	"go.uber.org/zap"
	"time"
)

func NewLogger()(*zap.Logger,error){
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./project.log",
	}
	return cfg.Build()
}

func main()  {
	logger, _ := NewLogger()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "127.0.0.1",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "127.0.0.1")
}