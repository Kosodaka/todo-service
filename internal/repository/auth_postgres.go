package repository

import (
	"github.com/Kosodaka/todo-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type AuthPostgres struct {
	db *bun.DB
}

func NewAuthPostgres(db *bun.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(c *gin.Context, user models.User) (int, error) {
	var id int
	err := r.db.NewInsert().Model(&user).Returning("id").Scan(c, &id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (r *AuthPostgres) GetUser(c *gin.Context, username, password string) (models.User, error) {
	var user models.User
	err := r.db.NewSelect().Model(&user).Column("id").Where("username=?", username).Where("password_hash=?", password).Scan(c, &user)

	return user, err
}
