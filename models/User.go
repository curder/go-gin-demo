package models

import "github.com/jinzhu/gorm"

// 定义用户结构体
type Users struct {
	gorm.Model

	Name     string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"size:255;not null"`
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
