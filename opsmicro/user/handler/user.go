package handler

import (
	"context"
	"github.com/micro/go-micro/v2/errors"
	"ops.was.ink/opsmicro/user/models"
	user "ops.was.ink/opsmicro/user/proto/user"
	"ops.was.ink/opsmicro/user/utils"
	"go.uber.org/zap"
)

type User struct{}

func (u *User) QueryUser(ctx context.Context, in *user.Request, out *user.Response) error {
	if in.Username != "" {
		if err := utils.GetDB(ctx).Model(models.User{}).Where("username = ?", in.Username).Scan(out).Error; err != nil {
			utils.GetLogger().Error("数据库查询信息失败", zap.String("errors", err.Error()))
			return errors.InternalServerError("com.ops.user.service.user.QueryUser", err.Error())
		}
		return nil
	} else {
		if err := utils.GetDB(ctx).Model(models.User{}).Where("phone = ?", in.Phone).Scan(out).Error; err != nil {
			utils.GetLogger().Error("数据库查询信息失败", zap.String("errors", err.Error()))
			return errors.InternalServerError("com.ops.user.service.user.QueryUser", err.Error())
		}
		return nil
	}
	return nil
}