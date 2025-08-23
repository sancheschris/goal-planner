package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sancheschris/goal-planner/internal/dto"
	"github.com/sancheschris/goal-planner/internal/entity"
	"github.com/sancheschris/goal-planner/internal/infra/database"
	entityPkg "github.com/sancheschris/goal-planner/pkg/entity"
)

type GoalHandler struct {
	GoalDB database.GoalInterface
}

func NewGoalHandler(db database.GoalInterface) *GoalHandler {
	return &GoalHandler{
		GoalDB: db,
	}
}

func (h *GoalHandler) CreateGoal(w http.ResponseWriter, r *http.Request) {
	var goal dto.GoalInput
	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	g := entity.NewGoal(goal.Goal, goal.Status, goal.Tasks)
	if err != nil {
		http.Error(w, "Error creating goal", http.StatusBadRequest)
		return
	}
	err = h.GoalDB.Create(g)
	if err != nil {
		http.Error(w, "Error saving goal", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *GoalHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	goals, err := h.GoalDB.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(goals)
}

func (h *GoalHandler) GetGoal(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	goal, err := h.GoalDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(goal)
} 

func (h *GoalHandler) UpdateGoal(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Goal ID is required", http.StatusBadRequest)
		return
	}
	var goal entity.Goal
	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return 
	}
	goal.ID, err = entityPkg.ParseID(id)
	if err != nil {
		http.Error(w, "Goal not found", http.StatusNotFound)
		return
	}
	err = h.GoalDB.Update(&goal)
	if err != nil {
		http.Error(w, "Error updating goal", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}