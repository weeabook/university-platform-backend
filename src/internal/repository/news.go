package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/weeabook/src/university-platform-backend/internal/model"
)

type NewsRepo struct {
	db *sqlx.DB
}

func NewNewsRepo(db *sqlx.DB) *NewsRepo {
	return &NewsRepo{db: db}
}

func (r *NewsRepo) GetNews() ([]model.News, error) {
	var news []model.News

	err := r.db.Get(&news, `SELECT ti.id, ti.img_url, ti.title, ti.content, ti.created_at, ul.name FROM news ti 
							INNER JOIN news_category li ON li.news_id=ti.id 
							INNER JOIN categories ul ON ul.id=li.category_id`)

	return news, err
}

func (r *NewsRepo) GetNewsByCategoryID(categoryId int) ([]model.News, error) {
	var news []model.News

	err := r.db.Get(&news, `SELECT ti.id, ti.img_url, ti.title, ti.content, ti.created_at, ul.name FROM news ti
							INNER JOIN news_category li ON li.news_id=ti.id
							INNER JOIN categories ul ON ul.id=li.category_id WHERE li.category_id=$1`, categoryId)

	return news, err
}

func (r *NewsRepo) GetNewsByID(newsId int) (model.News, error) {
	var news model.News

	err := r.db.Get(&news, `SELECT ti.id, ti.img_url, ti.title, ti.content, ti.created_at, ul.name FROM news ti
							INNER JOIN news_category li ON li.news_id=ti.id
							INNER JOIN categories ul ON ul.id=li.category_id WHERE ti.id=$1`, newsId)

	return news, err
}
