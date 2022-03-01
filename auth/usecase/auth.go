package usecase

import (
	"os"

	"github.com/golang-jwt/jwt"
)

type authUsecase struct {
}

func NewAuthUsecase() *authUsecase {
	return &authUsecase{}
}

func (s *authUsecase) GenerateToken(userId uint) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
