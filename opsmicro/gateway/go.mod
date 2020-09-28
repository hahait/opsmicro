module ops.was.ink/opsmicro/gateway

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
