package entity

type CreateCommentInput struct {
	Text     string `binding:"required"`
	SeriesId uint
	IsMovie  bool
	UserId   uint
}

type CreateCommentUri struct {
	Type     string `uri:"type" binding:"required"`
	SeriesId int    `uri:"seriesId" binding:"required"`
}
