package controller

import (
	"middleware/config"
	"middleware/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	var books []model.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	bookById := config.DB.First(&books, id)

	if err := bookById.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"book":    bookById,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	var books []model.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	bookById := config.DB.Delete(&books, id)

	if err := bookById.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book by id",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	var books []model.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	title := c.FormValue("title")
	author := c.FormValue("author")
	publisher := c.FormValue("publisher")

	updateBookById := config.DB.Model(&books).Where("id = ?", id).Updates(model.Book{Title: title, Author: author, Publisher: publisher})

	if err := updateBookById.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update book by id",
	})
}
