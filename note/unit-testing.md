# GO Unit testing

1. install test dependecies
```bash
go get github.com/stretchr/testify
```

2. Create test file with _test at the end (go philosophy is simply put the test file and real file in the same working folder)
```go
package service_test // you can also put _test at the end too

import (
	"go-fiber-snippets/internal/delivery/http/service"
	"go-fiber-snippets/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

func TestGetAllBooks_Success(t *testing.T) {
	mockRepo := new(mockBookRepo)
	bookService := service.NewBookService(mockRepo)

	mockBooks := []domain.Book{{ID: "1", Title: "Book 1"}, {ID: "2", Title: "Book 2"}}
	mockRepo.On("FindAll").Return(mockBooks, nil)

	result, err := bookService.GetAllBooks()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Book 1", result[0].Title)

	mockRepo.AssertExpectations(t)
}

func TestGetBookByID_EmptyID_ShouldReturnError(t *testing.T) {
	mockRepo := new(mockBookRepo)
	bookService := service.NewBookService(mockRepo)

	result, err := bookService.GetBookByID("")

	assert.Error(t, err)
	assert.Equal(t, "book ID cannot be empty", err.Error())
	assert.Nil(t, result)
}

```

3. run the test file
```bash
go test ./...
```

or
if you have makefile
```bash
make test
```
