package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGoal(t *testing.T) {
	tasks := []Task{{Name: "Learn Simple Past", Status: "Doing"}}
	goal, err := NewGoal("Learn English", "Todo", tasks)
	assert.Nil(t, err)
	assert.NotNil(t, goal)
	assert.NotEmpty(t, goal.Goal)
	assert.Equal(t, "Learn English", goal.Goal)
	assert.Equal(t, "Todo", goal.Status)
	assert.Equal(t, "Learn Simple Past", goal.Tasks[0].Name)
	assert.Equal(t, "Doing", goal.Tasks[0].Status)
}

func TestGoalWhenGoalIsRequired(t *testing.T) {
	g, err := NewGoal("", "Todo", []Task{})
	assert.Nil(t, g)
	assert.Equal(t, ErrGoalIsRequired, err)
}

func TestGoalWhenStatusIsRequired (t *testing.T) {
	g, err := NewGoal("Study DSA", "", []Task{})
	assert.Nil(t, g)
	assert.Equal(t, ErrStatusIsRequired, err)
}