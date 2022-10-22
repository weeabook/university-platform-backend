package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/weeabook/src/university-platform-backend/internal/model"
)

type RoleRepo struct {
	db *sqlx.DB
}

func NewRoleRepo(db *sqlx.DB) *RoleRepo {
	return &RoleRepo{db: db}
}

func (r *RoleRepo) GetRoles() ([]model.Role, error) {
	var roles []model.Role
	err := r.db.Get(&roles, "SELECT * FROM roles")

	return roles, err
}
