# Pagination

Reusable query-param extraction + response envelope. No external deps.

## Usage

### In the handler

```go
import "go-fiber-snippets/pkg/pagination"

func (h *Handler) List(c fiber.Ctx) error {
    params := pagination.FromQuery(c)   // ?page=1&limit=20

    books, total, err := h.svc.ListBooks(params.Offset(), params.Limit)
    if err != nil { … }

    return c.JSON(common.OK(pagination.PaginatedResponse[domain.Book]{
        Data: books,
        Meta: params.BuildMeta(total),
    }))
}
```

### Response shape

```json
{
  "status": "SUCCESS",
  "data": {
    "data": [{ "id": "1", "title": "…" }],
    "meta": {
      "page": 1,
      "limit": 20,
      "total": 157,
      "totalPages": 8
    }
  }
}
```

## API

| Function | Description |
|---|---|
| `pagination.FromQuery(c)` | Extracts `page` & `limit` from query string. Defaults: page=1, limit=20. Caps limit at 100. |
| `params.Offset()` | Returns `(page-1)*limit` — ready to use with SQL `OFFSET`. |
| `params.BuildMeta(total)` | Returns `Meta{Page, Limit, Total, TotalPages}`. |
| `pagination.PaginatedResponse[T]` | Generic struct wrapping `data []T` + `meta`. |
