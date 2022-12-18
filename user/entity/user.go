package entity

type CreateUserInput struct {
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type LoginInput struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type UpdateBioInput struct {
	Bio    string `binding:"required"`
	UserId uint
}
