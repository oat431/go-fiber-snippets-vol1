# Go fiber health check setup with actuator

we can use fiber health check middleware to create health check endpoints for our application. This is useful for monitoring the health of our application and ensuring that it is running properly. The health check middleware provides three endpoints: livez, readyz, and startupz.

```go
package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
)

func StartServer() {

	app := fiber.New()
	app.Get(healthcheck.LivenessEndpoint, healthcheck.New()) // livez
	app.Get(healthcheck.ReadinessEndpoint, healthcheck.New()) // readyz
	app.Get(healthcheck.StartupEndpoint, healthcheck.New()) // startupz

	port := ":" + os.Getenv("PORT")
	err := app.Listen(port)
	if err != nil {
		return
	}
}

```