package repository

import (
	"play/internal/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]model.Book, error)
	Find(id int) (*model.Book, error)
	Create(book *model.Book) error
	Update(book *model.Book) error
	Delete(id int) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) Find(id int) (*model.Book, error) {
	var book model.Book
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r *bookRepository) Create(book *model.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) Update(book *model.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id int) error {
	return r.db.Delete(&model.Book{}, id).Error
}
