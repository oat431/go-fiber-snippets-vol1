package handler

import (
	"log/slog"

	"go-fiber-snippets/internal/domain"
	"go-fiber-snippets/internal/service"
	"go-fiber-snippets/pkg/common"
	"go-fiber-snippets/pkg/pagination"
	"go-fiber-snippets/pkg/validator"

	"github.com/gofiber/fiber/v3"
)

// BookHandler groups all book-related HTTP handlers.
type BookHandler struct {
	svc service.BookService
}

func NewBookHandler(svc service.BookService) *BookHandler {
	return &BookHandler{svc: svc}
}

// CreateBookRequest maps the JSON body for POST /books.
type CreateBookRequest struct {
	Title string `json:"title" validate:"required,min=3,max=200"`
}

// ListBooks returns a paginated list of books.
//
//	GET /api/v1/books?page=1&limit=20
func (h *BookHandler) ListBooks(c fiber.Ctx) error {
	params := pagination.FromQuery(c)

	books, err := h.svc.GetAllBooks()
	if err != nil {
		slog.Error("failed to fetch books", "error", err)
		return c.Status(500).JSON(common.Err(
			common.APIError{Message: "failed to fetch books", HTTPCode: 500},
		))
	}

	total := len(books)
	offset := params.Offset()
	if offset > total {
		offset = total
	}
	end := offset + params.Limit
	if end > total {
		end = total
	}

	resp := pagination.PaginatedResponse[domain.Book]{
		Data: books[offset:end],
		Meta: params.BuildMeta(total),
	}

	slog.Debug("books listed", "page", params.Page, "limit", params.Limit, "total", total)
	return c.JSON(common.OK(resp))
}

// CreateBook validates and creates a new book.
//
//	POST /api/v1/books  {"title": "Learning Go"}
func (h *BookHandler) CreateBook(c fiber.Ctx) error {
	var req CreateBookRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(common.Fail(
			common.APIError{Message: "invalid JSON body", HTTPCode: 400},
		))
	}

	if errs := validator.Err(req); errs != nil {
		return c.Status(422).JSON(common.Fail(
			common.APIError{
				Message:  "validation failed",
				HTTPCode: 422,
				Details:  errs,
			},
		))
	}

	// In production: h.svc.CreateBook(req.Title) → persist and return
	book := domain.Book{ID: "new", Title: req.Title}
	slog.Info("book created", "title", req.Title)

	return c.Status(201).JSON(common.OK(book))
}
