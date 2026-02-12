# Setup Redirect Feature for Go Fiber Application

```go
    func ToGitHub(c fiber.Ctx) error {
        log.Info(c.Method() + " " + c.Path() + "\n")
        return c.Redirect().Status(301).To("https://example.com")
    }
```

301 Moved Permanently indicates that the resource requested has been permanently moved to a new URL provided in the Location header.