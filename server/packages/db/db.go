package db

import "gorm.io/gorm"

type GormDB struct {
	DB *gorm.DB
}

func NewGormDB(db *gorm.DB) *GormDB {
	return &GormDB{DB: db}
}

type Database interface {
	CreateTask(task *Task) error
	GetTask(id string) (*Task, error)
	UpdateTask(id string, task *Task) error
	DeleteTask(id string) error
	GetAllTasks() ([]Task, error)
}
