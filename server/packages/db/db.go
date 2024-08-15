package db

import "gorm.io/gorm"

type GormDB struct {
	DB *gorm.DB
}

func NewGormDB(db *gorm.DB) *GormDB {
	return &GormDB{DB: db}
}

type Database interface {
	GetAllTasks() ([]Task, error)
}
