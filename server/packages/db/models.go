package db

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskInput
}

type TaskInput struct {
	Title           string     `json:"title,omitempty"`
	Description     string     `json:"description,omitempty"`
	Interval        int        `json:"interval,omitempty"`
	LastCompletedAt *time.Time `json:"last_completed_at,omitempty"`
}
