package database

import "gorm.io/gorm"

type Goal struct {
	DB *gorm.DB
}

func NewGoal(db *gorm.DB) *Goal {
	return &Goal{DB: db}
}