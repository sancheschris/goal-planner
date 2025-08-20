package database

import (
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
	// err = goalDB.Create()
}

