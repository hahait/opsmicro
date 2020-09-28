package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"strings"
)

func JaegerHandler(c *gin.Context) {
	var md = make(map[string]string)
	r_url := strings.Split(c.Request.RequestURI, "?")[0]
	spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span := opentracing.GlobalTracer().StartSpan(r_url, opentracing.ChildOf(spanCtx))
	defer span.Finish()
	if err := opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md)); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx := opentracing.ContextWithSpan(context.Background(),span)
	ctx = metadata.NewContext(ctx, md)
	c.Set("traceContext", ctx)
	c.Next()
}
