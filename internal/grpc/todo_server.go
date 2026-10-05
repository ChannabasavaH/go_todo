package grpc

import (
	"context"

	"todo/internal/protos"
	"todo/internal/services"
)

type TodoServer struct {
	protos.UnimplementedTodoServer
	service *services.TodoService
}

func NewTodoServer(service *services.TodoService) *TodoServer {
	return &TodoServer{
		service: service,
	}
}

func (s *TodoServer) CreateTodo(ctx context.Context, req *protos.TodoItem) (*protos.TodoItem, error) {
	todo, err := s.service.CreateTodo(
		req.GetTitle(),
		req.GetDescription(),
	)

	if err != nil {
		return nil, err
	}

	return &protos.TodoItem{
		Title: todo.Title,
		Description: todo.Description,
	}, nil
}
