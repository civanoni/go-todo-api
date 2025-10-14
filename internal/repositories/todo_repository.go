package repositories

import "github.com/KaitoHasei/go-todo-api/internal/models"

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	GetById(id int) (*models.Todo, error)
	Create(todo *models.Todo) error
	Update(todo *models.Todo) error
}
