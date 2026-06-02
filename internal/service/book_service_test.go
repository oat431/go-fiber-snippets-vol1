package service_test

import (
	"testing"

	"go-fiber-snippets/internal/domain"
	"go-fiber-snippets/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mocks ---

type mockBookRepo struct {
	mock.Mock
}

func (m *mockBookRepo) FindAll() ([]domain.Book, error) {
	args := m.Called()
	return args.Get(0).([]domain.Book), args.Error(1)
}

func (m *mockBookRepo) FindByID(id string) (*domain.Book, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Book), args.Error(1)
}

// --- Tests ---

func TestGetAllBooks_ReturnsBooks(t *testing.T) {
	repo := new(mockBookRepo)
	svc := service.NewBookService(repo)

	expected := []domain.Book{
		{ID: "1", Title: "Book 1"},
		{ID: "2", Title: "Book 2"},
	}
	repo.On("FindAll").Return(expected, nil)

	books, err := svc.GetAllBooks()

	assert.NoError(t, err)
	assert.Len(t, books, 2)
	assert.Equal(t, "Book 1", books[0].Title)
	repo.AssertExpectations(t)
}

func TestGetBookByID_ValidID_ReturnsBook(t *testing.T) {
	repo := new(mockBookRepo)
	svc := service.NewBookService(repo)

	expected := &domain.Book{ID: "1", Title: "Book 1"}
	repo.On("FindByID", "1").Return(expected, nil)

	book, err := svc.GetBookByID("1")

	assert.NoError(t, err)
	assert.Equal(t, "Book 1", book.Title)
	repo.AssertExpectations(t)
}

func TestGetBookByID_EmptyID_ReturnsError(t *testing.T) {
	repo := new(mockBookRepo)
	svc := service.NewBookService(repo)

	// Empty ID should fail before hitting the repository
	book, err := svc.GetBookByID("")

	assert.Error(t, err)
	assert.EqualError(t, err, "book ID cannot be empty")
	assert.Nil(t, book)
	// Repo.FindByID must NOT be called
	repo.AssertNotCalled(t, "FindByID")
}

func TestGetBookByID_NotFound_ReturnsError(t *testing.T) {
	repo := new(mockBookRepo)
	svc := service.NewBookService(repo)

	repo.On("FindByID", "999").Return(nil, assert.AnError)

	book, err := svc.GetBookByID("999")

	assert.Error(t, err)
	assert.Nil(t, book)
	repo.AssertExpectations(t)
}
