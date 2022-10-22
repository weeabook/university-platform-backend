package model

import "time"

type News struct {
	Id        int       `json:"id" db:"id"`
	ImgUrl    string    `json:"imgUrl" db:"img_url"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Category  string    `json:"category" db:"name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
