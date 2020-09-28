package services

import (
	sentinelplg "github.com/alibaba/sentinel-golang/adapter/micro"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	popentracing "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	userconf "ops.was.ink/opsmicro/user/config"
	"ops.was.ink/opsmicro/user/utils"
	"time"
)

func serviceInit(sname string) micro.Service {
	// 1. 实例化一个 Registry 对象, 使用 etcd 作为注册中心
	reg := etcd.NewRegistry(
		registry.Addrs(userconf.EtcdAddress...),
	)

	// 2. 实例化一个 Service 对象
	service := micro.NewService(
		micro.Name(sname),
		micro.Version("latest"),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second * 30),
		micro.Address(":10300"),
		micro.Registry(reg),
		micro.WrapHandler(
			popentracing.NewHandlerWrapper(opentracing.GlobalTracer()),
			sentinelplg.NewHandlerWrapper(),
		),
	)

	// 3. 初始化 Service 对象, 并且打印初始化信息
	service.Init(micro.AfterStart(func() error {
			utils.GetLogger().Info("user 启动完成...")
			return nil
		}),
	)

	return service
}

func GetService(sname string) micro.Service {
	return serviceInit(sname)
}