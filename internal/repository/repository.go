package repository

import (
	"github.com/Kosodaka/todo-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type Authorization interface {
	CreateUser(c *gin.Context, user models.User) (int, error)
	GetUser(c *gin.Context, username, password string) (models.User, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoItem
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
