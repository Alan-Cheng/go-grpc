package service

import (
	"fmt"
	"go-grpc/internal/pkg/model"
	"go-grpc/internal/pkg/repo"

	"github.com/pkg/errors"
)

type BookService struct {
	booksRepo *repo.BookRepository
}

func GetNewBookService(bookRepo *repo.BookRepository) *BookService {
	return &BookService{booksRepo: bookRepo}
}

func (bs *BookService) AddBook(book *model.DBBook) {
	bs.booksRepo.AddBook(book)
}

func (bs *BookService) GetBook(isbn int) (*model.DBBook, error) {
	book := bs.booksRepo.GetBook(isbn)
	if book != nil {
		return book, nil
	}
	return nil, errors.New(fmt.Sprintf("book with isbn %d was not found", isbn))
}

func (bs *BookService) GetAllBooks() ([]*model.DBBook, error) {
	books, err := bs.booksRepo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, errors.New("No books present")
	}
	return books, nil
}

func (bs *BookService) RemoveBook(isbn int) {
	bs.booksRepo.RemoveBook(isbn)
}

func (bs *BookService) UpdateBook(book *model.DBBook) {
	bs.booksRepo.UpdateBook(book)
}
