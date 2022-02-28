package domain

import (
	"movie-app/user/entity"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Bio      string
}

type UserUsecase interface {
	Register(input entity.CreateUserInput) (User, error)
}

type UserRepository interface {
	Create(user User) (User, error)
	GetByEmail(email string) (User, error)
}
