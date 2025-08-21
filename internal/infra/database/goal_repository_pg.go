package database

import (
	"github.com/sancheschris/goal-planner/internal/entity"
	"gorm.io/gorm"
)

type Goal struct {
	DB *gorm.DB
}

func NewGoal(db *gorm.DB) *Goal {
	return &Goal{DB: db}
}

func (g *Goal) Create(goal *entity.Goal) error {
	return g.DB.Create(goal).Error
}

func (g *Goal) FindAll() ([]entity.Goal, error) {
	var goals []entity.Goal
	var err error
	return goals, err
}
