package util

import (
	"movie-app/domain"
	"net/http"
)

func GetErrorCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrEmailConflict:
		return http.StatusConflict
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrInputBinding:
		return http.StatusUnprocessableEntity
	case domain.ErrPassNotMatch:
		return http.StatusUnauthorized
	case domain.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
