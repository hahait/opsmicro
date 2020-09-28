package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	mw "ops.was.ink/opsmicro/gateway/middlewares"
)

var router *gin.Engine

func init() {
	// 1. 实例化一个 gin 对象
	router = gin.Default()
	// 2. 加载中间件
	router.Use(mw.HttpErrorHandler, mw.Authenticate, mw.JaegerHandler)
	// 3. 定义未知路由处理 handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 1,
			"errmsg": "请求的路由不存在",
		})
	})
}

func GetRouter() *gin.Engine {
	return router
}
