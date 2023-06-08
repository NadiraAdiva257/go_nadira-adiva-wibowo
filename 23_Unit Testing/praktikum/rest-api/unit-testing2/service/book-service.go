package service

import (
	"unit-testing2/config"
	"unit-testing2/model"
)

type IBookService interface {
	CreateBookController(book *model.Book) error
	UpdateBookController(book *model.Book, id int) error
	DeleteBookController(book *model.Book, id int) error
	GetBookController(book *model.Book, id int) error
	GetBooksController(books *[]model.Book) error
}

type BookRepository struct {
	Func IBookService
}

var bookRepository IBookService

func init() {
	ur := &BookRepository{}
	ur.Func = ur

	bookRepository = ur
}

func GetBookRepository() IBookService {
	return bookRepository
}

func SetBookRepository(ur IBookService) {
	bookRepository = ur
}

func (b *BookRepository) CreateBookController(book *model.Book) error {
	if err := config.DB.Save(&book).Error; err != nil {
		return err
	}

	return nil
}

func (b *BookRepository) UpdateBookController(book *model.Book, id int) error {
	if err := config.DB.Where("id = ?", id).Updates(&book).Error; err != nil {
		return err
	}

	return nil
}

func (b *BookRepository) DeleteBookController(book *model.Book, id int) error {
	if err := config.DB.Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}

	return nil
}

func (b *BookRepository) GetBookController(book *model.Book, id int) error {
	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return err
	}

	return nil
}

func (b *BookRepository) GetBooksController(books *[]model.Book) error {
	if err := config.DB.Find(&books).Error; err != nil {
		return err
	}

	return nil
}
