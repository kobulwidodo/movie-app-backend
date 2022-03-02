package domain

import (
	_movieEntity "movie-app/comment/entity"

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
}

type CommentUsecase interface {
	Create(input _movieEntity.CreateCommentInput, inputUril _movieEntity.CreateCommentUri) (Comment, error)
	GetCommentByUserId(userId uint) ([]Comment, error)
}
