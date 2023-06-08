package controller

import (
	"net/http"
	"strconv"
	"unit-testing2/model"
	"unit-testing2/service"

	"github.com/labstack/echo"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := service.GetBookRepository().GetBooksController(&books); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	var book model.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := service.GetBookRepository().GetBookController(&book, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"book":    book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	var book model.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if err := service.GetBookRepository().CreateBookController(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	var book model.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if err := service.GetBookRepository().UpdateBookController(&book, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update book by id",
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	var book model.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := service.GetBookRepository().DeleteBookController(&book, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book by id",
	})
}
