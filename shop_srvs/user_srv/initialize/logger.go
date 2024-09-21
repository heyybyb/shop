package initialize

import "go.uber.org/zap"

// InitLogger 初始化日志
func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
