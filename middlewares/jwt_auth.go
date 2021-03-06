package middlewares

import (
	"movie-app/domain"
	"movie-app/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	authUsecase domain.AuthUsecase
	userUsecase domain.UserUsecase
}

func NewAuthMiddleware(au domain.AuthUsecase, uu domain.UserUsecase) gin.HandlerFunc {
	return (&AuthMiddleware{authUsecase: au, userUsecase: uu}).Handle()
}

func (m *AuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &utils.Response{Message: "wrong token type"})
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := m.authUsecase.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &utils.Response{Message: "token invalid"})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &utils.Response{Message: "unauthorized"})
			return
		}

		userId := uint(claim["user_id"].(float64))
		var user domain.User
		user, err = m.userUsecase.GetUserById(userId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &utils.Response{Message: "failed to get user data"})
			return
		}

		c.Set("userLoggedin", user)
	}
}
