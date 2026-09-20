package service

import (
	"github.com/yuki-829/ToDo_GOtra1/internal/model"
	"github.com/yuki-829/ToDo_GOtra1/internal/repository"
)

func CreateTodo(task string) error {
	return repository.Create(task)
}

func GetTodos() ([]model.Todo, error) {
	return repository.FindAll()
}
