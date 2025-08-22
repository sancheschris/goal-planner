package dto

import "github.com/sancheschris/goal-planner/internal/entity"

type GoalInput struct {
	Goal string `json:"goal"`
	Status string `json:"status"`
	Tasks []entity.Task `json:"tasks"`
}