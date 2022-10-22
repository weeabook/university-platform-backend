package model

type Appeal struct {
	Id      int    `json:"id" db:"id"`
	UserId  int    `json:"-" db:"user_id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}

type InputAppeal struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
