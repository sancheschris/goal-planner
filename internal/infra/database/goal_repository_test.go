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

func TestFindAllGoals(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil { t.Fatal(err) }
	
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(1) // forces GORM’s connection pool to use only one connection
	_ = db.Exec("PRAGMA foreign_keys = ON").Error // when you have foreign key, this enforces and makes your test closer to reality

	if err := db.AutoMigrate(&entity.Goal{}, &entity.Task{}); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 5; i++ {
		task := entity.Task{
			Name: fmt.Sprintf("Substask %d", i),
			Status: "Todo",
		}   
		goal := entity.NewGoal(fmt.Sprintf("Task %d", i), "Todo", []entity.Task{task})
		if err := db.Create(goal).Error; err != nil {
			t.Fatalf("create goal %d failed: %v", i, err)
		}
	}
	goalRepo := NewGoal(db)
	goals, err := goalRepo.FindAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, goals)
	assert.Len(t, goals, 5)
	
	assert.Equal(t, "Task 0", goals[0].Goal)
	assert.Equal(t, "Task 4", goals[4].Goal)

	for _, g := range goals {
		assert.Len(t, g.Tasks, 1)
		assert.Contains(t, g.Tasks[0].Name, "Substask")
	}
}

func TestUpdateGoal(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Goal{}, &entity.Task{})
	tasks := []entity.Task{
			{Name: "Substask2", Status: "Todo"},
			    {Name: "Subtask 1", Status: "Doing"},
		}   
	goal := entity.NewGoal("Task", "Todo", tasks)
	assert.NoError(t, err)
	db.Create(goal)
	goal.Goal = "Updated Goal"
	goal.Status = "Done"
	goalDB := NewGoal(db)
	err = goalDB.Update(goal)
	assert.NoError(t, err)

	updatedGoal, err := goalDB.FindById(goal.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Updated Goal", updatedGoal.Goal)
	assert.Equal(t, "Done", updatedGoal.Status)
}

func TestGetGoal(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Goal{}, &entity.Task{})
	task := entity.Task{
			Name: "Substask",
			Status: "Todo",
		}   
	goal := entity.NewGoal("Task", "Todo", []entity.Task{task})
	assert.NoError(t, err)
	db.Create(goal)
	goalDB := NewGoal(db)
	goal, err = goalDB.FindById(goal.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Task", goal.Goal)
	assert.Equal(t, "Todo", goal.Status)
	assert.Len(t, goal.Tasks, 1)
	assert.Equal(t, "Substask", goal.Tasks[0].Name)
}