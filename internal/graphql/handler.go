package graphql

import (
	"github.com/gofiber/fiber/v3"
	"github.com/graphql-go/graphql"
)

// Handler returns a Fiber handler that executes GraphQL queries.
// POST to /graphql with JSON body: {"query": "...", "variables": {...}}
func Handler(schema graphql.Schema) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req struct {
			Query     string                 `json:"query"`
			Variables map[string]interface{} `json:"variables"`
		}

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid JSON body",
			})
		}

		if req.Query == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "query is required",
			})
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  req.Query,
			VariableValues: req.Variables,
			Context:        c.Context(),
		})

		return c.JSON(result)
	}
}
