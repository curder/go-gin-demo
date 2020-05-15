package repositories

import (
	"github.com/curder/go-gin-demo/commons"
	"github.com/curder/go-gin-demo/models"
	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

// 构造函数
func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{DB: commons.GetDB()}
}

// 创建
func (c CategoryRepository) Create(name string) (category *models.Category, err error) {
	category = &models.Category{Name: name}

	if err = c.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

// 更新分类名称
func (c CategoryRepository) Update(category models.Category, name string) (*models.Category, error) {
	var (
		err error
	)

	if err = c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// 通过ID查找分类
func (c CategoryRepository) FindByID(id int) (*models.Category, error) {
	var (
		category models.Category
		err      error
	)

	if err = c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// 通过ID删除分类
func (c CategoryRepository) DeleteByID(id int) error {
	var (
		category models.Category
		err      error
	)

	if err = c.DB.Delete(&category, id).Error; err != nil {
		return err
	}

	return nil
}
