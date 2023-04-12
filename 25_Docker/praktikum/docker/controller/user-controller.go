package controller

import (
	"net/http"
	"strconv"

	"docker/config"
	"docker/model"
	m "docker/pack-middleware"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	var users []model.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	userById := config.DB.First(&users, id)

	if err := userById.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	var users []model.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	userById := config.DB.Delete(&users, id)

	if err := userById.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	var users []model.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	updateUserById := config.DB.Model(&users).Where("id = ?", id).Updates(model.User{Name: name, Email: email, Password: password})

	if err := updateUserById.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user by id",
	})
}

func LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "fail login",
			"error":    err.Error(),
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
