package handlers

import (
	"log"
	"net/http"
	"strconv"
	"todo/internal/models"
	"todo/internal/repository"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Repo *repository.TodoRepository
}

func NewTodoHandler(repo *repository.TodoRepository) *TodoHandler {
	return &TodoHandler{
		Repo: repo,
	}
}

func (h *TodoHandler) CreateTodo(c *gin.Context){
	var todo models.Todos

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invaild body request",
		})
		return
	}

	if todo.Title == "" || todo.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title and Description is required",
		})
		return
	}

	if err := h.Repo.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Todo created successfully",
		"todo": todo,
	})
}

func (h *TodoHandler) ReadTodo(c *gin.Context){
	var todos []models.Todos

	if err := h.Repo.ReadTodo(&todos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Todos Retrived successfully",
		"todos": todos,
	})
}

func (h *TodoHandler) ReadTodoById(c *gin.Context){
	idParam := c.Param("id")
	log.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID Format",
		})
		return
	}

	var todo models.Todos
	if err := h.Repo.ReadTodoById(uint(id), &todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Todo Found",
		"todo": todo,
	})
}

func (h *TodoHandler) UpdateTodoById(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID Format",
		})
		return
	} 

	var todo models.Todos
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update" + err.Error(),
		})
		return
	}

	if todo.Title == "" || todo.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title and description are required",
		})
		return
	}

	if err := h.Repo.UpdateTodoById(uint(id), &todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Update Successfully",
		"todo": todo,
	})
}

func (h* TodoHandler) DeleteTodoById(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Id Format",
		})
		return
	}

	var todo models.Todos
	if err := h.Repo.DeleteTodoById(uint(id), &todo); err != nil{

		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Todo Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Deleted Successfully",
		"todo": todo,
	})
}