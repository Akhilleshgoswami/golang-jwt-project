package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-api/models"
	"todo-api/utils"
)

var todos []models.Todo

var nextID = 1

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {

	var newTodo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := utils.AddTodoToDB(&newTodo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	responseJson, err := json.Marshal(newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(responseJson)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("todo item create:", newTodo)
}

func ListTodosHandler(w http.ResponseWriter, r *http.Request) {

	todos , err := utils.GetAllTodosFromDB()
	if err !=  nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return 
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")

	responseJson,err := json.Marshal(todos)

	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return 
	}

	_,err  = w.Write(responseJson)

	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return 
	}

	fmt.Println("Fethced All todos")
}
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {

}
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {

}
func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Implement get todo by ID functionality here...
}
