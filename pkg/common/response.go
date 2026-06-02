package common

// Response status constants.
type Status string

const (
	StatusSuccess Status = "SUCCESS"
	StatusFail    Status = "FAIL"
	StatusError   Status = "ERROR"
)

// APIError carries structured error details.
type APIError struct {
	HTTPCode  int    `json:"httpCode"`
	ErrorCode string `json:"errorCode,omitempty"`
	Message   string `json:"message"`
	Details   any    `json:"details,omitempty"` // e.g., validation errors per field
}

// Response is the standard JSON envelope for all API responses.
type Response struct {
	Data   any       `json:"data"`
	Status Status    `json:"status"`
	Error  *APIError `json:"error,omitempty"`
}

// ---- Constructors ----

// OK returns a success response with the given data.
func OK(data any) Response {
	return Response{Data: data, Status: StatusSuccess}
}

// Fail returns a controlled failure (e.g., validation error).
func Fail(err APIError) Response {
	return Response{Status: StatusFail, Error: &err}
}

// Err returns a server-error response.
func Err(err APIError) Response {
	return Response{Status: StatusError, Error: &err}
}
