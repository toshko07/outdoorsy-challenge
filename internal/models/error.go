package models

import (
	"fmt"
)

const (
	InternalError = "InternalError"
	NotFoundError = "NotFoundError"
)

type ServiceError struct {
	Msg  string
	Code string
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("'%s': '%s'", e.Code, e.Msg)
}

func NewInternalError(message string) ServiceError {
	return ServiceError{
		Msg:  message,
		Code: InternalError,
	}
}

func NewNotFoundError(message string) ServiceError {
	return ServiceError{
		Msg:  message,
		Code: NotFoundError,
	}
}
