package service

import (
	"play/internal/model"
	"play/internal/repository"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAllBooks() ([]model.Book, error) {
	return s.repo.FindAll()
}

func (s *BookService) GetBook(id int) (*model.Book, error) {
	return s.repo.Find(id)
}

func (s *BookService) CreateBook(book *model.Book) error {
	return s.repo.Create(book)
}

func (s *BookService) UpdateBook(book *model.Book) error {
	return s.repo.Update(book)
}

func (s *BookService) DeleteBook(id int) error {
	return s.repo.Delete(id)
}
