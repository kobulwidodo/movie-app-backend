package usecase

import (
	"errors"
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

func (s *authUsecase) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Token Invalid")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}
