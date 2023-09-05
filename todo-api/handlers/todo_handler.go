package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-api/models"
	"todo-api/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var todos []models.Todo

var nextID = 1

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {

	var newTodo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTodo.ID = primitive.NilObjectID
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

	todos, err := utils.GetAllTodosFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	responseJson, err := json.Marshal(todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(responseJson)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Fethced All todos")
}

// DeleteTodoHandler handles the deletion of a todo item by ID.
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.DeleteTodoFromDB(todoID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

	// Implement get todo by ID functionality here...
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updateTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updateTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updateTodo.ID = todoID
	if err := utils.UpdateTodoInDB(todoID, &updateTodo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	responseJSON, err := json.Marshal(updateTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(responseJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("todo item updated: %+v\n", updateTodo)
}
