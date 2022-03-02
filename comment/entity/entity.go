package entity

type CreateCommentInput struct {
	Text   string `binding:"required"`
	UserId uint
}

type CreateCommentUri struct {
	Type     string `uri:"type" binding:"required"`
	SeriesId int    `uri:"seriesId" binding:"required"`
}

type GetCommentByIdUri struct {
	Id int `uri:"id" binding:"required"`
}

type CommentOutput struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type Movie struct {
	OriginalTitle string `json:"original_title"`
	ReleaseDate   string `json:"release_date"`
	PosterPath    string `json:"poster_path"`
}

type Tv struct {
	Name         string `json:"name"`
	FirstAirDate string `json:"first_air_date"`
	PosterPath   string `json:"poster_path"`
}
