package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// TodoDTO ..
type TodoDTO struct {
	Value string `json:"value"`
}

// CreateToDo ...
func CreateToDo(w http.ResponseWriter, req *http.Request) {
	var todoDTO TodoDTO
	err := json.NewDecoder(req.Body).Decode(&todoDTO)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(todoDTO)
}
