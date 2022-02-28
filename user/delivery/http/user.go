package http

import (
	"movie-app/domain"
	"movie-app/user/entity"
	"movie-app/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(r *gin.Engine, uu domain.UserUsecase) {
	handler := &UserHandler{UserUsecase: uu}
	r.POST("/register", handler.Register)
}

func (h *UserHandler) Register(c *gin.Context) {
	var input entity.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, &util.Response{Message: err.Error()})
		return
	}

	newUser, err := h.UserUsecase.Register(input)
	if err != nil {
		c.JSON(util.GetErrorCode(err), &util.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &util.Response{Data: &entity.GetUser{Name: newUser.Name, Email: newUser.Email, Bio: newUser.Bio}, Message: "successfully created a user"})
}
