package service

import "github.com/Kosodaka/todo-service/internal/repository"

type Authorization interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
