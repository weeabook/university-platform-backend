package model

type Role struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
