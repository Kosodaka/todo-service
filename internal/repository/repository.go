package repository

import "github.com/uptrace/bun"

type Authorization interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoItem
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{}
}
