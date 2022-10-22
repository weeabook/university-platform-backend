package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/weeabook/src/university-platform-backend/internal/model"
)

type CategoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) GetCategories() ([]model.Category, error) {
	var categories []model.Category

	err := r.db.Get(&categories, "SELECT * FROM categories")
	return categories, err
}
