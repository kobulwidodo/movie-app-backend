package domain

type AuthUsecase interface {
	GenerateToken(userId uint) (string, error)
}
