package model

type User struct {
	Id         int    `json:"-" db:"id"`
	Surname    string `json:"surname" db:"surname"`
	Name       string `json:"username" db:"username"`
	Patronymic string `json:"patronymic" db:"patronymic"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"password" db:"password_hash"`
	Recordbook int    `json:"recordbook" db:"recordbook"`
	GroupId    int    `json:"group_id" db:"group_id""`
	GroupName  string `json:"group_name" db:"group_name""`
	RoleId     string `json:"role_id" db:"-"`
}

type InputUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
