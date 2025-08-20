package database

import "github.com/sancheschris/goal-planner/internal/entity"

type GoalInterface interface {
	Create(goal *entity.Goal) error
	FindAll() ([]entity.Goal, error)
	Update(goal *entity.Goal) error
}