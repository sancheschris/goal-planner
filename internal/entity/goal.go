package entity

type Goal struct {
	Goal string `json:"goal"`
	Status string `json:"status"`
	Task []Task `json:"tasks"`
}

