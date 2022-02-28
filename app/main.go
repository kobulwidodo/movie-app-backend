package main

import (
	"fmt"
	"log"
	"movie-app/domain"
	_userHttpDelivery "movie-app/user/delivery/http"
	"movie-app/user/repository"
	"movie-app/user/usecase"
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

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	_userHttpDelivery.NewUserHandler(r, userUsecase)

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

	if err := DB.AutoMigrate(&domain.User{}); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return DB, nil
}
