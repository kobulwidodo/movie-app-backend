package repository

import (
	"movie-app/domain"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment domain.Comment) (domain.Comment, error) {
	if err := r.db.Create(&comment).Error; err != nil {
		return comment, err
	}

	return comment, nil
}
