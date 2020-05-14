package main

import (
	"github.com/curder/go-gin-demo/commons"
	"github.com/curder/go-gin-demo/configs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func main() {
	var (
		r    *gin.Engine
		db   *gorm.DB
		port string
	)
	// 初始化项目配置
	configs.InitConfig()

	// 初始化数据库连接
	db = commons.InitDB()
	defer db.Close()

	r = gin.Default()
	r = CollectRoute(r)
	port = viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}
