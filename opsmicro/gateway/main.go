package main

import (
        log "github.com/micro/go-micro/v2/logger"
        "github.com/micro/go-micro/v2/registry"
        "github.com/micro/go-micro/v2/registry/etcd"
        "github.com/micro/go-micro/v2/web"
        "github.com/opentracing/opentracing-go"
        "ops.was.ink/opsmicro/gateway/routers"
        "ops.was.ink/opsmicro/gateway/config"
        _ "ops.was.ink/opsmicro/gateway/handler"
        gjaeger "ops.was.ink/opsmicro/gateway/basic/jaeger"
        gcfg "ops.was.ink/opsmicro/gateway/config"
)

func main() {
        // 1. 初始化 Web Service
        reg := etcd.NewRegistry(
                registry.Addrs(config.EtcdAddress...),
        )
        svcname := gcfg.ServiceName
        t, c := gjaeger.JaegerInit(svcname)
        defer c.Close()
        opentracing.SetGlobalTracer(t)

        // 2. 实例化 Web Service 对象
        ws := web.NewService(
                web.Registry(reg),
                web.Address(":10000"),
                web.Handler(routers.GetRouter()),
        )

        if err := ws.Init(
                web.AfterStart(func() error {
                        log.Info(" gateway 启动成功...")
                        return nil
                }),
        ); err != nil {
                log.Fatal(err)
        }

        if err := ws.Run(); err != nil {
                log.Fatal(err)
        }
}
