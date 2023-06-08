package controller

import (
	"net/http"
	"strconv"

	"unit-testing2/model"
	m "unit-testing2/pack-middleware"
	"unit-testing2/service"

	"github.com/labstack/echo"
)

type Controller struct {
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := service.GetUserRepository().GetUsersController(&users); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	var user model.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := service.GetUserRepository().GetUserController(&user, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if err := service.GetUserRepository().CreateUserController(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	var user model.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if err := service.GetUserRepository().UpdateUserController(&user, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user by id",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	var user model.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := service.GetUserRepository().DeleteUserController(&user, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
	})
}

func LoginUserController(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if err := service.GetUserRepository().LoginUserController(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	token, err := m.CreateToken(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "fail login",
			"error":    err.Error(),
		})
	}

	userResponse := model.UserResponse{Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success login user",
		"user":     userResponse,
	})
}
