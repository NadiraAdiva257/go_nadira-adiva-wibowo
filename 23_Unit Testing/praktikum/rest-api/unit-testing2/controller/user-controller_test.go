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

var dataUser = model.User{
	Name:     "nadira",
	Email:    "nadira123@gmail.com",
	Password: "nadira123",
}

var dataUsers = []model.User{
	{
		Name:     "nadira",
		Email:    "nadira123@gmail.com",
		Password: "nadira123",
	},
	{
		Name:     "adiva",
		Email:    "adiva123@gmail.com",
		Password: "adiva123",
	},
}

func TestCreateUserController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("CreateUserController", &dataUser).Return(nil)

	e := echo.New()

	bDataUser, _ := json.Marshal(dataUser)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bDataUser))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateUserController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateUserController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("UpdateUserController", &dataUser, 1).Return(nil)

	e := echo.New()

	bDataUser, _ := json.Marshal(dataUser)
	req := httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewReader(bDataUser))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	UpdateUserController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteUserController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("DeleteUserController", &dataUser, 1).Return(nil)

	e := echo.New()

	bDataUser, _ := json.Marshal(dataUser)
	req := httptest.NewRequest(http.MethodDelete, "/users/1", bytes.NewReader(bDataUser))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	DeleteUserController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUserController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("GetUserController", &dataUser, 1).Return(nil)

	e := echo.New()

	bDataUser, _ := json.Marshal(dataUser)
	req := httptest.NewRequest(http.MethodGet, "/users/1", bytes.NewReader(bDataUser))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetUserController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUsersController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("GetUsersController", &dataUsers).Return(nil)

	e := echo.New()

	bDataUsers, _ := json.Marshal(dataUsers)
	req := httptest.NewRequest(http.MethodGet, "/users", bytes.NewReader(bDataUsers))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetUsersController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestLoginUserController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("LoginUserController", &dataUser).Return(nil)

	e := echo.New()

	bDataUser, _ := json.Marshal(dataUser)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bDataUser))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	LoginUserController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}
