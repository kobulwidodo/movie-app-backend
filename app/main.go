package main

import (
	"fmt"
	"log"
	_authUsecase "movie-app/auth/usecase"
	_commentHttpDelivery "movie-app/comment/delivery/http"
	_commentRepository "movie-app/comment/repository"
	_commentUsecase "movie-app/comment/usecase"
	"movie-app/domain"
	"movie-app/middlewares"
	_userHttpDelivery "movie-app/user/delivery/http"
	_userRepository "movie-app/user/repository"
	_userUsecase "movie-app/user/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	db, err := initDb()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	authUsecase := _authUsecase.NewAuthUsecase()

	userRepository := _userRepository.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepository)
	jwtMiddleware := middlewares.NewAuthMiddleware(authUsecase, userUsecase)
	_userHttpDelivery.NewUserHandler(r, userUsecase, authUsecase)

	commentRepository := _commentRepository.NewCommentRepository(db)
	commentUsecase := _commentUsecase.NewCommentRepository(commentRepository)
	_commentHttpDelivery.NewCommentHandler(r, commentUsecase, jwtMiddleware)

	r.Run()
}

func initDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := DB.AutoMigrate(&domain.User{}, &domain.Comment{}); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return DB, nil
}
