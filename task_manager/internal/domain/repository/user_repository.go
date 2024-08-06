package repository

import "github.com/kika1s1/task_manager/internal/domain/models"


type UserRepository interface {
    Create(user *models.User) error
    GetByUsername(username string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
    Save(user *models.User) error
}

