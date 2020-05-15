package controllers

import (
	"github.com/curder/go-gin-demo/models"
	"github.com/curder/go-gin-demo/repositories"
	"github.com/curder/go-gin-demo/responses"
	"github.com/curder/go-gin-demo/validations"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryControllerInterface interface {
	RestController
}

type CategoryController struct {
	Repository repositories.CategoryRepository
}

// 初始化
func NewCategoryController() CategoryControllerInterface {
	var (
		repository repositories.CategoryRepository
	)
	repository = repositories.NewCategoryRepository()

	// 设置数据库迁移
	repository.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
	repository.DB.AutoMigrate(&models.Category{})

	return CategoryController{Repository: repository}
}

// 列表
func (c CategoryController) Index(ctx *gin.Context) {
	panic("implement me")
}

// 添加
func (c CategoryController) Create(ctx *gin.Context) {
	var (
		requestCategory validations.CreateCategoryValidation
		category        *models.Category
		err             error
	)

	if err = ctx.ShouldBind(&requestCategory); err != nil {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, gin.H{}, "数据验证错误，分类名称必须提供")
		return
	}

	if category, err = c.Repository.Create(requestCategory.Name); err != nil {
		panic(err)
	}

	responses.Success(ctx, gin.H{"category": category}, "创建分类成功")
}

// 更新
func (c CategoryController) Update(ctx *gin.Context) {
	var (
		requestCategory validations.CreateCategoryValidation
		categoryID      int
		updateCategory  *models.Category
		category        *models.Category
		err             error
	)
	if err = ctx.ShouldBind(&requestCategory); err != nil {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, gin.H{}, "数据验证错误，分类名称必须提供")
		return
	}

	// 获取请求地址中的参数
	if categoryID, err = strconv.Atoi(ctx.Params.ByName("id")); err != nil {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, gin.H{}, "请求的参数有误")
		return
	}

	// 数据不存在
	if updateCategory, err = c.Repository.FindByID(categoryID); err != nil {
		responses.Fail(ctx, "分类不存在", nil)
		return
	}

	// 更新数据
	if category, err = c.Repository.Update(*updateCategory, requestCategory.Name); err != nil {
		responses.Fail(ctx, "更新数据有误", nil)
	}

	// 返回响应
	responses.Success(ctx, gin.H{"category": category}, "修改成功")
}

// 展示
func (c CategoryController) Show(ctx *gin.Context) {
	var (
		category   *models.Category
		categoryID int
		err        error
	)

	// 获取请求地址中的参数
	if categoryID, err = strconv.Atoi(ctx.Params.ByName("id")); err != nil {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, gin.H{}, "请求的参数有误")
		return
	}

	// 数据不存在
	if category, err = c.Repository.FindByID(categoryID); err != nil {
		responses.Fail(ctx, "分类不存在", nil)
		return
	}

	// 返回响应
	responses.Success(ctx, gin.H{"category": &category}, "查询分类成功")
}

// 删除
func (c CategoryController) Delete(ctx *gin.Context) {
	var (
		categoryID int
		err        error
	)

	// 获取请求地址中的参数
	if categoryID, err = strconv.Atoi(ctx.Params.ByName("id")); err != nil {
		responses.Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, gin.H{}, "请求的参数有误")
		return
	}

	if err = c.Repository.DeleteByID(categoryID); err != nil {
		responses.Fail(ctx, "删除失败，请稍后重试", nil)
		return
	}

	responses.Success(ctx, nil, "删除成功")
}
