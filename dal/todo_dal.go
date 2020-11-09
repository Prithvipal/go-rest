package dal

import "github.com/Prithvipal/go-rest/entity"

var todoList []entity.Todo

// SaveTodo ...
func SaveTodo(todo entity.Todo) error {
	todoList = append(todoList, todo)
	return nil
}

// GetTodoList ...
func GetTodoList() ([]entity.Todo, error) {
	return todoList, nil
}
