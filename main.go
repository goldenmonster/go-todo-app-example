package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ToDo struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Hours int `json:"hours"`
	Done bool `json:"done"`
}


var todos = []ToDo {
	{ID: "1", Title: "Todo 1", Hours: 4, Done: false},
	{ID: "2", Title: "Todo 2", Hours: 3, Done: true},
	{ID: "3", Title: "Todo 3", Hours: 2, Done: true},
}

func getTodoList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var todo ToDo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todos = append(todos, todo)

	c.IndentedJSON(http.StatusCreated, todo)
}

func getTodoById(c *gin.Context) {
	id := c.Param("id")

	for _, todo := range todos {
		if todo.ID == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	for i, a := range todos {
		if a.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	c.Status(http.StatusNoContent)
}

func main() {
	r := gin.Default();
	r.GET("/", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	todoRoutes := r.Group("/todos")
	todoRoutes.GET("/", getTodoList)
	todoRoutes.POST("/", createTodo)
	todoRoutes.GET("/:id", getTodoById)
	todoRoutes.DELETE("/:id", deleteTodo)
	r.Run()	
}
