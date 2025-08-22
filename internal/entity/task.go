package entity

import (
	"github.com/sancheschris/goal-planner/pkg/entity"
)

type Task struct {
    ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
    Name   string `json:"name"`
    Status string `json:"status"`
    GoalID entity.ID `json:"goal_id" gorm:"index"`
}