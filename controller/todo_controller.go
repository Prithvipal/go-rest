package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Prithvipal/go-rest/dto"
	"github.com/Prithvipal/go-rest/service"
	"github.com/gorilla/mux"
)

// CreateTodo ...
func CreateTodo(w http.ResponseWriter, req *http.Request) {
	var todoDTO dto.TodoBaseDTO
	err := json.NewDecoder(req.Body).Decode(&todoDTO)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.CreateTodo(todoDTO)
	if err != nil {
		writeSerivceError(w, err)
		return
	}
	msg := dto.Message{Msg: "TODO created successfully"}
	writeJSON(w, msg)
}

// GetTodoList ...
func GetTodoList(w http.ResponseWriter, req *http.Request) {
	pageable, err := getPageable(req, 0, 10)
	if err != nil {
		log.Println("Error while parsing pagination details: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	todos, err := service.GetTodoList(pageable)
	if err != nil {
		log.Println("Error while getting TODO List", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, todos)
}

// GetTodoByID ...
func GetTodoByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID := vars["ID"]
	todo, err := service.GetTodoByID(ID)
	if err != nil {
		writeSerivceError(w, err)
		return
	}
	writeJSON(w, todo)
}

// GetTodoDetailByID handler/api function to return Todo Detail by ID
func GetTodoDetailByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID := vars["ID"]
	todo, err := service.GetTodoDetailByID(ID)
	if err != nil {
		writeSerivceError(w, err)
		return
	}
	writeJSON(w, todo)
}

// DeleteTodoByID ...
func DeleteTodoByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID := vars["ID"]
	err := service.DeleteTodoByID(ID)
	if err != nil {
		writeSerivceError(w, err)
		return
	}
	msg := dto.Message{Msg: "TODO deleted successfully"}
	writeJSON(w, msg)
}

// UpdateTodoByID ...
func UpdateTodoByID(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	ID := vars["ID"]
	var todoDTO dto.TodoBaseDTO
	err := json.NewDecoder(req.Body).Decode(&todoDTO)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.UpdateTodoByID(ID, todoDTO)
	if err != nil {
		writeSerivceError(w, err)
		return
	}
	msg := dto.Message{Msg: "TODO updated successfully"}
	writeJSON(w, msg)
}
