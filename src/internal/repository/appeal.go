package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/weeabook/src/university-platform-backend/internal/model"
)

type AppealRepo struct {
	db *sqlx.DB
}

func NewAppealRepo(db *sqlx.DB) *AppealRepo {
	return &AppealRepo{db: db}
}

func (r *AppealRepo) Create(userId int, appeal model.InputAppeal) error {
	_, err := r.db.Exec("INSERT INTO appeals (user_id, title, content) VALUES($1, $2, $3)",
		userId, appeal.Title, appeal.Content)

	return err
}
