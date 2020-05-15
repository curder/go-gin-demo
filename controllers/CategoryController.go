package controllers

import (
	"fmt"
	"github.com/curder/go-gin-demo/commons"
	"github.com/curder/go-gin-demo/models"
	"github.com/curder/go-gin-demo/responses"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type CategoryControllerInterface interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

// 初始化
func NewCategoryController() CategoryControllerInterface {
	var (
		db *gorm.DB
	)
	db = commons.GetDB()

	// 设置数据库迁移
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
	db.AutoMigrate(&models.Category{})

	return CategoryController{DB: db}
}

// 列表
func (c CategoryController) Index(ctx *gin.Context) {
	panic("implement me")
}

// 添加
func (c CategoryController) Store(ctx *gin.Context) {
	var (
		requestCategory models.Category
		err             error
	)

	if err = ctx.Bind(&requestCategory); err != nil {
		fmt.Printf("request bind err: %s", err.Error())
		return
	}

	if requestCategory.Name == "" {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, gin.H{}, "数据验证错误，分类名称必须提供")
		return
	}

	c.DB.Create(&requestCategory)

	responses.Success(ctx, gin.H{"category": requestCategory}, "创建分类成功")
}

// 更新
func (c CategoryController) Update(ctx *gin.Context) {
	var (
		requestCategory models.Category
		categoryID      int
		updateCategory  models.Category
		err             error
	)

	// 获取请求主体中的参数
	_ = ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, gin.H{}, "数据验证错误，分类名称必须提供")
		return
	}

	// 获取请求地址中的参数
	if categoryID, err = strconv.Atoi(ctx.Params.ByName("id")); err != nil {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, gin.H{}, "请求的参数有误")
		return
	}

	// 数据不存在
	if c.DB.First(&updateCategory, categoryID).RecordNotFound() {
		responses.Fail(ctx, "分类不存在", nil)
	}

	// 更新数据
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	// 返回响应
	responses.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

// 展示
func (c CategoryController) Show(ctx *gin.Context) {
	var (
		category   models.Category
		categoryID int
		err        error
	)

	// 获取请求地址中的参数
	if categoryID, err = strconv.Atoi(ctx.Params.ByName("id")); err != nil {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, gin.H{}, "请求的参数有误")
		return
	}

	// 数据不存在
	if c.DB.First(&category, categoryID).RecordNotFound() {
		responses.Fail(ctx, "分类不存在", nil)
	}

	// 返回响应
	responses.Success(ctx, gin.H{"category": category}, "查询分类成功")
}

// 删除
func (c CategoryController) Delete(ctx *gin.Context) {
	var (
		category   models.Category
		categoryID int
		err        error
	)

	// 获取请求地址中的参数
	if categoryID, err = strconv.Atoi(ctx.Params.ByName("id")); err != nil {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, gin.H{}, "请求的参数有误")
		return
	}

	if err = c.DB.Delete(&category, categoryID).Error; err != nil {
		responses.Fail(ctx, "删除失败，请稍后重试", nil)
	}

	responses.Success(ctx, nil, "删除成功")
}
