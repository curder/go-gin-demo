package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Users struct {
	gorm.Model

	Name     string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"size:255;not null"`
}

func main() {

	var (
		r  *gin.Engine
		db *gorm.DB
	)

	// 初始化数据库连接
	db = InitDB()

	defer db.Close()

	r = gin.Default()

	r.POST("/api/auth/register", func(context *gin.Context) {
		var (
			name     string
			phone    string
			password string
			user     Users
		)
		// 获取参数
		name = context.PostForm("name")
		phone = context.PostForm("phone")
		password = context.PostForm("password")

		// 验证数据
		if len(phone) != 11 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "手机号必须是11位数"})
			return
		}

		if len(password) < 4 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "密码不能小于4位"})
			return
		}

		if len(name) == 0 {
			name = RandomString(12)
		}

		log.Println(name, phone, password)
		// 判断手机号是否存在
		if IsPhoneExists(db, phone) {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "用户已经存在"})
			return
		}

		// 创建用户
		user = Users{
			Name:     name,
			Phone:    phone,
			Password: password,
		}
		db.Create(&user)

		// 返回结果
		context.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
		})
	})

	panic(r.Run())
}

func RandomString(n int) string {
	var (
		letters []byte
		index   int
		result  []byte
	)

	letters = []byte("abcdefhijklmnopqrstuvwxyzABCDEFHIJKLMNOPQRSTUVWXYZ1234597890")
	result = make([]byte, n)
	rand.Seed(time.Now().Unix())

	for index = range result {
		result[index] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

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
	db.AutoMigrate(&Users{})

	return
}

// 判断手机号是否存在
func IsPhoneExists(db *gorm.DB, phone string) bool {
	var user Users

	db.Where("phone = ?", phone).First(&user)

	if user.ID != 0 {
		return true
	}

	return false
}
