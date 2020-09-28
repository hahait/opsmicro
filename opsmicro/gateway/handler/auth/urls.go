package auth

import (
	"github.com/gin-gonic/gin"
)

func Routers( e *gin.Engine) {
	e.POST("/login", LoginHandler)
	e.GET("/", IndexHandler)
	//userGroup := e.Group("/user")
	//{
	//	userGroup.POST("", createUserHandler)
	//	userGroup.DELETE("", deleteUserHandler)
	//	userGroup.PUT("", updateUserHandler)
	//	userGroup.GET("", queryUserHandler)
	//	userGroup.GET("/info", queryUserInfoHandler)
	//	userGroup.PUT("/pwd", updateUserPasswordHandler)
	//	userGroup.PUT("/admin", updateUserAdminHandler)
	//	userGroup.PUT("/status", updateUserStatusHandler)
	//}
}