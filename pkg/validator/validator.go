package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidationError maps field paths to human-readable messages.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Err returns all validation errors for the given struct.
// Usage:
//
//	type CreateBookReq struct {
//	    Title string `json:"title" validate:"required,min=3,max=100"`
//	}
//
//	if errs := validator.Err(req); errs != nil {
//	    return c.Status(400).JSON(common.Fail(…, errs))
//	}
func Err(s interface{}) []ValidationError {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	var errs []ValidationError
	for _, e := range err.(validator.ValidationErrors) {
		errs = append(errs, ValidationError{
			Field:   jsonField(e),
			Message: msgForTag(e),
		})
	}
	return errs
}

// --- helpers ---

// jsonField returns the json tag name (falls back to struct field name).
func jsonField(e validator.FieldError) string {
	// e.StructField() gives the Go field name; we want the json tag.
	// We parse the tag from e.StructNamespace() or just lowercase the field.
	name := e.Field()
	// Attempt to use the actual namespace (which matches json tag in practice)
	if ns := e.Namespace(); ns != "" {
		parts := strings.Split(ns, ".")
		if len(parts) > 0 {
			name = parts[len(parts)-1]
		}
	}
	// Fall back to lowercase Go field name for readability
	return strings.ToLower(name[:1]) + name[1:]
}

func msgForTag(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "this field is required"
	case "min":
		return fmt.Sprintf("must be at least %s characters", e.Param())
	case "max":
		return fmt.Sprintf("must be at most %s characters", e.Param())
	case "email":
		return "must be a valid email address"
	case "url":
		return "must be a valid URL"
	case "oneof":
		return fmt.Sprintf("must be one of [%s]", e.Param())
	default:
		return fmt.Sprintf("failed validation: %s", e.Tag())
	}
}
