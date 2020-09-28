package utils

import (
	"go.uber.org/zap"
)

func init() {
	// 1. 初始化 log 对象
	LogInit()
	// 2. 初始化 gorm 对象
	if err := DBInit(); err != nil {
		log.Fatal("gorm 初始化失败",
			zap.String("errors", err.Error()),
		)
	} else {
		log.Info("gorm 初始化完成...")
	}
}
