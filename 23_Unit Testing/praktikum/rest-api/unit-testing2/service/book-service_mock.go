package service

import (
	"errors"
	"unit-testing2/model"

	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	Mock mock.Mock
}

func (brm *BookRepositoryMock) CreateBookController(book *model.Book) error {
	if book == nil {
		return errors.New("error")
	} else {
		return nil
	}
}

func (brm *BookRepositoryMock) UpdateBookController(book *model.Book, id int) error {
	var books []model.Book
	var err error

	for _, book := range books {
		if int(book.ID) == id {
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

func (brm *BookRepositoryMock) DeleteBookController(book *model.Book, id int) error {
	var books []model.Book
	var err error

	for _, book := range books {
		if int(book.ID) == id {
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

func (brm *BookRepositoryMock) GetBookController(book *model.Book, id int) error {
	var books []model.Book
	var err error

	for _, book := range books {
		if int(book.ID) == id {
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

func (brm *BookRepositoryMock) GetBooksController(books *[]model.Book) error {
	var err error

	for _, book := range *books {
		if book.Title != "" {
			err = nil
		}
		if book.Author != "" {
			err = nil
		}
		if book.Publisher != "" {
			err = nil
		}
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}
