package repository

import (
	"errors"
	"go-fiber-snippets/internal/domain"
)

type BookRepository interface {
	FindAll() ([]domain.Book, error)
	FindByID(id string) (*domain.Book, error)
}

type mockBookRepository struct {
	mockData []domain.Book
}

func NewMockBookRepository() BookRepository {
	return &mockBookRepository{
		// ใส่ข้อมูลจำลองตั้งต้น (Seed Data)
		mockData: []domain.Book{
			{ID: "1", Title: "Go Programming Fundamentals"},
			{ID: "2", Title: "Clean Architecture in Go"},
			{ID: "3", Title: "Mastering GraphQL"},
		},
	}
}

func (r *mockBookRepository) FindAll() ([]domain.Book, error) {
	return r.mockData, nil
}

func (r *mockBookRepository) FindByID(id string) (*domain.Book, error) {
	for _, book := range r.mockData {
		if book.ID == id {
			return &book, nil
		}
	}

	return nil, errors.New("book not found in mock database")
}
