package controllers

import (
	"github.com/curder/go-gin-demo/commons"
	"github.com/curder/go-gin-demo/models"
	"github.com/curder/go-gin-demo/resources"
	"github.com/curder/go-gin-demo/responses"
	"github.com/curder/go-gin-demo/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// 注册
func Register(ctx *gin.Context) {
	var (
		requestUser    models.Users
		name           string
		phone          string
		password       string
		user           models.Users
		hashedPassword []byte
		err            error
		token          string
	)

	// 获取参数
	_ = ctx.Bind(&requestUser)
	name = requestUser.Name
	phone = requestUser.Phone
	password = requestUser.Password

	// 验证数据
	if len(phone) != 11 {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, nil, "手机号必须是11位数")
		return
	}

	if len(password) < 4 {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, nil, "密码不能小于4位")
		return
	}

	if len(name) == 0 {
		name = utils.RandomString(12)
	}

	// 判断手机号是否存在
	if models.IsPhoneExists(commons.DB, phone) {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, nil, "用户已经存在")
		return
	}

	// 创建用户
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		responses.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, nil, "密码加密错误")
		return
	}
	user = models.Users{
		Name:     name,
		Phone:    phone,
		Password: string(hashedPassword),
	}
	commons.DB.Create(&user)

	// 生成token
	if token, err = commons.ReleaseToken(user); err != nil {
		responses.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, nil, "用户token生成失败")
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	responses.Success(ctx, gin.H{"token": token}, "注册成功")
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
	_ = ctx.Bind(&user)
	phone = user.Phone
	password = user.Password

	// 数据验证
	if len(phone) != 11 {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, nil, "手机号必须是11位数")
		return
	}

	if len(password) < 4 {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, nil, "密码不能小于4位")
		return
	}

	// 判断手机号是否存在
	commons.GetDB().Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, nil, "用户不存在")
		return
	}

	// 判断密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		responses.Fail(ctx, "用户不存在", nil)
		return
	}

	// 生成token
	if token, err = commons.ReleaseToken(user); err != nil {
		responses.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, nil, "用户token生成失败")
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	responses.Success(ctx, gin.H{"token": token}, "登录成功")
}

// 用户信息
func UserInfo(ctx *gin.Context) {
	var (
		user   interface{}
		exists bool
	)
	if user, exists = ctx.Get(`user`); !exists {
		responses.Response(ctx, http.StatusUnauthorized, http.StatusUnauthorized, nil, "没有操作权限")
		return
	}

	responses.Success(ctx, gin.H{"user": resources.ToUserResource(user.(models.Users))}, "查询成功")
}
