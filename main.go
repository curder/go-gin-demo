package main

import (
	"github.com/curder/go-gin-demo/commons"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	var (
		r  *gin.Engine
		db *gorm.DB
	)

	// 初始化数据库连接
	db = commons.InitDB()
	defer db.Close()

	r = gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
