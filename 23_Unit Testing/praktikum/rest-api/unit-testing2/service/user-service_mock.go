package service

import (
	"errors"
	"unit-testing2/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (urm *UserRepositoryMock) CreateUser(user *model.User) error {
	if user == nil {
		return errors.New("error")
	} else {
		return nil
	}
}
