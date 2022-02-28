package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrEmailConflict       = errors.New("the email has been already registered")
	ErrInputBinding        = errors.New("wrong parameter")
)
