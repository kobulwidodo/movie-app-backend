package repository

import (
	"movie-app/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user domain.User) (domain.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetById(userId uint) (domain.User, error) {
	var user domain.User
	if err := r.db.Where("id = ?", userId).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Update(user domain.User) (domain.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
