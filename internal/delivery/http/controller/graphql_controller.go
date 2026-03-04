package controller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/graphql-go/graphql"
)

type graphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
	// OperationName string `json:"operationName"` // (Optional) ใส่ไว้เผื่ออนาคตได้ครับ
}

type GraphQLController struct {
	schema graphql.Schema
}

func NewGraphQLController(schema graphql.Schema) *GraphQLController {
	return &GraphQLController{
		schema: schema,
	}
}

func (ctrl *GraphQLController) ExecuteQuery(c fiber.Ctx) error {
	req := new(graphQLRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid JSON payload",
			"details": err.Error(),
		})
	}

	if req.Query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "GraphQL query is required",
		})
	}

	params := graphql.Params{
		Schema:         ctrl.schema,
		RequestString:  req.Query,
		VariableValues: req.Variables,
		Context:        c.Context(),
	}

	result := graphql.Do(params)

	return c.Status(fiber.StatusOK).JSON(result)
}
