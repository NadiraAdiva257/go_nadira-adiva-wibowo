package controller

import (
	"code-competence/config"
	"code-competence/middleware"
	"code-competence/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateAdminController(c echo.Context) error {
	var admin model.Admin
	var admins []model.Admin
	var sameEmailUsername bool

	c.Bind(&admin)

	if err := config.DB.Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, value := range admins {
		if admin.Email == value.Email || admin.Username == value.Username {
			sameEmailUsername = true
			break
		}
	}

	if sameEmailUsername {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "email or username admin is same with other",
		})
	} else {
		if err := config.DB.Save(&admin).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new admin",
		})
	}
}

func LoginAdminController(c echo.Context) error {
	var admin model.Admin
	var admins model.Admin
	var idAdmin int

	c.Bind(&admin)

	cekUsername := config.DB.First(&admins, "username = ?", admin.Username)
	if err := cekUsername.Error; err != nil {
		return echo.ErrUnauthorized
	} else {
		if admins.Password != admin.Password {
			return echo.ErrUnauthorized
		} else {
			idAdmin = int(admins.ID)
		}
	}

	token, err := middleware.CreateToken(idAdmin, admin.Username, admin.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "admin fail login",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "admin successfully login",
		"your token": token,
	})
}
