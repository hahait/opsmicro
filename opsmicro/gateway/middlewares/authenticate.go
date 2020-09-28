package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"ops.was.ink/opsmicro/gateway/handler/auth"
	authproto "ops.was.ink/opsmicro/gateway/proto/auth"
	"strings"
)

func Authenticate(c *gin.Context) {
	req_path := strings.Split(c.Request.RequestURI, "?")[0]
	if req_path == "/login" {
		c.Next()
	} else {
		// 1. 获取 http header 中的 token
		cookie := c.Request.Header.Get("Cookie")
		tokenstr := strings.TrimPrefix(cookie, "Bearer=")
		if cookie == "" || (!strings.HasPrefix(cookie, "Bearer")) || tokenstr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 1,
				"msg": "Token 验证失败, 请重新登陆...",
			})
			return
		}
		rsp, err := auth.AuthClient.ValidateAccessToken(context.Background(), &authproto.ValidateTokenRequest{Token: tokenstr})

		if err != nil || (!rsp.Success) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": 1,
				"msg": "Token 验证失败",
				"errmsg": err.Error(),
			})
			return
		}
		c.Next()
	}
}
