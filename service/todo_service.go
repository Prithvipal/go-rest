package service

import (
	"log"
	"time"

	"github.com/Prithvipal/go-rest/dal"
	"github.com/Prithvipal/go-rest/dto"
	"github.com/Prithvipal/go-rest/entity"
	"github.com/Prithvipal/go-rest/models"
)

// CreateTodo ...
func CreateTodo(todoDTO dto.TodoBaseDTO) (err error) {
	todo := entity.Todo{
		Value:     todoDTO.Value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = dal.SaveTodo(todo)
	if err != nil {
		log.Printf("Error while saving Todo: %v", err)
	}
	return
}

// GetTodoList ...
func GetTodoList(pageable models.Pageable) (todoPageable dto.TodoPageableDto, err error) {
	page := pageable.Page * pageable.Limit
	todoList, err := dal.GetTodoList(page, pageable.Limit)
	if err != nil {
		log.Printf("Error while geting Todo List: %v", err)
		return
	}
	var todos []dto.TodoDTO
	for _, todo := range todoList {
		todoDto := dto.TodoDTO{
			ID: todo.ID,
		}
		todoDto.Value = todo.Value
		todos = append(todos, todoDto)
	}
	total, err := dal.GetTodoCount()
	if err != nil {
		log.Printf("Error while geting Todo Count: %v", err)
		return
	}

	todoPageable.TodoList = todos
	todoPageable.Limit = pageable.Limit
	todoPageable.Page = pageable.Page
	todoPageable.Count = len(todos)
	todoPageable.TotalCount = total
	todoPageable.TotalPages = total / pageable.Limit

	return
}

// GetTodoByID ...
func GetTodoByID(ID string) (todoDto dto.TodoDTO, err error) {
	todo, err := dal.GetTodoByID(ID)
	if err != nil {
		log.Printf("Error while getting Todo by ID[%v]: %v", ID, err)
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
		log.Printf("Error while getting Todo Details by ID[%v]: %v", ID, err)
		return
	}
	todoDto.ID = todo.ID
	todoDto.Value = todo.Value
	todoDto.CreatedAt = todo.CreatedAt
	todoDto.UpdatedAt = todo.UpdatedAt
	return
}

// DeleteTodoByID ...
func DeleteTodoByID(ID string) (err error) {
	err = dal.DeleteTodoByID(ID)
	if err != nil {
		log.Printf("Error while deleting Todo By ID[%v]: %v", ID, err)
	}
	return
}

// UpdateTodoByID a service method to update Todo
func UpdateTodoByID(ID string, todoDTO dto.TodoBaseDTO) (err error) {
	todo := entity.Todo{}
	todo.Value = todoDTO.Value
	todo.UpdatedAt = time.Now()
	err = dal.UpdateTodoByID(ID, todo)
	if err != nil {
		log.Printf("Error while updating Todo By ID[%v]: %v", ID, err)
	}
	return
}
