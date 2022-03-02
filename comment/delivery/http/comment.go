package http

import (
	"movie-app/comment/entity"
	"movie-app/domain"
	"movie-app/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	CommentUsecase domain.CommentUsecase
}

func NewCommentHandler(r *gin.Engine, cu domain.CommentUsecase, jwtMiddleware gin.HandlerFunc) {
	handler := &CommentHandler{CommentUsecase: cu}
	r.POST("/comment/:type/:seriesId", jwtMiddleware, handler.Create)
	r.GET("/comment/user/:id", handler.GetCommentByUserId)
	r.GET("/comment/:id", handler.GetCommentBySeriesId)
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

	userLoggedin := c.MustGet("userLoggedin").(domain.User)
	input.UserId = userLoggedin.ID

	var newComment domain.Comment
	newComment, err := h.CommentUsecase.Create(input, inputUri)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), &utils.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &utils.Response{Data: &entity.CommentOutput{Id: int(newComment.ID), Text: newComment.Text, CreatedAt: newComment.CreatedAt.String()}, Message: "comment has ben created"})
}

func (h *CommentHandler) GetCommentByUserId(c *gin.Context) {
	var inputUri entity.GetCommentByIdUri
	if err := c.ShouldBindUri(&inputUri); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{Message: err.Error()})
		return
	}

	comments, err := h.CommentUsecase.GetCommentByUserId(uint(inputUri.Id))
	if err != nil {
		c.JSON(utils.GetErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, &utils.Response{Data: CommentsResponse(comments), Message: "successfully get comments"})
}

func (h *CommentHandler) GetCommentBySeriesId(c *gin.Context) {
	var inputUri entity.GetCommentByIdUri
	if err := c.ShouldBindUri(&inputUri); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{Message: err.Error()})
		return
	}

	comments, err := h.CommentUsecase.GetCommentBySeriesId(strconv.Itoa(inputUri.Id))
	if err != nil {
		c.JSON(utils.GetErrorCode(err), &utils.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &utils.Response{Data: CommentsSeriesResponse(comments), Message: "successfully get comments"})
}
