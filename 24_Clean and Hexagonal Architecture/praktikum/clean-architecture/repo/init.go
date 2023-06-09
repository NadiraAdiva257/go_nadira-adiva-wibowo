package repo

import (
	"clean-architecture/entity/user"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetAllUsers(users *[]user.User) error
	CreateUser(user *user.User) error
	LoginUser(user *user.User) error
}

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserRepo {
	return &userRepo{
		db,
	}
}
