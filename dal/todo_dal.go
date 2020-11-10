package dal

import (
	"strconv"

	"github.com/Prithvipal/go-rest/entity"
)

var todoList []entity.Todo

// SaveTodo ...
func SaveTodo(todo entity.Todo) error {
	todo.ID = strconv.Itoa(len(todoList) + 1)
	todoList = append(todoList, todo)
	return nil
}

// GetTodoList ...
func GetTodoList() ([]entity.Todo, error) {
	return todoList, nil
}
