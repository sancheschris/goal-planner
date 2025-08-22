package entity

import "github.com/sancheschris/goal-planner/pkg/entity"

type Goal struct {
    ID     entity.ID `json:"id" gorm:"primaryKey"` // make it the PK and use ID (not Id)
    Goal   string `json:"goal"`
    Status string `json:"status"`
    Tasks  []Task `json:"tasks" gorm:"foreignKey:GoalID;constraint:OnDelete:CASCADE"`
}


func NewGoal(goal string, status string, tasks []Task) *Goal {
	return &Goal{
		ID: entity.NewId(),
		Goal: goal,
		Status: status,
		Tasks: tasks,
	}
}
