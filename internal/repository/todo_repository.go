package repository

import (
	"github.com/yuki-829/ToDo_GOtra1/internal/db"
	"github.com/yuki-829/ToDo_GOtra1/internal/model"
)

func Create(task string) error {
	todo := model.Todo{Task: task}
	return db.DB.Create(&todo).Error
}

func FindAll() ([]model.Todo, error) {
	var todos []model.Todo
	err := db.DB.Find(&todos).Error
	return todos, err
}
