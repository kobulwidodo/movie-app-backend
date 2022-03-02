package http

import "movie-app/domain"

type commentResponse struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	User      user   `json:"user"`
	Series    series `json:"series"`
}

type user struct {
	Name string `json:"name"`
}

type series struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Year        string `json:"year"`
	ImagePoster string `json:"image_poster"`
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
	series := series{
		Id:          comment.SeriesId,
		Title:       comment.Series.Title,
		Year:        comment.Series.Year,
		ImagePoster: comment.Series.ImagePoster,
	}
	commentRes := commentResponse{
		Id:        int(comment.ID),
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt.String(),
		User:      user,
		Series:    series,
	}
	return commentRes
}
