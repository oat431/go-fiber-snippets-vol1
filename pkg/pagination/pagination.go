package pagination

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// Params are extracted from query string: ?page=1&limit=20.
// Caps limit at 100 for safety.
type Params struct {
	Page  int
	Limit int
}

// Meta is the pagination envelope returned alongside data.
type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
}

// FromQuery extracts pagination params from the request.
// Defaults: page=1, limit=20.
func FromQuery(c fiber.Ctx) Params {
	p := Params{Page: 1, Limit: 20}

	if v := c.Query("page"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			p.Page = n
		}
	}
	if v := c.Query("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			p.Limit = n
		}
	}

	// Safety cap
	if p.Limit > 100 {
		p.Limit = 100
	}

	return p
}

// BuildMeta computes pagination metadata from total count and params.
func (p Params) BuildMeta(total int) Meta {
	totalPages := int(math.Ceil(float64(total) / float64(p.Limit)))
	if totalPages < 1 {
		totalPages = 1
	}
	return Meta{
		Page:       p.Page,
		Limit:      p.Limit,
		Total:      total,
		TotalPages: totalPages,
	}
}

// Offset returns the DB offset for this page.
func (p Params) Offset() int {
	return (p.Page - 1) * p.Limit
}

// PaginatedResponse wraps data with pagination metadata.
// Embed this in your API response.
type PaginatedResponse[T any] struct {
	Data []T  `json:"data"`
	Meta Meta `json:"meta"`
}
