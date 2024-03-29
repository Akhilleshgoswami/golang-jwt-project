package routes

import (
	"github.com/gorilla/mux"
	"todo-api/handlers"
)

func SetupRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.CreateTodoHandler).Methods("POST")
	r.HandleFunc("/todos", handlers.ListTodosHandler).Methods("GET")
	// r.HandleFunc("/todos/{id}", handlers.GetTodoHandler).Methods("GET")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodoHandler).Methods("DELETE")
	return r
}
