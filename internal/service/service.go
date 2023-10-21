package service

import (
	"github.com/Kosodaka/todo-service/internal/models"
	"github.com/Kosodaka/todo-service/internal/repository"
	"github.com/gin-gonic/gin"
)

type Authorization interface {
	CreateUser(c *gin.Context, user models.User) (int, error)
	CreateToken(c *gin.Context, username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
