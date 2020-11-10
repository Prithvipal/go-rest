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

// GetTodoByID ...
func GetTodoByID(ID string) (todoDto dto.TodoDTO, err error) {
	todo, err := dal.GetTodoByID(ID)
	if err != nil {
		return
	}
	todoDto.ID = todo.ID
	todoDto.Value = todo.Value
	return
}

// GetTodoDetailByID a service function to return Todo Detail
func GetTodoDetailByID(ID string) (todoDto dto.TodoDetailDTO, err error) {
	todo, err := dal.GetTodoByID(ID)
	if err != nil {
		return
	}
	todoDto.ID = todo.ID
	todoDto.Value = todo.Value
	todoDto.CreatedAt = todo.CreatedAt
	todoDto.UpdatedAt = todo.UpdatedAt
	return
}

// DeleteTodoByID ...
func DeleteTodoByID(ID string) error {
	return dal.DeleteTodoByID(ID)
}

// UpdateTodoByID a service method to update Todo
func UpdateTodoByID(ID string, todoDTO dto.TodoBaseDTO) (err error) {
	todo := entity.Todo{}
	todo.Value = todoDTO.Value
	todo.UpdatedAt = time.Now()
	return dal.UpdateTodoByID(ID, todo)
}
