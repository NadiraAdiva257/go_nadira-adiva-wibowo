package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	// ID int ``
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Blogs    []Blog `gorm:"foreignKey:UserID"`
}

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}

type Blog struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	UserID  int    `json:"id_user" gorm:"column:user_id"`
	// User    User
	// IdUser int `json:"id_user" form:"id_user" gorm:"unique; not null; foreignKey:IdUser"`
	// User    User
}
