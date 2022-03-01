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
	AuthUsecase domain.AuthUsecase
}

func NewUserHandler(r *gin.Engine, uu domain.UserUsecase, au domain.AuthUsecase) {
	handler := &UserHandler{UserUsecase: uu, AuthUsecase: au}
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
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

func (h *UserHandler) Login(c *gin.Context) {
	var input entity.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, &util.Response{Message: err.Error()})
		return
	}

	var userLogged domain.User
	userLogged, err := h.UserUsecase.Login(input)
	if err != nil {
		c.JSON(util.GetErrorCode(err), &util.Response{Message: err.Error()})
		return
	}

	token, err := h.AuthUsecase.GenerateToken(userLogged.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &util.Response{Message: domain.ErrInternalServerError.Error()})
		return
	}

	c.JSON(http.StatusOK, &util.Response{Data: &entity.GetUser{Name: userLogged.Name, Email: userLogged.Email, Bio: userLogged.Bio, Token: token}, Message: "Success Login"})
}
