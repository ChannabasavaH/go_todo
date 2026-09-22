package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"todo/internal/config"
	"todo/internal/database"
	"todo/internal/handlers"
	"todo/internal/repository"
	"todo/internal/routes"
)

func main(){

	cfg := config.LoadConfig()

	db := database.ConnectPostgres(cfg)

	todoRepo := repository.NewTodoRepository(db)

	todoHandler := handlers.NewTodoHandler(todoRepo)

	router := gin.Default()

	routes.SetUpRoutes(router, todoHandler)

	fmt.Println("Todo API Server Running ")
	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Println("Server failed to start:", err)
	}
	_=db
}
