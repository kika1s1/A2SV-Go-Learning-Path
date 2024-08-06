package usecase

import (
	"github.com/kika1s1/task_manager/internal/domain/models"
	"github.com/kika1s1/task_manager/internal/domain/repository"
)

type TaskUsecase struct {
    TaskRepo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) *TaskUsecase {
    return &TaskUsecase{
        TaskRepo: repo,
    }
}

func (uc *TaskUsecase) CreateTask(task *models.Task) error {
    return uc.TaskRepo.Create(task)
}

func (uc *TaskUsecase) GetTaskByID(id string) (*models.Task, error) {
    return uc.TaskRepo.GetByID(id)
}

func (uc *TaskUsecase) UpdateTask(task *models.Task) error {
    return uc.TaskRepo.Update(task)
}

func (uc *TaskUsecase) DeleteTask(id string) error {
    return uc.TaskRepo.Delete(id)
}

func (uc *TaskUsecase) GetAllTasks() ([]*models.Task, error) {
    return uc.TaskRepo.GetAll()
}
