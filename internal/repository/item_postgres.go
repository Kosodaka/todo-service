package repository

import (
	"database/sql"
	"fmt"
	"github.com/Kosodaka/todo-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"strings"
	"time"
)

type TodoItemPostgres struct {
	db *bun.DB
}

func NewTodoItemPostgres(db *bun.DB) *TodoItemPostgres {
	return &TodoItemPostgres{
		db: db,
	}
}

func (r *TodoItemPostgres) Create(c *gin.Context, userId int, item models.TodoItem) (int, error) {

	tx, err := r.db.BeginTx(c, &sql.TxOptions{})
	if err != nil {
		return 0, err
	}
	var id int
	timeParse(&item.CreatedAt)
	err = r.db.NewInsert().Model(&item).Returning("id").Scan(c, &id)
	if err != nil {
		err = tx.Rollback()
		return 0, err
	}
	createUsersItemsQuery := fmt.Sprintf("INSERT INTO %s (user_id, item_id) VALUES (?, ?)", usersItemsTable)
	_, err = r.db.Exec(createUsersItemsQuery, &userId, &id)
	if err != nil {
		err = tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(c *gin.Context, userId int) ([]models.TodoItem, error) {
	var items []models.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description,ti.created_at, ti.done FROM %s ti INNER JOIN %s ui on ti.id=ui.item_id WHERE ui.user_id = ?", todoItemTable, usersItemsTable)
	_, err := r.db.NewRaw(query, &userId).Exec(c, &items)
	return items, err
}
func (r *TodoItemPostgres) GetById(c *gin.Context, userId, itemId int) (models.TodoItem, error) {
	var item models.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description,ti.created_at, ti.done FROM %s ti 
                        INNER JOIN %s ui on ti.id=ui.item_id WHERE ui.user_id = ? AND ui.item_id=?`, todoItemTable, usersItemsTable)
	_, err := r.db.NewRaw(query, &userId, &itemId).Exec(c, &item)
	return item, err
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s ui WHERE ti.id = ui.item_id AND ui.user_id=? AND ui.item_id=?",
		todoItemTable, usersItemsTable)
	_, err := r.db.Exec(query, &userId, &itemId)

	return err
}

func (r *TodoItemPostgres) Update(c *gin.Context, userId, itemId int, input models.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=?"))
		args = append(args, *input.Title)
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=?"))
		args = append(args, *input.Description)
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s ui WHERE ti.id = ui.item_id AND ui.item_id=? AND ui.user_id=?",
		todoItemTable, setQuery, usersItemsTable)
	args = append(args, itemId, userId)

	_, err := r.db.NewRaw(query, args...).Exec(c)
	return err
}

func timeParse(currentDate *time.Time) time.Time {
	*currentDate = time.Now()
	location, err := time.LoadLocation("Asia/Yekaterinburg")
	if err != nil {
		log.Fatal()
	}

	return currentDate.In(location)
}
