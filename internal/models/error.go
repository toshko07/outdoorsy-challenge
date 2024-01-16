package models

import (
	"fmt"
)

const (
	InternalErrorCode = "InternalError"
	NotFoundErrorCode = "NotFoundError"
)

type ServiceError struct {
	Msg  string
	Code string
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("'%s': '%s'", e.Code, e.Msg)
}

func NewServiceError(message, code string) ServiceError {
	return ServiceError{
		Msg:  message,
		Code: code,
	}
}

type InternalError ServiceError

func (e InternalError) Error() string {
	return e.Msg
}

func NewInternalError(msg string) InternalError {
	return InternalError(NewServiceError(msg, InternalErrorCode))
}

type NotFoundError ServiceError

func (e NotFoundError) Error() string {
	return e.Msg
}

func NewNotFoundError(msg string) NotFoundError {
	return NotFoundError(NewServiceError(msg, NotFoundErrorCode))
}
