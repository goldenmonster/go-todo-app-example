package main

import (
	"example/todo-app/database"
	"example/todo-app/models"
	"example/todo-app/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {

	database.Connect()
	database.Database.AutoMigrate(&models.Todo{})
}

func main() {
	loadEnv()
	loadDatabase()
	r := gin.Default();
	r.GET("/", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	
	routes.SetupTodoRoutes(r)

	r.Run()	
}
