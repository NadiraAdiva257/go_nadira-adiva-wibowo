package route

import (
	"clean-architecture/common"
	"clean-architecture/middleware"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func StartRoute(handler common.Handler) *echo.Echo {
	e := echo.New()

	jwtMiddleware := echojwt.JWT([]byte(middleware.SECRET_JWT))

	e.POST("/users/regis", handler.UserHandler.CreateUser)
	e.POST("/users/login", handler.UserHandler.LoginUser)

	userGroup := e.Group("/users")
	userGroup.Use(jwtMiddleware)
	userGroup.GET("", handler.UserHandler.GetAllUsers)

	return e
}
