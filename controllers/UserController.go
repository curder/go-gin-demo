package controllers

import (
	"github.com/curder/go-gin-demo/commons"
	"github.com/curder/go-gin-demo/models"
	"github.com/curder/go-gin-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册
func Register(context *gin.Context) {
	var (
		name     string
		phone    string
		password string
		user     models.Users
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
		name = utils.RandomString(12)
	}

	// 判断手机号是否存在
	if models.IsPhoneExists(commons.DB, phone) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "用户已经存在"})
		return
	}

	// 创建用户
	user = models.Users{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	commons.DB.Create(&user)

	// 返回结果
	context.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}
