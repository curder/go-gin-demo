package models

// 文章分类模型
type Category struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"type:varchar(100);not null;unique"`
	CreatedAt Time   `json:"created_at" gorm:"type:datetime"`
	UpdatedAt Time   `json:"updated_at" gorm:"type:datetime"`
	DeletedAt *Time  `json:"deleted_at" gorm:"type:datetime" sql:"index"`
}
