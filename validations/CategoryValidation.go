package validations

type CreateCategoryValidation struct {
	Name string `form:"name" json:"name" binding:"required"`
}
