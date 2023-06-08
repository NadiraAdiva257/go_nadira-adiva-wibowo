package service

import (
	"errors"
	"unit-testing2/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (urm *UserRepositoryMock) CreateUserController(user *model.User) error {
	if user == nil {
		return errors.New("error")
	} else {
		return nil
	}
}

func (urm *UserRepositoryMock) UpdateUserController(user *model.User, id int) error {
	var users []model.User
	var err error

	for _, user := range users {
		if int(user.ID) == id {
			err = nil
			break
		}
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (urm *UserRepositoryMock) DeleteUserController(user *model.User, id int) error {
	var users []model.User
	var err error

	for _, user := range users {
		if int(user.ID) == id {
			err = nil
			break
		}
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (urm *UserRepositoryMock) GetUserController(user *model.User, id int) error {
	var users []model.User
	var err error

	for _, user := range users {
		if int(user.ID) == id {
			err = nil
			break
		}
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (urm *UserRepositoryMock) GetUsersController(users *[]model.User) error {
	var err error

	for _, user := range *users {
		if user.Email != "" {
			err = nil
		}
		if user.Password != "" {
			err = nil
		}
		if user.Name != "" {
			err = nil
		}
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (urm *UserRepositoryMock) LoginUserController(user *model.User) error {
	if user == nil {
		return errors.New("error")
	} else {
		return nil
	}
}
