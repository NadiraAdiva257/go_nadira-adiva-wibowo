package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username    string `json:"username" form:"username"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Electronics []Electronic
}
