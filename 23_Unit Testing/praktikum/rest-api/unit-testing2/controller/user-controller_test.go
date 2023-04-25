package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"unit-testing2/model"
	"unit-testing2/service"

	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func InitEchoTestAPI() *echo.Echo {
// 	config.InitDB()
// 	e := echo.New()
// 	return e
// }

// func InsertDataUserForGetUsers() error {
// 	user := model.User{
// 		Name:     "eline",
// 		Email:    "eline789@gmail.com",
// 		Password: "eline789",
// 	}

// 	err := config.DB.Save(&user).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func TestController_CreateUserController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		m       *Controller
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if err := tt.m.CreateUserController(tt.args.c); (err != nil) != tt.wantErr {
			// 	t.Errorf("Controller.CreateUserController() error = %v, wantErr %v", err, tt.wantErr)
			// }

			data := model.User{
				Name:     "eline",
				Email:    "eline789@gmail.com",
				Password: "eline789",
			}
			userRepository.Mock.On("CreateUser", &data).Return(errors.New("error"))

			e := echo.New()

			bData, _ := json.Marshal(data)
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bData))
			req.Header.Set("content-type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			controller := Controller{}
			controller.CreateUserController(c)

			assert.Equal(t, http.StatusOK, rec.Code)

			var resultJSON map[string]interface{}
			json.Unmarshal(rec.Body.Bytes(), &resultJSON)

			expectResult := map[string]interface{}{
				"message": "success create new user",
			}
			assert.Equal(t, expectResult, resultJSON)
		})
	}
}
