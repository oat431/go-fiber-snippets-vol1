# Setup Go Fiber Framework

## Prerequisites
- Go installed on your machine (version 1.16 or higher recommended)
- Basic understanding of Go programming language
- Go modules enabled
- A code editor (like VSCode)

## Step 1: Create a New Go Project
1. Open your terminal.
2. Create a new directory for your project and navigate into it:
3. ```bash
    mkdir my-fiber-app
    cd my-fiber-app
    ```
4. Initialize a new Go module:
5. ```bash
    go mod init my-fiber-app
    ```
## Step 2: Install Fiber Framework
1. Run the following command to install the Fiber package:
2. ```bash
    go get github.com/gofiber/fiber/v2
    ```
## Step 3: Create a Simple Fiber Application
1. Create a new file named `main.go` in your project directory.
2. Open `main.go` in your code editor and add the following code:
3. ```go
    package main

    import (
        "github.com/gofiber/fiber/v3"
    )

    func main() {
        // Create a new Fiber instance
        app := fiber.New()

        // Define a simple route
        app.Get("/", func(c *fiber.Ctx) error {
            return c.SendString("Hello, World!")
        })

        // Start the server on port 3000
        app.Listen(":3000")
    }
    ```
## Step 4: Run Your Fiber Application
1. In your terminal, run the following command to start your Fiber application:
2. ```bash
    go run main.go
    ```
3. Open your web browser and navigate to `http://localhost:3000`. You should
4. see "Hello, World!" displayed on the page.
5. Congratulations! You have successfully set up a basic Go Fiber application.


