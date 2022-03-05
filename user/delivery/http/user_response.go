package http

import "movie-app/domain"

type userAuthResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Bio   string `json:"bio"`
	Token string `json:"token"`
}

func UserAuthResponse(user domain.User, token string) userAuthResponse {
	res := userAuthResponse{
		Name:  user.Name,
		Email: user.Email,
		Bio:   user.Bio,
		Token: token,
	}

	return res
}
