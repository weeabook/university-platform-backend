package repository

import "github.com/jmoiron/sqlx"

type TimeTableRepo struct {
	db *sqlx.DB
}

func (t TimeTableRepo) GetAll() ([]string, error) {
	var result []string
	err := t.db.Select(&result, "SELECT filename from schedulefile")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t TimeTableRepo) GetByGroup(group int) (string, error) {
	var result string
	err := t.db.Get(&result, "SELECT filename FROM schedulefile where group_id = $1", group)
	if err != nil {
		return "", nil
	}
	return result, nil
}

func NewTimeTableRepo(db *sqlx.DB) *TimeTableRepo {
	return &TimeTableRepo{db: db}
}
