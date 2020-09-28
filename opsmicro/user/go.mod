module ops.was.ink/opsmicro/user

go 1.13

require (
	github.com/alibaba/sentinel-golang v0.6.1
	github.com/codahale/hdrhistogram v0.9.0 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/logger/zap/v2 v2.9.1 // indirect
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/zap v1.15.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.0.0
	gorm.io/gorm v1.20.0
)

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
