package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrEmailConflict       = errors.New("the email has been already registered")
	ErrInputBinding        = errors.New("wrong parameter")
	ErrNotFound            = errors.New("not found")
	ErrPassNotMatch        = errors.New("your login credentials dont match in our system")
	ErrBadRequest          = errors.New("bad request")
	ErrForbidden           = errors.New("forbidden request")
)
