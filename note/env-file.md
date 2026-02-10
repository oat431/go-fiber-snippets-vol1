# Setup .env file for Go Fiber Framework

## Step 1: install godotenv package
```bash
go get github.com/joho/godotenv
```

## Step 2: Create a .env file
1. In the root directory of your Go Fiber project, create a new file named `.env
2. Open the `.env` file in your code editor and add your environment variables in the following format:
```env
PORT=8080
DATABASE_URL=your_database_url
SECRET_KEY=your_secret_key
```

## Step 3: Load .env file in your Go Fiber application
1. Open your `main.go` file and import the `godotenv` package:
2. Load the `.env` file at the beginning of your `main` function and access the environment variables using `os.Getenv`:
```go
import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    port := os.Getenv("PORT")
    databaseURL := os.Getenv("DATABASE_URL")
    secretKey := os.Getenv("SECRET_KEY")

    // Use the environment variables in your application
    log.Printf("Server will run on port: %s", port)
    log.Printf("Database URL: %s", databaseURL)
    log.Printf("Secret Key: %s", secretKey)

    // Continue with your Fiber application setup...
}
```
