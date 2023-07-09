package main

import (
	"example/todo-app/database"
	"example/todo-app/middlewares"
	"example/todo-app/models"
	"example/todo-app/routes"
	"io"
	"log"
	"net/http"
	"os"

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


func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	loadEnv()
	loadDatabase()
	setupLogOutput()

	r := gin.New();
	r.Static("/css", "./templates/css")
	r.LoadHTMLGlob("templates/*.html")

	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	r.GET("/", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	
	routes.SetupTodoRoutes(r)

	r.Run()	
}
