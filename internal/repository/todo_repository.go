package repository

import (
	"errors"
	"todo/internal/models"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		DB: db,
	}
}

func (r *TodoRepository) CreateTodo(todo *models.Todos) error {
	return r.DB.Create(todo).Error
}

func (r *TodoRepository) ReadTodo(todos *[]models.Todos) error {
	return r.DB.Find(todos).Error
}

func (r *TodoRepository) ReadTodoById(id uint, todo *models.Todos) error {
	return r.DB.First(todo, id).Error
}

func (r *TodoRepository) UpdateTodoById(id uint, todos *models.Todos) error {
	return r.DB.Model(&models.Todos{}).Where("id=?", id).Updates(todos).Error
}

func (r *TodoRepository) DeleteTodoById(id uint,todos *models.Todos) error {
	result := r.DB.Delete(todos, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0{
		return errors.New("record not found")
	}

	return nil
}