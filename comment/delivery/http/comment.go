package http

import (
	"movie-app/comment/entity"
	"movie-app/domain"
	"movie-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	CommentUsecase domain.CommentUsecase
}

func NewCommentHandler(r *gin.Engine, cu domain.CommentUsecase) {
	handler := &CommentHandler{CommentUsecase: cu}
	r.POST("/comment/:type/:seriesId", handler.Create)
}

func (h *CommentHandler) Create(c *gin.Context) {
	var inputUri entity.CreateCommentUri
	if err := c.ShouldBindUri(&inputUri); err != nil {
		c.JSON(http.StatusUnprocessableEntity, &utils.Response{Message: err.Error()})
		return
	}

	var input entity.CreateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, &utils.Response{Message: err.Error()})
		return
	}
}
