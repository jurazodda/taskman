package models

import "time"

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	Done        *bool     `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type CreateTaskParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MarkTaskParams struct {
	IsDone bool `json:"is_done"`
}

type UpdateTaskParams struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
