package restbooksserver

import "go-grpc/internal/pkg/model"

func DBBook(book *model.Book) *model.DBBook {
	return &model.DBBook{
		Isbn:      book.Isbn,
		Name:      book.Name,
		Publisher: book.Publisher,
	}
}

func Book(dbBook *model.DBBook) *model.Book {
	return &model.Book{
		Isbn:      dbBook.Isbn,
		Name:      dbBook.Name,
		Publisher: dbBook.Publisher,
	}
}
