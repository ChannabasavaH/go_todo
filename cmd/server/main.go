package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodo(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}

func createTodo(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"method": "POST"})
}

func updateTodo(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"method": "PUT"})
}

func deleteTodo(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}

func main(){
	router := gin.Default()
	
	router.GET("/todos", getTodo)
	router.POST("/todos", createTodo)
	router.PUT("/todos/:id", updateTodo)
	router.DELETE("todos/:id", deleteTodo)
		
	fmt.Println("Todo API Server Running ")
	err := router.Run(); if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
