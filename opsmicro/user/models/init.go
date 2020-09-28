package models

import (
	"context"
	"go.uber.org/zap"
	"ops.was.ink/opsmicro/user/utils"
)


func init() {
	err := utils.GetDB(context.TODO()).AutoMigrate(
		&User{},
	)
	if err != nil {
		utils.GetLogger().Fatal("自动迁移 models 到数据库失败",
			zap.String("errors", err.Error()),
		)
	} else {
		utils.GetLogger().Info("自动迁移 models 到数据库完成...")
	}
}
