package main

import (
	"github.com/curder/go-gin-demo/controllers"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controllers.Register)

	return r
}
