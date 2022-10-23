package service

import (
	"github.com/weeabook/src/university-platform-backend/internal/model"
	"github.com/weeabook/src/university-platform-backend/internal/repository"
)

type AppealService interface {
	Create(userId int, appeal model.InputAppeal) error
}

type AuthService interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
	GetUserInfo(userId int) (model.User, error)
}

type CategoryService interface {
	GetCategories() ([]model.Category, error)
}

type NewsService interface {
	GetNews() ([]model.News, error)
	GetNewsByCategoryID(categoryId int) ([]model.News, error)
	GetNewsByID(newsId int) (model.News, error)
}

type RoleService interface {
	GetRoles() ([]model.Role, error)
}

type TimetableService interface {
	GetAll() ([]string, error)
	GetByGroup(group int) (string, error)
}

type Service struct {
	AppealService
	AuthService
	CategoryService
	NewsService
	RoleService
	TimetableService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AppealService:    NewAppeal(repos.AppealRepository),
		AuthService:      NewAuth(repos.AuthRepository),
		CategoryService:  NewCategory(repos.CategoryRepository),
		NewsService:      NewNews(repos.NewsRepository),
		RoleService:      NewRole(repos.RoleRepository),
		TimetableService: nil,
	}
}
