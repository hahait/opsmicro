package jaeger

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func JaegerInit(sname string) (opentracing.Tracer, io.Closer){
	cfg := jaegercfg.Configuration{
		ServiceName:         sname,
		Sampler:            &jaegercfg.SamplerConfig{
			Type:                     jaeger.SamplerTypeConst,
			Param:                    1,
		},
		Reporter:            &jaegercfg.ReporterConfig{
			LogSpans:                   true,
			BufferFlushInterval: time.Second,
			//LocalAgentHostPort: "127.0.0.1:6831",
			//CollectorEndpoint: "http://127.0.0.1:14268/api/traces",
		},
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		fmt.Println("实例化 Tracer 对象失败...")
	}
	return tracer, closer
}
