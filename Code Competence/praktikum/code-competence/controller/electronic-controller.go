package controller

import (
	"code-competence/config"
	"code-competence/middleware"
	"code-competence/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateElectronicController(c echo.Context) error {
	var electronic model.Electronic
	var electronics []model.Electronic
	var sameName bool

	electronic.AdminID = middleware.GetClaims(c).Id
	c.Bind(&electronic)

	if err := config.DB.Find(&electronics).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, value := range electronics {
		if electronic.Name == value.Name {
			sameName = true
			break
		}
	}

	if sameName {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "item name is same with other",
		})
	} else {
		if err := config.DB.Save(&electronic).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new electronic data",
		})
	}
}

func UpdateElectronicController(c echo.Context) error {
	var electronic model.Electronic
	electronic.AdminID = middleware.GetClaims(c).Id
	c.Bind(&electronic)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := config.DB.Where("id = ?", id).Updates(&electronic).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update electronic data by id",
	})
}

func DeleteElectronicController(c echo.Context) error {
	var electronics []model.Electronic

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := config.DB.Where("id = ?", id).Delete(&electronics).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete electronic data by id",
	})
}

func GetElectronicController(c echo.Context) error {
	if item_name := c.QueryParam("keyword"); item_name != "" {
		GetElectronicByItemNameController(c)
	} else {
		GetAllElectronicController(c)
	}

	return nil
}

func GetAllElectronicController(c echo.Context) error {
	var electronics []model.Electronic

	if err := config.DB.Find(&electronics).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success get all electronic data",
		"electronic data": electronics,
	})
}

func GetElectronicByItemNameController(c echo.Context) error {
	var electronics []model.Electronic

	item_name := c.QueryParam("keyword")

	if err := config.DB.Where("name LIKE ?", "%"+item_name+"%").Find(&electronics).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success get electronic data by item name",
		"electronic data": electronics,
	})
}

func GetElectronicByIdController(c echo.Context) error {
	var electronics []model.Electronic

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := config.DB.Where("id = ?", id).Find(&electronics).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success get electronic data by id",
		"electronic data": electronics,
	})
}

func GetElectronicByCategoryIdController(c echo.Context) error {
	var electronics []model.Electronic

	category_id := c.Param("category_id")

	if err := config.DB.Where("category = ?", category_id).Find(&electronics).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success get electronic data by category id",
		"electronic data": electronics,
	})
}
