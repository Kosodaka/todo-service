package models

import (
	"errors"
	"time"
)

type TodoItem struct {
	Id          int       `json:"id"  bun:"id,pk,autoincrement"`
	Title       string    `json:"title" bun:"title"`
	Description string    `json:"description" bun:"description"`
	CreatedAt   time.Time `json:"created_at" bin:"created_at"`
	Done        bool      `json:"done" bin:"done, default:false"`
}

type UsersItems struct {
	Id     int `json:"id" bun:"id,pk,autoincrement"`
	ItemId int `json:"item_id" bun:"item_id"`
	UserId int `json:"user_id" bun:"user_id"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
