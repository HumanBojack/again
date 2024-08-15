package db

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskInput
	LastCompletedAt *time.Time `json:"last_completed_at"`
}

type TaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Interval    int    `json:"interval"`
}
