package services

import (
	"github.com/micro/go-micro/v2"
	userproto "ops.was.ink/opsmicro/auth/proto/user"
)

var (
	UserClient userproto.UserService
)

func clientInit(svc micro.Service) {
	UserClient = userproto.NewUserService("com.ops.user.service.user", svc.Client())
}



