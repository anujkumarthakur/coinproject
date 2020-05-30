package main

import (
	//"github.com/gin-gonic/gin"
	"coinproject/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//var router *gin.Engine

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the routes
	r := routes.InitializeRoutes()

	// Start serving the application
	//router.Run()
	r.Run(":" + os.Getenv("PORT"))

}
