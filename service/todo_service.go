package service

import (
	"time"

	"github.com/Prithvipal/go-rest/dal"
	"github.com/Prithvipal/go-rest/dto"
	"github.com/Prithvipal/go-rest/entity"
)

// CreateTodo ...
func CreateTodo(todoDTO dto.TodoBaseDTO) error {
	todo := entity.Todo{
		Value:     todoDTO.Value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return dal.SaveTodo(todo)
}

// GetTodoList ...
func GetTodoList() (todos []dto.TodoDTO, err error) {
	todoList, err := dal.GetTodoList()
	if err != nil {
		return
	}
	for _, todo := range todoList {
		todoDto := dto.TodoDTO{
			ID: todo.ID,
		}
		todoDto.Value = todo.Value
		todos = append(todos, todoDto)
	}
	return
}
