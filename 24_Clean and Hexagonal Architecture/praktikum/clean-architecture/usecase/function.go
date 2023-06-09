package usecase

import (
	"clean-architecture/entity/user"
	"clean-architecture/middleware"
)

func (uc *userUseCase) GetAllUsers(users *[]user.User) error {
	return uc.userRepo.GetAllUsers(users)
}

func (uc *userUseCase) CreateUser(user *user.User) error {
	return uc.userRepo.CreateUser(user)
}

func (uc *userUseCase) LoginUser(user *user.User) (string, error) {
	if err := uc.userRepo.LoginUser(user); err != nil {
		return "", err
	}

	token, err := middleware.CreateToken(int(user.ID), user.Email, user.Password)
	if err != nil {
		return "", err
	}

	return token, nil
}
