package main

import (
	"TaskManager/db"
	"TaskManager/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db.InitDB()

	fmt.Println("Database connected successfully")

	r := mux.NewRouter()

	r.HandleFunc("/tasks", handlers.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/create", handlers.Create).Methods("POST")
	r.HandleFunc("/delete/{id}", handlers.Delete).Methods("DELETE")
	r.HandleFunc("/update/{id}", handlers.Update).Methods("PUT")

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
