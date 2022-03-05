package entity

type CreateUserInput struct {
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Bio      string `binding:"required"`
}

type LoginInput struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
