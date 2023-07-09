package controllers

import (
	"example/todo-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo

	err := models.GetAllTodos(&todos)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodoListPage (c *gin.Context) {
	var todos []models.Todo
	err := models.GetAllTodos((&todos))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	
	data := gin.H{
		"title": "Todo List page",
		"todos": todos,
	}

	c.HTML(http.StatusOK, "index.html", data)
}


func CreateATodo (c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)  // required

	err := models.CreateATodo(&todo)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo

	err := models.GetATodo(&todo, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func UpdateATodo (c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.GetATodo(&todo, id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)

	err = models.UpdateATodo(&todo, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func DeleteATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
}