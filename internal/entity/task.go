package entity

type Task struct {
	Name string `json:"name"`
	Status string `json:"status"`
	GoalID string
}