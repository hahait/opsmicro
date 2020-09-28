package handler

import (
	"github.com/gin-gonic/gin"
	"ops.was.ink/opsmicro/gateway/handler/auth"
	"ops.was.ink/opsmicro/gateway/routers"
)

var r = routers.GetRouter()

type Option func(*gin.Engine)

var options = []Option{
	auth.Routers,
}

func init() {
	for _, opt := range options {
		opt(r)
	}
}
