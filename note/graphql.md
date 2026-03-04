# Create GraphQL API for Go fiber

1. Install necessary packages:
```bash
    go get github.com/graphql-go/graphql
```

2. Because of Graphql is single endpoint, so we need to create a handler for GraphQL and register it to fiber app:

graphql handler:
```go
import (
	"github.com/gofiber/fiber/v3"
	"github.com/graphql-go/graphql"
)

type graphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
	// OperationName string `json:"operationName"` // (Optional)
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
```

create boostrap configuration for GraphQL:
```go
type AppContainer struct {
	GraphQLController *controller.GraphQLController
}

func NewAppContainer() *AppContainer {

	//repo1 := repository.NewRepo1()
	//repo2 := repository.NewRepo2()
	//service1 := service.NewService1(repo1)
    //service2 := service.NewService2(repo2)

	schema, err := graphql_api.SetupSchema(/*service...*/)
	if err != nil {
		log.Fatal("Failed to setup GraphQL schema: ", err)
	}

	gqlController := controller.NewGraphQLController(schema)

	return &AppContainer{
		GraphQLController: gqlController,
	}
}
```

register to fiber app:
```go
func main() {
    app := fiber.New()

    container := NewAppContainer()

    // GraphQL Endpoint
    app.Post("/graphql", container.GraphQLController.ExecuteQuery)

    log.Fatal(app.Listen(":3000"))
}
```
if you already did this step, you can skip to next step

---

now let's create graphql api step by step

1. Create Model (That Linked to Database):
```go
    type Book struct {
        ID    string `json:"id"`
        Title string `json:"title"`
    }
```

2. Create GraphQL Model that related to the model
```go
import "github.com/graphql-go/graphql"

var BookType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Book",
        Fields: graphql.Fields{
            "id":    &graphql.Field{Type: graphql.String},
            "title": &graphql.Field{Type: graphql.String},
        },
    })
```

3. Create Business Layer and Data Access Layer (Service and Repository) like normal REST API (This is not related to GraphQL, but it is a good practice to separate concerns)

Repository:
```go
type BookRepository interface {
	FindAll() ([]domain.Book, error)
	FindByID(id string) (*domain.Book, error)
}

type mockBookRepository struct {
	mockData []domain.Book
}

// normally, you will implement this repository with database connection, but for simplicity, we will use mock data here
func NewMockBookRepository() BookRepository {
	return &mockBookRepository{
		mockData: []domain.Book{
			{ID: "1", Title: "Go Programming Fundamentals"},
			{ID: "2", Title: "Clean Architecture in Go"},
			{ID: "3", Title: "Mastering GraphQL"},
		},
	}
}

func (r *mockBookRepository) FindAll() ([]domain.Book, error) {
	return r.mockData, nil
}

func (r *mockBookRepository) FindByID(id string) (*domain.Book, error) {
	for _, book := range r.mockData {
		if book.ID == id {
			return &book, nil
		}
	}

	return nil, errors.New("book not found in mock database")
}
```

Service:
```go
type BookService interface {
	GetAllBooks() ([]domain.Book, error)
	GetBookByID(id string) (*domain.Book, error)
}

type bookService struct {
	repo repository.BookRepository
}

func (b bookService) GetAllBooks() ([]domain.Book, error) {
	return b.repo.FindAll()
}

func (b bookService) GetBookByID(id string) (*domain.Book, error) {
	return b.repo.FindByID(id)
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}
```

4. Create Query or Mutaion (it is considered as a controller in REST API)
```go
func GetBooksQuery(bookService service.BookService) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.BookType),
		Description: "Get all books from database",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return bookService.GetAllBooks()
		},
	}
}

func GetBookByIDQuery(bookService service.BookService) *graphql.Field {
	return &graphql.Field{
		Type:        types.BookType,
		Description: "Get a single book by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			return bookService.GetBookByID(id)
		},
	}
}
```

5. Registered Query or Mutation to Schema
```go
func SetupSchema(bookService service.BookService) (graphql.Schema, error) {

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			// "nameOfQuery":  queries.NameOfQuery(service),
			"GetAllBook":  queries.GetBooksQuery(bookService),
			"GetBookById": queries.GetBookByIDQuery(bookService),

			// if you have more than one service, you can add more query here:
			// "users": queries.GetUsersQuery(userService),
			// "links": queries.GetLinksQuery(linkService),
		},
	})

	/* // Mutations (Insert/Update)
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createBook": mutations.CreateBookMutation(bookService),
		},
	})
	*/

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
		// Mutation: rootMutation,
	})
}
```

6. then back to register all the essential component to AppContainer and run the app, you can test the GraphQL API 
```graphql
query AllBookFunction {
    GetAllBook {
        id
        title
    }
}
```

response 
```json
{
  "data": {
    "GetAllBook": [
      {
        "id": "1",
        "title": "Go Programming Fundamentals"
      },
      {
        "id": "2",
        "title": "Clean Architecture in Go"
      },
      {
        "id": "3",
        "title": "Mastering GraphQL"
      }
    ]
  }
}
```
