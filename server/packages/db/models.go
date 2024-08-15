package db

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title           string
	Description     string
	LastCompletedAt *time.Time
	Interval        time.Duration
}
