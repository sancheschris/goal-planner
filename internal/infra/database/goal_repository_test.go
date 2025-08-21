package database

import (
	"fmt"
	"testing"

	"github.com/sancheschris/goal-planner/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewGoal(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Goal{})
	goal := entity.NewGoal("Sample Title", "Sample Description", []entity.Task{})
	assert.NoError(t, err)
	assert.NotNil(t, goal.Goal)

	goalDB := NewGoal(db)
	err = goalDB.Create(goal)
	assert.NoError(t, err)
	assert.NotEmpty(t, goal.Goal)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Goal{})
	for i := 1; i < 5; i++ {
		task := entity.Task{
			Name: fmt.Sprintf("Subtask %d", i),
			Status: "Todo",
		}
		goal := entity.NewGoal(fmt.Sprintf("Task %d", i), "Todo", []entity.Task{task})
		assert.NoError(t, err)
		db.Create(goal)
	}
	goalDB := NewGoal(db)
	goals, err := goalDB.FindAll()
	assert.NoError(t, err)
	assert.Len(t, goals, 5)
	assert.Equal(t, "Task 1", goals[0].Goal)
	assert.Equal(t, "Task 2", goals[0].Goal)

}

