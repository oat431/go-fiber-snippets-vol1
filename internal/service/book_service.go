package service

import (
	"errors"
	"go-fiber-snippets/internal/domain"
	"go-fiber-snippets/internal/repository"
)

// BookService defines the business-logic contract.
type BookService interface {
	GetAllBooks() ([]domain.Book, error)
	GetBookByID(id string) (*domain.Book, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() ([]domain.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetBookByID(id string) (*domain.Book, error) {
	if id == "" {
		return nil, errors.New("book ID cannot be empty")
	}
	return s.repo.FindByID(id)
}
