package main

import (
	"elogistics-backend/database"
	"elogistics-backend/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectMongoDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
