package usecase

import (
	"movie-app/domain"
	"movie-app/user/entity"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (s *userUsecase) Register(input entity.CreateUserInput) (domain.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}

	existedUser, _ := s.userRepository.GetByEmail(input.Email)
	if existedUser != (domain.User{}) {
		return domain.User{}, domain.ErrEmailConflict
	}

	user := domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(passwordHash),
		Bio:      input.Bio,
	}

	newUser, err := s.userRepository.Create(user)
	if err != nil {
		return newUser, domain.ErrInternalServerError
	}

	return newUser, nil
}

func (s *userUsecase) Login(input entity.LoginInput) (domain.User, error) {
	var user domain.User
	user, err := s.userRepository.GetByEmail(input.Email)
	if err != nil {
		return user, domain.ErrInternalServerError
	}

	if user.ID == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, domain.ErrPassNotMatch
	}

	return user, nil
}

func (s *userUsecase) GetUserById(userId uint) (domain.User, error) {
	var user domain.User
	user, err := s.userRepository.GetById(userId)
	if err != nil {
		return user, domain.ErrNotFound
	}

	return user, nil
}
