package entity

type Status string

type Task struct {
	Name string `json:"name"`
	Status Status
}

type GoalPlanner struct {
	Goal string `json:"goal"`
	Status Status `json:"status"`
	Task []Task `json:"tasks"`
}

