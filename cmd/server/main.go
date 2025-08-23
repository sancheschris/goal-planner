package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sancheschris/goal-planner/internal/entity"
	"github.com/sancheschris/goal-planner/internal/infra/database"
	"github.com/sancheschris/goal-planner/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// starting point
	// configs, err := configs.LoadConfig(".")
	// if err != nil {
	// 	panic(err)
	// }
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Goal{}, &entity.Task{})
	goalDB := database.NewGoal(db)
	goalHandler := handlers.NewGoalHandler(goalDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	r.Post("/goals", goalHandler.CreateGoal)
	r.Get("/goals", goalHandler.FindAll)

	http.ListenAndServe(":8080", r)
}

