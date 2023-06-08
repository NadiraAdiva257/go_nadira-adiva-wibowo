package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"unit-testing2/model"
	"unit-testing2/service"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var dataBook = model.Book{
	Title:     "sunset bersama rosie",
	Author:    "tere liye",
	Publisher: "gramedia",
}

var dataBooks = []model.Book{
	{
		Title:     "sunset bersama rosie",
		Author:    "tere liye",
		Publisher: "gramedia",
	},
	{
		Title:     "bulan",
		Author:    "tere liye",
		Publisher: "gramedia",
	},
}

func TestCreateBookController(t *testing.T) {
	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("CreateBookController", &dataBook).Return(nil)

	e := echo.New()

	bDataBook, _ := json.Marshal(dataBook)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(bDataBook))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateBookController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateBookController(t *testing.T) {
	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("UpdateBookController", &dataBook, 1).Return(nil)

	e := echo.New()

	bDataBook, _ := json.Marshal(dataBook)
	req := httptest.NewRequest(http.MethodPut, "/books/1", bytes.NewReader(bDataBook))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	UpdateBookController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteBookController(t *testing.T) {
	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("DeleteBookController", &dataBook, 1).Return(nil)

	e := echo.New()

	bDataBook, _ := json.Marshal(dataBook)
	req := httptest.NewRequest(http.MethodDelete, "/books/1", bytes.NewReader(bDataBook))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	DeleteBookController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetBookController(t *testing.T) {
	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("GetBookController", &dataBook, 1).Return(nil)

	e := echo.New()

	bDataBook, _ := json.Marshal(dataBook)
	req := httptest.NewRequest(http.MethodGet, "/books/1", bytes.NewReader(bDataBook))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetBookController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetBooksController(t *testing.T) {
	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("GetBooksController", &dataBooks).Return(nil)

	e := echo.New()

	bDataBooks, _ := json.Marshal(dataBooks)
	req := httptest.NewRequest(http.MethodGet, "/books", bytes.NewReader(bDataBooks))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetBooksController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}
