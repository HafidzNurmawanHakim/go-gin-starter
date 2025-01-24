package main

import (
	"gin-template/internal/models"
	"gin-template/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var SECRET_KEY string

func main() {
	r := gin.Default()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	SECRET_KEY = os.Getenv("SECRET_KEY")

	config := models.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}

	models.InitDB(config)

	routes.AuthRoutes(r)
	routes.UserRoute(r)
	routes.Role(r)

	r.Run(":8000")
}
