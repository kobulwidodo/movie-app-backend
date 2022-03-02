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

func (r *commentRepository) GetByUserId(userId uint) ([]domain.Comment, error) {
	var comments []domain.Comment
	if err := r.db.Preload("User").Where("user_id = ?", userId).Find(&comments).Error; err != nil {
		return comments, err
	}

	return comments, nil
}

func (r *commentRepository) GetBySeriesId(seriesId string) ([]domain.Comment, error) {
	var comments []domain.Comment
	if err := r.db.Preload("User").Where("series_id = ?", seriesId).Find(&comments).Error; err != nil {
		return comments, err
	}

	return comments, nil
}
