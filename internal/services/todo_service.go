package services

import "github.com/KaitoHasei/go-todo-api/internal/models"

type TodoService interface {
	GetAll() ([]models.Todo, error)
	GetById(id int) (*models.Todo, error)
	Create(title string, desc string) error
	Update(id int, title string, desc string, completed bool) error
}
