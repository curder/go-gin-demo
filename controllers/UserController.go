package controllers

import (
	"github.com/curder/go-gin-demo/commons"
	"github.com/curder/go-gin-demo/models"
	"github.com/curder/go-gin-demo/resources"
	"github.com/curder/go-gin-demo/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// 注册
func Register(ctx *gin.Context) {
	var (
		name           string
		phone          string
		password       string
		user           models.Users
		hashedPassword []byte
		err            error
	)
	// 获取参数
	name = ctx.PostForm("name")
	phone = ctx.PostForm("phone")
	password = ctx.PostForm("password")

	// 验证数据
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "手机号必须是11位数"})
		return
	}

	if len(password) < 4 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "密码不能小于4位"})
		return
	}

	if len(name) == 0 {
		name = utils.RandomString(12)
	}

	// 判断手机号是否存在
	if models.IsPhoneExists(commons.DB, phone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "用户已经存在"})
		return
	}

	// 创建用户
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "密码加密错误"})
		return
	}
	user = models.Users{
		Name:     name,
		Phone:    phone,
		Password: string(hashedPassword),
	}
	commons.DB.Create(&user)

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "注册成功",
		"data":    gin.H{},
	})
}

// 用户登录
func Login(ctx *gin.Context) {
	var (
		phone    string
		password string
		user     models.Users
		err      error
		token    string
	)
	// 获取参数
	phone = ctx.PostForm("phone")
	password = ctx.PostForm("password")

	// 数据验证
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "手机号必须是11位数"})
		return
	}

	if len(password) < 4 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "密码不能小于4位"})
		return
	}

	// 判断手机号是否存在
	commons.GetDB().Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "用户不存在"})
		return
	}

	// 判断密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "密码错误"})
		return
	}

	// 生成token
	if token, err = commons.ReleaseToken(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "用户token生成失败"})
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "登录成功",
		"data":    gin.H{"token": token},
	})
}

// 用户信息
func UserInfo(ctx *gin.Context) {
	var (
		user   interface{}
		exists bool
	)
	if user, exists = ctx.Get(`user`); !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "没有操作权限"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": gin.H{"user": resources.ToUserResource(user.(models.Users))}})
}
