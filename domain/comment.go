package domain

import (
	_movieEntity "movie-app/comment/entity"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text     string `gorm:"type:text"`
	SeriesId int
	IsMovie  bool
	UserId   uint
	User     User
}

type CommentRepository interface {
	Create(comment Comment) (Comment, error)
}

type CommentUsecase interface {
	Create(input _movieEntity.CreateCommentInput) (Comment, error)
}
