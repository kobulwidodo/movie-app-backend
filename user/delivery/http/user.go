package http

import (
	"movie-app/domain"
	"movie-app/user/entity"
	"movie-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
	AuthUsecase domain.AuthUsecase
}

func NewUserHandler(r *gin.Engine, uu domain.UserUsecase, au domain.AuthUsecase) {
	handler := &UserHandler{UserUsecase: uu, AuthUsecase: au}
	api := r.Group("/auth")
	{
		api.POST("/register", handler.Register)
		api.POST("/login", handler.Login)
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var input entity.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, &utils.Response{Message: err.Error()})
		return
	}

	newUser, err := h.UserUsecase.Register(input)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), &utils.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &utils.Response{Data: UserAuthResponse(newUser, ""), Message: "successfully created a user"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var input entity.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, &utils.Response{Message: err.Error()})
		return
	}

	var userLogged domain.User
	userLogged, err := h.UserUsecase.Login(input)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), &utils.Response{Message: err.Error()})
		return
	}

	token, err := h.AuthUsecase.GenerateToken(userLogged.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{Message: domain.ErrInternalServerError.Error()})
		return
	}

	c.JSON(http.StatusOK, &utils.Response{Data: UserAuthResponse(userLogged, token), Message: "Success Login"})
}
