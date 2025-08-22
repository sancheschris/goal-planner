package database

import (
	"github.com/sancheschris/goal-planner/internal/entity"
	"gorm.io/gorm"
)

type GoalRepo struct {
	DB *gorm.DB
}

func NewGoal(db *gorm.DB) *GoalRepo {
	return &GoalRepo{DB: db}
}

func (g *GoalRepo) Create(goal *entity.Goal) error {
	return g.DB.Create(goal).Error
}

func (g *GoalRepo) FindAll() ([]entity.Goal, error) {
	var goals []entity.Goal
	err := g.DB.Preload("Tasks").Find(&goals).Error
	return goals, err
}
