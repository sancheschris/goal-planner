package entity

import "github.com/sancheschris/goal-planner/pkg/entity"

type Goal struct {
	Id entity.ID `json:"id"`
	Goal string `json:"goal"`
	Status string `json:"status"`
	Tasks []Task `json:"tasks"`
}


func NewGoal(goal string, status string, tasks []Task) *Goal {
	return &Goal{
		Id: entity.NewId(),
		Goal: goal,
		Status: status,
		Tasks: tasks,
	}
	
}
