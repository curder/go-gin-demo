package main

import (
	"github.com/curder/go-gin-demo/controllers"
	"github.com/curder/go-gin-demo/middlewares"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middlewares.CorsMiddleware())
	r.POST("/api/auth/register", controllers.Register)                          // 用户注册
	r.POST("/api/auth/login", controllers.Login)                                // 用户登录
	r.GET("/api/auth/info", middlewares.AuthMiddleware(), controllers.UserInfo) // 用户信息

	return r
}
