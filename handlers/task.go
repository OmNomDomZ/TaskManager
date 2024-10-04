package handlers

import (
	"TaskManager/db"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

//var tasks []Task
//
//func init() {
//	tasks = []Task{
//		{ID: 1, Title: "First Task", Description: "Simple task 1", CreatedAt: time.Now()},
//		{ID: 2, Title: "Second Task", Description: "Simple task 2", CreatedAt: time.Now()},
//	}
//}

func GetTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var task Task

	if err := db.DB.First(&task, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w).Encode(task)
	if e != nil {
		panic(e)
	}
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	if err := db.DB.Find(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w).Encode(tasks)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var task Task
	e := json.NewDecoder(r.Body).Decode(&task)
	if e != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	task.CreatedAt = time.Now()

	if err := db.DB.Create(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	e = json.NewEncoder(w).Encode(task)
	if e != nil {
		panic(e)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := db.DB.Delete(&Task{}, id); err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var task Task
	if err := db.DB.First(&task, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if e := json.NewDecoder(r.Body).Decode(&task); e != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := db.DB.Save(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
