# Setup WebSocket Go Fiber

## 1. Install the required package

```bash
go get github.com/gofiber/contrib/v3/websocket
```

## 2. Create WebSocket handler

```go
func WebSocketExample(c *websocket.Conn) {
	log.Info("WebSocket connection established ", c.Locals("Host"))
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Info("read:", err)
			break
		}
		log.Info("recv: %s", msg)
		reply := "from computer: " + string(msg)
		err = c.WriteMessage(mt, []byte(reply))
		if err != nil {
			log.Info("write:", err)
			break
		}
	}
}
```

## 3. Register WebSocket route

```go
func StartServer() {
    app := fiber.New()

    app.Get("/chat", websocket.New(controller.WebSocketExample))

    port := ":" + os.Getenv("PORT")
    err := app.Listen(port)
    if err != nil {
        return
    }
}
```

## 4. Middleware for WebSocket (Optional)

```go
func WebSocketMiddleware(c fiber.Ctx) error {
	if c.Get("Host") == "localhost:8000" {
		c.Locals("Host", "localhost:8000")
		return c.Next()
	}
	return c.Status(fiber.StatusForbidden).SendString("Request origin not allowed")
}
```

register middleware

```go
app.Use("/chat", WebSocketMiddleware)
```
