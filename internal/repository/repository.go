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
	Create(c *gin.Context, userId int, item models.TodoItem) (int, error)
	GetAll(c *gin.Context, userId int) ([]models.TodoItem, error)
	GetById(c *gin.Context, userId, itemId int) (models.TodoItem, error)
	Delete(userId, itemId int) error
	Update(c *gin.Context, userId, itemId int, input models.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoItem
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
