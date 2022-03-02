package domain

import "github.com/golang-jwt/jwt"

type AuthUsecase interface {
	GenerateToken(userId uint) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
