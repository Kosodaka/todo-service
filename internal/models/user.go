package models

type User struct {
	Id       int    `json:"-" bun:"id,pk,autoincrement" `
	Name     string `json:"name" binding:"required" bun:"name"`
	Username string `json:"username" binding:"required" bun:"username"`
	Password string `json:"password" binding:"required" bun:"password_hash"`
}
