package usecase

import (
	"clean-architecture/entity/user"
	ur "clean-architecture/repo"
)

type UserUsecase interface {
	GetAllUsers(users *[]user.User) error
	CreateUser(user *user.User) error
	LoginUser(user *user.User) (string, error)
}

type userUseCase struct {
	userRepo ur.UserRepo
}

func New(userRepo ur.UserRepo) *userUseCase {
	return &userUseCase{
		userRepo,
	}
}
