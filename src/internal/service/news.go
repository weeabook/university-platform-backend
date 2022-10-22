package service

import (
	"github.com/weeabook/src/university-platform-backend/internal/model"
	"github.com/weeabook/src/university-platform-backend/internal/repository"
)

type News struct {
	repo repository.NewsRepository
}

func NewNews(repo repository.NewsRepository) *News {
	return &News{repo: repo}
}

func (s *News) GetNews() ([]model.News, error) {
	return s.repo.GetNews()
}

func (s *News) GetNewsByCategoryID(categoryId int) ([]model.News, error) {
	return s.repo.GetNewsByCategoryID(categoryId)
}

func (s *News) GetNewsByID(newsId int) (model.News, error) {
	return s.repo.GetNewsByID(newsId)
}
