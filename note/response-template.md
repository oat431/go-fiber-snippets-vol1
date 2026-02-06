# Setup Response Temp# API Response Template

My Personal API Response Template, inspired by [Jsend](https://github.com/omniti-labs/jsend) format.

```json
{
  "status": "SUCCESS | FAIL | ERROR",
  "data": {
    // Your response data goes here
  },
  "error" : {
    "httpCode": "HTTP_STATUS_CODE",
    "errorCode": "PROGRAM_ERROR_CODE",
    "message": "Detailed error message"
  }
}
```

---

Go struct Representation:

reponse_dto_status.go
```go
package common

type ResponseDTOStatus string

const (
SUCCESS ResponseDTOStatus = "SUCCESS"
FAIL    ResponseDTOStatus = "FAIL"
ERROR   ResponseDTOStatus = "ERROR"
)

```

response_dto_error.go
```go
package common

type ResponseDTOError struct {
	httpCode  int
	errorCode string
	message   string
}
```

response_dto.go
```go
package common

type ResponseDTO[T any] struct {
	Data   T                 `json:"data"`
	Status ResponseDTOStatus `json:"status"`
	Error  *ResponseDTOError `json:"error"`
}

```