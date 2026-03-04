package service

import (
	"go-fiber-snippets/internal/delivery/http/repository"
	"go-fiber-snippets/internal/domain"
)

type BookService interface {
	GetAllBooks() ([]domain.Book, error)
	GetBookByID(id string) (*domain.Book, error)
}

type bookService struct {
	repo repository.BookRepository
}

func (b bookService) GetAllBooks() ([]domain.Book, error) {
	return b.repo.FindAll()
}

func (b bookService) GetBookByID(id string) (*domain.Book, error) {
	return b.repo.FindByID(id)
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}
