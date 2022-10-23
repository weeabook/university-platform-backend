package model

type TimeTable struct {
	Id       int    `json:"id"`
	FileName string `json:"file_name"`
	GroupId  int    `json:"group_id"`
}
