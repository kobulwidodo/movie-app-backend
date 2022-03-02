package usecase

import (
	"encoding/json"
	"io/ioutil"
	"movie-app/comment/entity"
	"movie-app/domain"
	"net/http"
	"strconv"
)

type commentUsecase struct {
	commentRepository domain.CommentRepository
}

func NewCommentRepository(cr domain.CommentRepository) *commentUsecase {
	return &commentUsecase{commentRepository: cr}
}

func (s *commentUsecase) Create(input entity.CreateCommentInput, inputUri entity.CreateCommentUri) (domain.Comment, error) {
	comment := domain.Comment{
		Text:     input.Text,
		SeriesId: strconv.Itoa(inputUri.SeriesId),
		UserId:   input.UserId,
	}

	if inputUri.Type == "movie" {
		comment.IsMovie = true
	} else if inputUri.Type == "tv" {
		comment.IsMovie = false
	} else {
		return comment, domain.ErrBadRequest
	}

	newComment, err := s.commentRepository.Create(comment)
	if err != nil {
		return comment, domain.ErrInternalServerError
	}

	return newComment, nil
}

func (s *commentUsecase) GetCommentByUserId(userId uint) ([]domain.Comment, error) {
	var comments []domain.Comment
	comments, err := s.commentRepository.GetByUserId(userId)
	if err != nil {
		return comments, domain.ErrInternalServerError
	}

	if comments[0].ID == 0 {
		return comments, domain.ErrNotFound
	}

	var newComments []domain.Comment

	for _, comment := range comments {
		var url string
		if comment.IsMovie {
			url = "/movie/" + comment.SeriesId + "?api_key=6386ea8d1021b8ae5f21896f87ee2a09"
		} else {
			url = "/tv/" + comment.SeriesId + "?api_key=6386ea8d1021b8ae5f21896f87ee2a09"
		}
		res, err := http.Get("https://api.themoviedb.org/3" + url)
		if err != nil {
			return comments, err
		}

		resData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return comments, err
		}

		if comment.IsMovie {
			var movie entity.Movie
			json.Unmarshal(resData, &movie)
			comment.Series.Title = movie.OriginalTitle
			comment.Series.Year = movie.ReleaseDate
			comment.Series.ImagePoster = movie.PosterPath
		} else {
			var tv entity.Tv
			json.Unmarshal(resData, &tv)
			comment.Series.Title = tv.Name
			comment.Series.Year = tv.FirstAirDate
			comment.Series.ImagePoster = tv.PosterPath
		}
		newComments = append(newComments, comment)
	}

	return newComments, nil
}

func (s *commentUsecase) GetCommentBySeriesId(seriesId string) ([]domain.Comment, error) {
	var comments []domain.Comment
	comments, err := s.commentRepository.GetBySeriesId(seriesId)
	if err != nil {
		return comments, err
	}

	return comments, nil
}
