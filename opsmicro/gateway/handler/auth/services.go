package auth

import (
	"github.com/micro/go-micro/v2/client"
	authproto "ops.was.ink/opsmicro/gateway/proto/auth"
)

var (
	// 请求 auth 服务的接口
	AuthClient = authproto.NewAuthService("com.ops.auth.service.auth", client.DefaultClient)
)

