package main

import (
	"github.com/curder/go-gin-demo/controllers"
	"github.com/curder/go-gin-demo/middlewares"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	var (
		categoryRouters    *gin.RouterGroup
		categoryController controllers.CategoryControllerInterface
	)

	r.Use(middlewares.CorsMiddleware(), middlewares.RecoveryMiddleware())       // 跨域和Panic处理中间件
	r.POST("/api/auth/register", controllers.Register)                          // 用户注册
	r.POST("/api/auth/login", controllers.Login)                                // 用户登录
	r.GET("/api/auth/info", middlewares.AuthMiddleware(), controllers.UserInfo) // 用户信息

	categoryRouters = r.Group("/api/categories")              // 路由分组
	categoryController = controllers.NewCategoryController()  // 获取分类控制器实例
	categoryRouters.GET("", categoryController.Index)         // 分类列表
	categoryRouters.POST("", categoryController.Create)       // 创建分类
	categoryRouters.PUT("/:id", categoryController.Update)    // 更新分类
	categoryRouters.GET("/:id", categoryController.Show)      // 查看分类
	categoryRouters.DELETE("/:id", categoryController.Delete) // 删除分类

	return r
}
