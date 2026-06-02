package repository

import (
	"errors"
	"go-fiber-snippets/internal/domain"
)

// BookRepository defines the data-access contract.
// Swap this mock for a real DB implementation without touching the service layer.
type BookRepository interface {
	FindAll() ([]domain.Book, error)
	FindByID(id string) (*domain.Book, error)
}

// --- Mock implementation (no database required) ---

type mockBookRepo struct {
	data []domain.Book
}

// NewMockBookRepo returns an in-memory repository seeded with sample data.
func NewMockBookRepo() BookRepository {
	return &mockBookRepo{
		data: []domain.Book{
			{ID: "1", Title: "Go Programming Fundamentals"},
			{ID: "2", Title: "Clean Architecture in Go"},
			{ID: "3", Title: "Mastering GraphQL"},
		},
	}
}

func (r *mockBookRepo) FindAll() ([]domain.Book, error) {
	return r.data, nil
}

func (r *mockBookRepo) FindByID(id string) (*domain.Book, error) {
	for i := range r.data {
		if r.data[i].ID == id {
			return &r.data[i], nil
		}
	}
	return nil, errors.New("book not found")
}
