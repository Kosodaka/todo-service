package models

type TodoItem struct {
	Id          int    `bun:"id"`
	Title       string `bun:"title"`
	Description string `bun:"description"`
	Date        string `bun:"date"`
	Done        bool   `bun:"done"`
}
