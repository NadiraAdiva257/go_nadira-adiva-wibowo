package route

import (
	"docker/controller"

	"docker/constants"
	m "docker/pack-middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	m.LogMiddleware(e)

	e.POST("/user", controller.CreateUserController)
	e.POST("/login", controller.LoginUserController)

	eJwt := e.Group("/users")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("", controller.GetUsersController)
	eJwt.GET("/:id", controller.GetUserController)
	eJwt.DELETE("/:id", controller.DeleteUserController)
	eJwt.PUT("/:id", controller.UpdateUserController)

	eJwt2 := e.Group("/books")
	eJwt2.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt2.GET("", controller.GetBooksController)
	eJwt2.GET("/:id", controller.GetBookController)
	eJwt2.POST("", controller.CreateBookController)
	eJwt2.DELETE("/:id", controller.DeleteBookController)
	eJwt2.PUT("/:id", controller.UpdateBookController)

	return e
}
