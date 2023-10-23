package service

import (
	"github.com/Kosodaka/todo-service/internal/models"
	"github.com/Kosodaka/todo-service/internal/repository"
	"github.com/gin-gonic/gin"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(c *gin.Context, userId int, item models.TodoItem) (int, error) {
	return s.repo.Create(c, userId, item)
}

func (s *TodoItemService) GetById(c *gin.Context, userId, itemId int) (models.TodoItem, error) {
	return s.repo.GetById(c, userId, itemId)
}
func (s *TodoItemService) GetAll(c *gin.Context, userId int) ([]models.TodoItem, error) {
	return s.repo.GetAll(c, userId)
}
func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}
func (s *TodoItemService) Update(c *gin.Context, userId, itemId int, input models.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(c, userId, itemId, input)
}
