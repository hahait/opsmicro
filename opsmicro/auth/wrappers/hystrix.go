package wrappers

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	userproto "ops.was.ink/opsmicro/auth/proto/user"
)

type hystrixClientWrapper struct {
	client.Client
}

func hystrixFallback(rsp interface{}) {
	bb, _ := rsp.(*userproto.Response)
	bb.Username = "ping.zhang"
	bb.Password = "Abcd1234!"
}

func (c *hystrixClientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service()+"."+req.Endpoint()
	fmt.Println("调用的 方法名: ", cmdName)
	hystrix.ConfigureCommand(cmdName,
		hystrix.CommandConfig{
			Timeout:                100,
			MaxConcurrentRequests:  10,
			RequestVolumeThreshold: 1,
			SleepWindow:            5,
			ErrorPercentThreshold:  10,
		},
	)
	return hystrix.Do(
		cmdName,
		func() error {
			return c.Client.Call(ctx, req, rsp, opts...)
		},nil)
}

func NewHystrixClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &hystrixClientWrapper{c}
	}
}
