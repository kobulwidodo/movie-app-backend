package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrEmailConflict       = errors.New("the email has been already registered")
	ErrInputBinding        = errors.New("wrong parameter")
	ErrNotFound            = errors.New("credential not found")
	ErrPassNotMatch        = errors.New("your login credentials dont match in our system")
)
