package repositories

import (
	"github.com/KaitoHasei/go-todo-api/internal/models"
	"gorm.io/gorm"
)

type todoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository { // DI pattern: Inject DB qua constructor
	return &todoRepositoryImpl{db: db}
}

func (r *todoRepositoryImpl) GetAll() ([]models.Todo, error) {
	var todos []models.Todo

	return todos, r.db.Find(&todos).Error
}

func (r *todoRepositoryImpl) GetById(id int) (*models.Todo, error) {
	var todo models.Todo

	err := r.db.First(&todo, id).Error

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *todoRepositoryImpl) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepositoryImpl) Update(todo *models.Todo) error {
	return r.db.Save(todo).Error
}
