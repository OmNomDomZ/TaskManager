package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

var tasks []Task

func init() {
	tasks = []Task{
		{ID: 1, Title: "First Task", Description: "Simple task 1", CreatedAt: time.Now()},
		{ID: 2, Title: "Second Task", Description: "Simple task 2", CreatedAt: time.Now()},
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			e := json.NewEncoder(w).Encode(task)
			if e != nil {
				panic(e)
			}
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w).Encode(tasks)
	if e != nil {
		panic(e)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var t Task
	e := json.NewDecoder(r.Body).Decode(&t)
	if e != nil {
		panic(e)
	}
	t.ID = len(tasks) + 1
	t.CreatedAt = time.Now()
	tasks = append(tasks, t)

	w.Header().Set("Content-Type", "application/json")

	e = json.NewEncoder(w).Encode(tasks)
	if e != nil {
		panic(e)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var t Task
	e := json.NewDecoder(r.Body).Decode(&t)
	if e != nil {
		panic(e)
	}

	for _, task := range tasks {
		if task.ID == id {
			task.Title = t.Title
			task.Description = t.Description
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
