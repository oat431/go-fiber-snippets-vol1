# Request Validation

Uses `go-playground/validator/v10` — the de-facto Go validation library.

## Install

```bash
go get github.com/go-playground/validator/v10
```

## Usage

1. Add `validate` struct tags to your request DTO:

```go
type CreateBookRequest struct {
    Title string `json:"title" validate:"required,min=3,max=200"`
}
```

2. Call `validator.Err(req)` in your handler:

```go
import "go-fiber-snippets/pkg/validator"

func (h *Handler) Create(c fiber.Ctx) error {
    var req CreateBookRequest
    if err := c.Bind().Body(&req); err != nil {
        return c.Status(400).JSON(common.Fail(…))
    }

    if errs := validator.Err(req); errs != nil {
        return c.Status(422).JSON(common.Fail(common.APIError{
            Message:  "validation failed",
            HTTPCode: 422,
            Details:  errs,  // per-field errors
        }))
    }

    // … proceed with valid req …
}
```

3. Response shape:

```json
{
  "status": "FAIL",
  "error": {
    "httpCode": 422,
    "message": "validation failed",
    "details": [
      {"field": "title", "message": "this field is required"}
    ]
  }
}
```

## Supported Tags

| Tag | Example | Message |
|---|---|---|
| `required` | `validate:"required"` | "this field is required" |
| `min` | `validate:"min=3"` | "must be at least 3 characters" |
| `max` | `validate:"max=100"` | "must be at most 100 characters" |
| `email` | `validate:"email"` | "must be a valid email address" |
| `url` | `validate:"url"` | "must be a valid URL" |
| `oneof` | `validate:"oneof=admin user"` | "must be one of [admin user]" |

Add more tag-to-message mappings in `pkg/validator/validator.go` → `msgForTag()`.
