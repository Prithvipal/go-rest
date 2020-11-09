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
	var todoDTO dto.TodoDTO
	err := json.NewDecoder(req.Body).Decode(&todoDTO)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.CreateToDo(todoDTO)
	if err != nil {
		log.Println("Error while creating TODO", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
