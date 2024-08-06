package repository

import "github.com/kika1s1/task_manager/internal/domain/models"

type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id string) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id string) error
	GetAll() ([]*models.Task, error)
}
