package entity

import "time"

type Task struct {
	TaskId    string    `json:"task_id" gorm:"primaryKey;column:task_id;autoIncrement:false;"`
	UserId    string    `json:"user_id" gorm:"column:user_id;"`
	Title     string    `json:"title" gorm:"column:title;"`
	Content   string    `json:"content" gorm:"column:content;"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}
