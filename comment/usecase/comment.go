package usecase

import (
	"movie-app/comment/entity"
	"movie-app/domain"
)

type commentUsecase struct {
	commentRepository domain.CommentRepository
}

func NewCommentRepository(cr domain.CommentRepository) *commentUsecase {
	return &commentUsecase{commentRepository: cr}
}

func (s *commentUsecase) Create(input entity.CreateCommentInput) (domain.Comment, error) {
	comment := domain.Comment{
		Text:     input.Text,
		SeriesId: int(input.SeriesId),
		IsMovie:  input.IsMovie,
		UserId:   input.UserId,
	}

	newComment, err := s.commentRepository.Create(comment)
	if err != nil {
		return comment, domain.ErrInternalServerError
	}

	return newComment, nil
}
