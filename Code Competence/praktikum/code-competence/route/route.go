package route

import (
	"code-competence/controller"
	"code-competence/middleware"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func New() *echo.Echo {
	e := echo.New()

	middleware.LogMiddleware(e)

	e.POST("/admins", controller.CreateAdminController)
	e.POST("/admins/login", controller.LoginAdminController)

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middleware.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	groupElectronic := e.Group("/items")
	groupElectronic.Use(echojwt.WithConfig(config))
	groupElectronic.POST("", controller.CreateElectronicController)
	groupElectronic.PUT("/:id", controller.UpdateElectronicController)
	groupElectronic.DELETE("/:id", controller.DeleteElectronicController)
	groupElectronic.GET("", controller.GetElectronicController)
	groupElectronic.GET("/:id", controller.GetElectronicByIdController)
	groupElectronic.GET("/category/:category_id", controller.GetElectronicByCategoryIdController)
	// groupElectronic.GET("", controller.GetElectronicByItemNameController)
	// groupElectronic.GET("", controller.GetAllElectronicController)

	return e
}
