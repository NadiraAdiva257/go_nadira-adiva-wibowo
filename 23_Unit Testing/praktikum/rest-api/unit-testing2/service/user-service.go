package service

import (
	"unit-testing2/config"
	"unit-testing2/model"
)

type IUserService interface {
	CreateUserController(user *model.User) error
	UpdateUserController(user *model.User, id int) error
	DeleteUserController(user *model.User, id int) error
	GetUserController(user *model.User, id int) error
	GetUsersController(users *[]model.User) error
	LoginUserController(user *model.User) error
}

type UserRepository struct {
	Func IUserService
}

var userRepository IUserService

func init() {
	ur := &UserRepository{}
	ur.Func = ur

	userRepository = ur
}

func GetUserRepository() IUserService {
	return userRepository
}

func SetUserRepository(ur IUserService) {
	userRepository = ur
}

func (u *UserRepository) CreateUserController(user *model.User) error {
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) UpdateUserController(user *model.User, id int) error {
	if err := config.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) DeleteUserController(user *model.User, id int) error {
	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUserController(user *model.User, id int) error {
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUsersController(users *[]model.User) error {
	if err := config.DB.Find(&users).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) LoginUserController(user *model.User) error {
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}

	return nil
}
