package handler

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"ops.was.ink/opsmicro/gateway/config"
	"ops.was.ink/opsmicro/gateway/routers"
)

func serviceInit() web.Service {
	// 1. 初始化 Web Service
	reg := etcd.NewRegistry(
		registry.Addrs(config.EtcdAddress...),
	)

	// 2. 实例化 Web Service 对象
	ws := web.NewService(
		web.Registry(reg),
		web.Address(":10000"),
		web.Handler(routers.GetRouter()),
	)
	
	return ws
}

func GetService() web.Service {
	return serviceInit()
}

