package main

import (
	"github.com/opentracing/opentracing-go"
	ujaeger "ops.was.ink/opsmicro/user/basic/jaeger"
	ucfg "ops.was.ink/opsmicro/user/config"
	"ops.was.ink/opsmicro/user/handler"
	_ "ops.was.ink/opsmicro/user/models"
	user "ops.was.ink/opsmicro/user/proto/user"
	"ops.was.ink/opsmicro/user/utils"
	"ops.was.ink/opsmicro/user/services"
)

func main() {
	svcname := ucfg.ServiceName
	t, c := ujaeger.JaegerInit(svcname)
	defer c.Close()
	opentracing.SetGlobalTracer(t)

	svc := services.GetService(svcname)
	user.RegisterUserHandler(svc.Server(), new(handler.User))

	if err := svc.Run(); err != nil {
		utils.GetLogger().Fatal(err.Error())
	}
}
