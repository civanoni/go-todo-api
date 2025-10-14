package services

import (
	"github.com/KaitoHasei/go-todo-api/internal/models"
	"github.com/KaitoHasei/go-todo-api/internal/repositories"
)

type todoServiceImpl struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoServiceImpl{repo: repo}
}

func (s *todoServiceImpl) GetAll() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoServiceImpl) GetById(id int) (*models.Todo, error) {
	return s.repo.GetById(id)
}

func (s *todoServiceImpl) Create(title string, desc string) error {
	todo := &models.Todo{
		Title:       title,
		Description: desc,
	}

	return s.repo.Create(todo)
}

func (s *todoServiceImpl) Update(id int, title string, desc string, completed bool) error {
	todo, err := s.repo.GetById(id)

	if err != nil {
		return err
	}

	todo.Title = title
	todo.Description = desc
	todo.Completed = completed

	return s.repo.Update(todo)
}
