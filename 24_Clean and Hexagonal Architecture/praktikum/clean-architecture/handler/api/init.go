package api

import uc "clean-architecture/usecase"

type UserHandler struct {
	userUseCase uc.UserUsecase
}

func New(userUseCase uc.UserUsecase) *UserHandler {
	return &UserHandler{
		userUseCase,
	}
}
