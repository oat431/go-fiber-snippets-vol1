package domain

type ResponseDTOStatus string

const (
	SUCCESS ResponseDTOStatus = "SUCCESS"
	FAIL    ResponseDTOStatus = "FAIL"
	ERROR   ResponseDTOStatus = "ERROR"
)
