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

type user struct {
	Name string `json:"name"`
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
