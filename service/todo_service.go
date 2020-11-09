package service

import (
	"time"

	"github.com/Prithvipal/go-rest/dal"
	"github.com/Prithvipal/go-rest/dto"
	"github.com/Prithvipal/go-rest/entity"
)

// CreateToDo ...
func CreateToDo(todoDTO dto.TodoDTO) error {
	todo := entity.Todo{
		Value:     todoDTO.Value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return dal.SaveTodo(todo)
}
