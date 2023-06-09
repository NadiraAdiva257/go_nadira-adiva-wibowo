package app

import (
	"clean-architecture/cmd/route"
	"clean-architecture/common"
	"clean-architecture/config"
	userHandler "clean-architecture/handler/api"
	userRepo "clean-architecture/repo"
	userUsecase "clean-architecture/usecase"

	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	config.Init()

	userRepo := userRepo.New(config.DB)
	userUsecase := userUsecase.New(userRepo)
	userHandler := userHandler.New(userUsecase)

	handler := common.Handler{
		UserHandler: userHandler,
	}

	router := route.StartRoute(handler)
	return router
}
