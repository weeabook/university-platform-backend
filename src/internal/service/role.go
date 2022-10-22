package service

import (
	"github.com/weeabook/src/university-platform-backend/internal/model"
	"github.com/weeabook/src/university-platform-backend/internal/repository"
)

type Role struct {
	repo repository.RoleRepository
}

func NewRole(repo repository.RoleRepository) *Role {
	return &Role{repo: repo}
}

func (s *Role) GetRoles() ([]model.Role, error) {
	return s.repo.GetRoles()
}
