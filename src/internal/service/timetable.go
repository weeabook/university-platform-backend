package service

import (
	"github.com/weeabook/src/university-platform-backend/internal/model"
	"github.com/weeabook/src/university-platform-backend/internal/repository"
)

type TimetableServiceImpl struct {
	repo repository.TimetableRepository
}

func NewTimetableServiceImpl(repo repository.TimetableRepository) *TimetableServiceImpl {
	return &TimetableServiceImpl{repo: repo}
}

func (t TimetableServiceImpl) GetAll() ([]model.TimeTable, error) {
	return t.repo.GetAll()
}

func (t TimetableServiceImpl) GetByGroup(group int) (string, error) {
	return t.repo.GetByGroup(group)
}
