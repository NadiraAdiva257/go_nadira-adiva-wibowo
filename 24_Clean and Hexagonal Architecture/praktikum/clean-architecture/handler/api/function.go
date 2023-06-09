package api

import (
	"clean-architecture/entity/user"

	"github.com/labstack/echo/v4"
)

func (uh *UserHandler) GetAllUsers(c echo.Context) error {
	var users []user.User

	if err := uh.userUseCase.GetAllUsers(&users); err != nil {
		return c.JSON(500, echo.Map{
			"Message": "Fail to get all user",
		})
	}

	return c.JSON(200, echo.Map{
		"Message": "Success get all user",
		"Data":    users,
	})
}

func (uh *UserHandler) CreateUser(c echo.Context) error {
	var user user.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"Message": "Gagal",
		})
	}

	if err := uh.userUseCase.CreateUser(&user); err != nil {
		return c.JSON(500, echo.Map{
			"Message": "Fail to get all user",
		})
	}

	return c.JSON(200, echo.Map{
		"Message": "Success create user",
	})
}

func (uh *UserHandler) LoginUser(c echo.Context) error {
	var user user.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"Message": "Gagal",
		})
	}

	token, err := uh.userUseCase.LoginUser(&user)
	if err != nil {
		return c.JSON(500, echo.Map{
			"Message": "Fail to get all user",
		})
	}

	return c.JSON(200, echo.Map{
		"Message": "Login successful",
		"Token":   token,
	})
}
