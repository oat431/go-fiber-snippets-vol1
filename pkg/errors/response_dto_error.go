package errors

type ResponseDTOError struct {
	httpCode  int
	errorCode string
	message   string
}
