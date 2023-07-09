package routes

import (
	"example/todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(r *gin.Engine)  {

	todoApiRoutes := r.Group("/api/todos")

	todoApiRoutes.GET("/", controllers.GetTodos)
	todoApiRoutes.GET("/:id", controllers.GetATodo)
	todoApiRoutes.POST("/", controllers.CreateATodo)
	todoApiRoutes.PUT("/:id", controllers.UpdateATodo)
	todoApiRoutes.DELETE("/:id", controllers.DeleteATodo)

	todoViewRoutes := r.Group("todos")
	todoViewRoutes.GET("/", controllers.GetTodoListPage)
}