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

func (bs *BookService) AddBook(book *model.DBBook) error {
	err := bs.booksRepo.AddBook(book)
	if err != nil {
		return errors.Wrap(err, "failed to add book")
	}
	return nil
}

func (bs *BookService) GetBook(isbn int) (*model.DBBook, error) {
	book, err := bs.booksRepo.GetBook(isbn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get book with isbn %d", isbn)
	}
	if book == nil {
		return nil, errors.New(fmt.Sprintf("book with isbn %d was not found", isbn))
	}
	return book, nil
}

func (bs *BookService) GetAllBooks() ([]*model.DBBook, error) {
	books, err := bs.booksRepo.GetAllBooks()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all books")
	}
	if len(books) == 0 {
		return nil, errors.New("no books present")
	}
	return books, nil
}

func (bs *BookService) RemoveBook(isbn int) error {
	err := bs.booksRepo.RemoveBook(isbn)
	if err != nil {
		return errors.Wrapf(err, "failed to remove book with isbn %d", isbn)
	}
	return nil
}

func (bs *BookService) UpdateBook(book *model.DBBook) error {
	err := bs.booksRepo.UpdateBook(book)
	if err != nil {
		return errors.Wrapf(err, "failed to update book with isbn %d", book.Isbn)
	}
	return nil
}
