package domain

import "go-fiber-snippets/pkg/errors"

type ResponseDTO[T any] struct {
	Data   T                        `json:"data"`
	Status ResponseDTOStatus        `json:"status"`
	Error  *errors.ResponseDTOError `json:"error"`
}
