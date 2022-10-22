package service

import (
	"github.com/weeabook/src/university-platform-backend/internal/model"
	"github.com/weeabook/src/university-platform-backend/internal/repository"
)

type Category struct {
	repo repository.CategoryRepository
}

func NewCategory(repo repository.CategoryRepository) *Category {
	return &Category{repo: repo}
}

func (s *Category) GetCategories() ([]model.Category, error) {
	return s.repo.GetCategories()
}
