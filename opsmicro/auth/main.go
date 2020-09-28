package main

import (
	"github.com/opentracing/opentracing-go"
	"log"
	_ "ops.was.ink/opsmicro/auth/basic"
	acfg "ops.was.ink/opsmicro/auth/basic/config"
	ajaeger "ops.was.ink/opsmicro/auth/basic/jaeger"
	"ops.was.ink/opsmicro/auth/handler"
	authproto "ops.was.ink/opsmicro/auth/proto/auth"
	"ops.was.ink/opsmicro/auth/services"
)

func main() {
	svcname := acfg.ServiceName
	// 实例化 jaeger tracer 对象, 并设置成全局单例模式对象
	t, c := ajaeger.JaegerInit(svcname)
	defer c.Close()
	opentracing.SetGlobalTracer(t)

	s := services.GetService(svcname)
	if err := authproto.RegisterAuthHandler(s.Server(), new(handler.Auth)); err != nil {
		log.Fatal("注册服务失败, 错误信息: ", err.Error())
	}

	if err := s.Run(); err != nil {
		log.Fatal("启动服务失败, 错误信息: ", err.Error())
	}
}
