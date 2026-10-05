package services

import (
	"todo/internal/models"

	"todo/internal/repository"

)

type TodoService struct {
	Repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{
		Repo: repo,
	}
}

func (s *TodoService) CreateTodo(title, description string) (*models.Todos, error){
	todo := &models.Todos{
		Title: title,
		Description: description,
	}

	if err := s.Repo.CreateTodo(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) GetTodos() ([]models.Todos, error){
	var todos []models.Todos

	if err := s.Repo.ReadTodo(&todos); err!= nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodoService) GetTodoById(id uint) (*models.Todos, error) {
	var todo *models.Todos

	if err := s.Repo.ReadTodoById(id, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) UpdateTodoById(id uint, title string, description string, completed bool) (*models.Todos, error){
	todo := &models.Todos{
		ID: id,
		Title: title,
		Description: description,
		Completed: completed,
	}

	if err := s.Repo.UpdateTodoById(id, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) DeleteTodoById(id uint) (*models.Todos, error) {
	var todo *models.Todos
	if err := s.Repo.DeleteTodoById(id, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

