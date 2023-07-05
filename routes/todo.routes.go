package routes

import (
	"example/todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(r *gin.Engine)  {

	todoRoutes := r.Group("/todos")

	todoRoutes.GET("/", controllers.GetTodos)
	todoRoutes.GET("/:id", controllers.GetATodo)
	todoRoutes.POST("/", controllers.CreateATodo)
	todoRoutes.PUT("/:id", controllers.UpdateATodo)
	todoRoutes.DELETE("/:id", controllers.DeleteATodo)
}