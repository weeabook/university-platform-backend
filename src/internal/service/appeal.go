package service

import (
	"github.com/weeabook/src/university-platform-backend/internal/model"
	"github.com/weeabook/src/university-platform-backend/internal/repository"
)

type Appeal struct {
	repo repository.AppealRepository
}

func NewAppeal(repo repository.AppealRepository) *Appeal {
	return &Appeal{repo: repo}
}

func (s *Appeal) Create(userId int, appeal model.InputAppeal) error {
	return s.repo.Create(userId, appeal)
}
