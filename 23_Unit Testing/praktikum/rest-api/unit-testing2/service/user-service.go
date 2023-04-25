package service

import (
	"unit-testing2/config"
	"unit-testing2/model"
)

type IUserService interface {
	CreateUser(*model.User) error
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

func (u *UserRepository) CreateUser(user *model.User) error {
	err := config.DB.Save(&user)
	if err != nil {
		return err.Error
	}

	return nil
}
