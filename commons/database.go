package commons

import (
	"fmt"
	"github.com/curder/go-gin-demo/models"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() (db *gorm.DB) {
	var (
		driverName string
		host       string
		port       int
		database   string
		user       string
		password   string
		charset    string

		args string
		err  error
	)

	driverName = "mysql"
	host = "localhost"
	port = 33060
	database = "go_gin_demo"
	user = "root"
	password = "root"
	charset = "utf8"
	args = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		user,
		password,
		host,
		port,
		database,
		charset,
	)
	if db, err = gorm.Open(driverName, args); err != nil {
		panic("failed to connect database,err: " + err.Error())
	}

	// 设置引擎和字符集
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")

	// 自动创建数据表
	db.AutoMigrate(&models.Users{})

	DB = db
	return
}

// 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
