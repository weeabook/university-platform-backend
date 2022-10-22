package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/weeabook/src/university-platform-backend/internal/model"
)

type AppealRepository interface {
	Create(userId int, appeal model.InputAppeal) error
}

type AuthRepository interface {
	CreateUser(user model.User) (int, error)
	GetUser(user model.InputUser) (model.User, error)
	GetUserInfo(userId int) (model.User, error)
}

type CategoryRepository interface {
	GetCategories() ([]model.Category, error)
}

type NewsRepository interface {
	GetNews() ([]model.News, error)
	GetNewsByCategoryID(categoryId int) ([]model.News, error)
	GetNewsByID(newsId int) (model.News, error)
}

type RoleRepository interface {
	GetRoles() ([]model.Role, error)
}

type TimetableRepository interface {
	// TODO
}

type Repository struct {
	AppealRepository
	AuthRepository
	CategoryRepository
	NewsRepository
	RoleRepository
	TimetableRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AppealRepository:    NewAppealRepo(db),
		AuthRepository:      NewAuthRepo(db),
		CategoryRepository:  NewCategoryRepo(db),
		NewsRepository:      NewNewsRepo(db),
		RoleRepository:      NewRoleRepo(db),
		TimetableRepository: nil,
	}
}
