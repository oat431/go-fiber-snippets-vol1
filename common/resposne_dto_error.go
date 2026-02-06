package common

type ResponseDTOError struct {
	httpCode  int
	errorCode string
	message   string
}
