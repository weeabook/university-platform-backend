package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/weeabook/src/university-platform-backend/internal/model"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(user model.User) (int, error) {
	var args []any
	var userId int

	fields := "username, surname, email, password_hash"
	values := "$1, $2, $3"
	args = append(args, user.Name, user.Surname, user.Email, user.Password)

	if user.Patronymic != "" {
		fields = fields + ", patronymic"
		values = values + ", $4"
		args = append(args, user.Patronymic)
	}

	row := r.db.QueryRow(fmt.Sprintf("INSERT INTO users (%s) VALUES (%s) RETURNING id", fields, values), args...)
	if err := row.Scan(&userId); err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *AuthRepo) GetUser(user model.InputUser) (model.User, error) {
	var data model.User

	err := r.db.Get(&data, "SELECT * FROM users WHERE email=$1 AND password_hash=$2", user.Email, user.Password)
	return data, err
}

func (r *AuthRepo) GetUserInfo(userId int) (model.User, error) {
	var user model.User

	err := r.db.Get(&user, "SELECT * FROM users WHERE id=$1", userId)
	return user, err
}
