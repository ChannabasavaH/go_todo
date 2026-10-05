package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"


	"todo/internal/config"
	"todo/internal/database"
	"todo/internal/handlers"
	"todo/internal/repository"
	"todo/internal/routes"
	"todo/internal/services"
	"todo/internal/protos"
	grcpHandler "todo/internal/grpc"
)

func main(){

	cfg := config.LoadConfig()

	db := database.ConnectPostgres(cfg)

	todoRepo := repository.NewTodoRepository(db)

	todoService := services.NewTodoService(todoRepo)

	todoHandler := handlers.NewTodoHandler(todoService)

	todoServer := grcpHandler.NewTodoServer(todoService)

	grpcServer := grpc.NewServer()

	protos.RegisterTodoServer(
		grpcServer,
		todoServer,
	)

	go func() {
		lis, err := net.Listen("tcp", ":50051")

		if err != nil {
			log.Fatal(err)
		}

		log.Println("gRPC server running on: 50051")

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()

	routes.SetUpRoutes(router, todoHandler)

	fmt.Println("Todo API Server Running ")
	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Println("Server failed to start:", err)
	}
	_=db
}
