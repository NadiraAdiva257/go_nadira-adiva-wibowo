package controller

import (
	"net/http"
	"orm-mvc-eksplorasi/config"
	"orm-mvc-eksplorasi/model"
	"strconv"

	"github.com/labstack/echo"
)

// get all blogs
func GetBlogsController(c echo.Context) error {
	var blogs []model.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	var blogs []model.Blog

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	blogById := config.DB.First(&blogs, id)

	if err := blogById.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by id",
		"blog":    blogs,
	})
}

// create new blog
func CreateBlogController(c echo.Context) error {
	blog := model.Blog{}
	c.Bind(&blog)

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	var blogs []model.Blog

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	blogById := config.DB.Delete(&blogs, id)

	if err := blogById.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog by id",
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	var blogs []model.Blog

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	title := c.FormValue("title")
	content := c.FormValue("content")
	user_id, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		return err
	}

	updateBlogById := config.DB.Model(&blogs).Where("id = ?", id).Updates(model.Blog{Title: title, Content: content, UserID: uint(user_id)})

	if err := updateBlogById.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update blog by id",
	})
}
