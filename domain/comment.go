package domain

import (
	_commentEntity "movie-app/comment/entity"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text     string `gorm:"type:text"`
	SeriesId string
	IsMovie  bool
	UserId   uint
	User     User
	Series   Series `gorm:"-"`
}

type Series struct {
	Title       string `json:"title"`
	Year        string `json:"year"`
	ImagePoster string `json:"image_poster"`
}

type CommentRepository interface {
	Create(comment Comment) (Comment, error)
	GetByUserId(userId uint) ([]Comment, error)
	GetBySeriesId(seriesId string) ([]Comment, error)
	GetById(id uint) (Comment, error)
	Delete(comment Comment) error
}

type CommentUsecase interface {
	Create(input _commentEntity.CreateCommentInput, inputUri _commentEntity.CreateCommentUri) (Comment, error)
	GetCommentByUserId(userId uint) ([]Comment, error)
	GetCommentBySeriesId(seriesId string) ([]Comment, error)
	DeleteComment(input _commentEntity.GetCommentByIdUri, userId uint) error
}
