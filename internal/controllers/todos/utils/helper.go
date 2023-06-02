package utils

import (
	"errors"

	"example.com/restful-api-example/internal/models"
	"gorm.io/gorm"
)

func GetTodoByID(data_todos *gorm.DB, id string) (*models.Todo, error) {
	todo := new(models.Todo)

	res := data_todos.Where("id = ?", id).Table("todos").Find(todo)

	if res.Error == nil {
		if res.RowsAffected <= 0 {
			return nil, errors.New("todo not found")
		}
		return todo, nil
	}

	return nil, res.Error
}

func GetTodoByOwner(data_todos *gorm.DB, owner string) ([]models.Todo, error) {
	var todo []models.Todo

	res := data_todos.Where("owner = ?", owner).Table("todos").Find(&todo)
	if res.Error == nil {
		if res.RowsAffected <= 0 {
			return nil, errors.New("todo not found")
		}
		return todo, nil
	}

	return nil, res.Error
}
