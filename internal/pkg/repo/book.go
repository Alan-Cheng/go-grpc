package repo

import (
	"go-grpc/internal/pkg/model"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func GetNewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) AddBook(book *model.DBBook) error {
	result := br.db.Create(book)
	return result.Error
}

func (br *BookRepository) UpdateBook(book *model.DBBook) error {
	result := br.db.Model(&model.DBBook{}).Where("isbn = ?", book.Isbn).Updates(map[string]interface{}{
		"name":      book.Name,
		"publisher": book.Publisher,
	})
	return result.Error
}

func (br *BookRepository) GetBook(isbn int) (*model.DBBook, error) {
	var book model.DBBook
	result := br.db.First(&book, isbn)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &book, nil
}

func (br *BookRepository) GetAllBooks() ([]*model.DBBook, error) {
	books := make([]*model.DBBook, 0)
	err := br.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (br *BookRepository) RemoveBook(isbn int) error {
	result := br.db.Delete(&model.DBBook{}, isbn)
	return result.Error
}
