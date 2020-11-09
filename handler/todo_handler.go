package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Prithvipal/go-rest/dto"
	"github.com/Prithvipal/go-rest/service"
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
		log.Println("Error while creating TODO", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// GetTodoList ...
func GetTodoList(w http.ResponseWriter, req *http.Request) {
	todos, err := service.GetTodoList()
	if err != nil {
		log.Println("Error while getting TODO List", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	writeJSON(w, todos)
}

func writeJSON(w http.ResponseWriter, records interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(records)
	if err != nil {
		log.Println("Error while getting TODO List", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write(data)
}
