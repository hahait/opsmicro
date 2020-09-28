package services

import (
	"context"
	"fmt"
	sentinelplg "github.com/alibaba/sentinel-golang/adapter/micro"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	popentracing "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"ops.was.ink/opsmicro/auth/basic/config"
	//ajaeger "ops.was.ink/opsmicro/auth/basic/jaeger"
	"time"
)


func Haha(c context.Context, cli client.Request, berr *base.BlockError) error {
	fmt.Println("熔断类型: ", berr.BlockType(), "熔断阈值: ", berr.TriggeredValue(), "熔断规则: ", berr.TriggeredRule())
	return nil
}

func serviceInit(sname string) micro.Service {
	reg := etcd.NewRegistry(
		registry.Addrs(config.EtcdAddress...),
	)

	service := micro.NewService(
		micro.Name(sname),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.RegisterInterval(time.Second * 30),
		micro.RegisterTTL(time.Minute),
		micro.Address(":10200"),
		micro.WrapClient(sentinelplg.NewClientWrapper(sentinelplg.WithClientBlockFallback(Haha))),
		micro.WrapHandler(popentracing.NewHandlerWrapper(opentracing.GlobalTracer())),
		//micro.WrapClient(popentracing.NewClientWrapper(opentracing.GlobalTracer())),
	)

	clientInit(service)

	service.Init(
		micro.AfterStart(func() error {
			fmt.Println("Auth 启动成功...")
			return nil
		}),
	)

	return service
}

func GetService(sname string) micro.Service {
	return serviceInit(sname)
}


