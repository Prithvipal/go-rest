package dal

import (
	"strconv"

	"github.com/Prithvipal/go-rest/apierrors"
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

// GetTodoByID ...
func GetTodoByID(ID string) (entity.Todo, error) {
	for _, todo := range todoList {
		if todo.ID == ID {
			return todo, nil
		}
	}
	return entity.Todo{}, apierrors.ErrNotFount
}

// DeleteTodoByID ...
func DeleteTodoByID(ID string) error {
	index := getIndex(ID)
	if index == -1 {
		return apierrors.ErrNotFount
	}
	todoList = append(todoList[:index], todoList[index+1:]...)
	return nil
}

func getIndex(ID string) int {
	for indx, todo := range todoList {
		if todo.ID == ID {
			return indx
		}
	}
	return -1
}

// UpdateTodoByID to update TODO by id in data access layer
func UpdateTodoByID(ID string, todo entity.Todo) error {
	index := getIndex(ID)
	if index == -1 {
		return apierrors.ErrNotFount
	}
	todoList[index].UpdatedAt = todo.UpdatedAt
	todoList[index].Value = todo.Value
	return nil
}
