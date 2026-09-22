package routes

import (
	"todo/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes (router *gin.Engine, todoHandler *handlers.TodoHandler){
	api := router.Group("/api")
	{
		api.POST("/todos", todoHandler.CreateTodo)
		api.GET("/todos", todoHandler.ReadTodo)
		api.GET("/todos/:id", todoHandler.ReadTodoById)
		api.PUT("/todos/:id", todoHandler.UpdateTodoById)
		api.DELETE("/todos/:id", todoHandler.DeleteTodoById)
	}
}