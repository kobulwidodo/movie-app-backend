package http

import "movie-app/domain"

type commentResponse struct {
	Id        int           `json:"id"`
	Text      string        `json:"text"`
	SeriesId  string        `json:"series_id"`
	Type      string        `json:"type"`
	CreatedAt string        `json:"created_at"`
	User      user          `json:"user"`
	Series    domain.Series `json:"series"`
}

type commentSeriesResponse struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	User      user   `json:"user"`
}

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func CommentsSeriesResponse(comments []domain.Comment) []commentSeriesResponse {
	commentsRes := []commentSeriesResponse{}
	for _, comment := range comments {
		commentsRes = append(commentsRes, CommentSeriesResponse(comment))
	}

	return commentsRes
}

func CommentSeriesResponse(comment domain.Comment) commentSeriesResponse {
	user := user{
		Id:   int(comment.User.ID),
		Name: comment.User.Name,
	}
	commentRes := commentSeriesResponse{
		Id:        int(comment.ID),
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt.String(),
		User:      user,
	}

	return commentRes
}

func CommentsResponse(comments []domain.Comment) []commentResponse {
	commentsRes := []commentResponse{}
	for _, comment := range comments {
		commentsRes = append(commentsRes, CommentResponse(comment))
	}
	return commentsRes
}

func CommentResponse(comment domain.Comment) commentResponse {
	user := user{
		Id:   int(comment.User.ID),
		Name: comment.User.Name,
	}
	var _type string
	if comment.IsMovie {
		_type = "movie"
	} else {
		_type = "tv"
	}
	commentRes := commentResponse{
		Id:        int(comment.ID),
		Text:      comment.Text,
		SeriesId:  comment.SeriesId,
		Type:      _type,
		CreatedAt: comment.CreatedAt.String(),
		User:      user,
		Series:    comment.Series,
	}
	return commentRes
}
