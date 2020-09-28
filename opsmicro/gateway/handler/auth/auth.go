package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/micro/go-micro/v2/errors"
	"github.com/opentracing/opentracing-go"
	"net/http"
	authproto "ops.was.ink/opsmicro/gateway/proto/auth"
	"ops.was.ink/opsmicro/gateway/utils"
)

type Login struct {
	Username string
	Password string
}

func LoginHandler(c *gin.Context) {
	var (
		lg Login
	)
	contx := c.MustGet("traceContext").(context.Context)
	span, ctx := opentracing.StartSpanFromContext(contx,"loginHandler")
	defer span.Finish()
	utils.ErrorHandler(c.ShouldBindBodyWith(&lg, binding.JSON), http.StatusBadRequest, "获取用户登陆信息失败")
	// RPC 请求 auth 服务认证用户
	rsp, err := AuthClient.Login(ctx, &authproto.Request{Username: lg.Username, Password: lg.Password})
	if err != nil {
		eor := errors.Parse(err.Error())
		utils.ErrorHandler(eor, eor.Code, "用户认证失败")
	}
	c.SetCookie("Bearer", rsp.Token, 3600,"/", "*", false, true)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": "认证成功",
	})
}

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"results": "Haha, hello world",
	})
}