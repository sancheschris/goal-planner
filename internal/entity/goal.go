package entity

import (
	"errors"

	"github.com/sancheschris/goal-planner/pkg/entity"
)

var (
	ErrGoalIsRequired = errors.New("Goal is required")
	ErrStatusIsRequired = errors.New("Status is required")
)

type Goal struct {
    ID     entity.ID `json:"id" gorm:"primaryKey"` // make it the PK and use ID (not Id)
    Goal   string `json:"goal"`
    Status string `json:"status"`
    Tasks  []Task `json:"tasks" gorm:"foreignKey:GoalID;constraint:OnDelete:CASCADE"`
}


func NewGoal(goal string, status string, tasks []Task) (*Goal, error) {
	newGoal := &Goal{
		ID:     entity.NewId(),
		Goal:   goal,
		Status: status,
		Tasks:  tasks,
	}
	err := newGoal.Validate()
	if err != nil {
		return nil, err
	}
	return newGoal, nil
}

func (g *Goal) Validate() error {
	if g.Goal == "" {
		return ErrGoalIsRequired
	}
	if g.Status == "" {
		return ErrStatusIsRequired
	}
	return nil
}
