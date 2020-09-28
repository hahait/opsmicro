package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/errors"
	authproto "ops.was.ink/opsmicro/auth/proto/auth"
	userproto "ops.was.ink/opsmicro/auth/proto/user"
	"ops.was.ink/opsmicro/auth/services"
	"ops.was.ink/opsmicro/auth/utils"
)

type Auth struct{}

type Login struct {
	Username string
	Password string
}

func (a *Auth) Login(ctx context.Context, in *authproto.Request, out *authproto.Response) error {
	userinfo, err := services.UserClient.QueryUser(ctx, &userproto.Request{Username: in.Username})
	if err != nil {
		return errors.InternalServerError("com.ops.auth.service.auth.Login", "RPC 请求 com.ops.user.service.user.Queryuser 失败", err.Error())
	}

	if ok := utils.CheckPassword(in.Password, userinfo.Password); !ok {
		return errors.InternalServerError("com.ops.auth.service.auth.Login", "密码校验失败, 请重新输入密码")
	}

	j := new(services.JwtAuth)
	token, err := j.GenerateJwtToken(in.Username)
	if err != nil {
		return errors.InternalServerError("com.ops.auth.service.auth.Login", fmt.Sprintf(" Token 生成失败, 错误信息: %s", err.Error()))
	}

	out.Token = token

	return nil
}

func (a *Auth) ValidateAccessToken(ctx context.Context, in *authproto.ValidateTokenRequest, out *authproto.ValidateTokenResponse) error {
	j := new(services.JwtAuth)
	b, err := j.ValidateJwtToken(in.Token)
	out.Success = b
	if err != nil {
		return err
	}
	return nil
}