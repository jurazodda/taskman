package service

import (
	"context"
	"taskman/internal/models"
	"taskman/internal/repository"
)

type TaskService struct {
	repo *repository.Repository
}

func NewTaskService(repo *repository.Repository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, params models.CreateTaskParams) (task *models.Task, err error) {
	return s.repo.CreateTask(ctx, params)
}

func (s *TaskService) GetTaskByID(ctx context.Context, taskID int) (task *models.Task, err error) {
	return s.repo.GetTaskByID(ctx, taskID)
}

func (s *TaskService) GetTasks(ctx context.Context) (tasks []*models.Task, err error) {
	return s.repo.GetTasks(ctx)
}

func (s *TaskService) UpdateTask(ctx context.Context, taskID int, params models.UpdateTaskParams) (task *models.Task, err error) {
	return s.repo.UpdateTask(ctx, taskID, params)
}

func (s *TaskService) DeleteTask(ctx context.Context, taskID int) (err error) {
	return s.repo.DeleteTask(ctx, taskID)
}

func (s *TaskService) MarkTask(ctx context.Context, taskID int, params models.MarkTaskParams) error  {
	return s.repo.MarkTask(ctx, taskID, params)
}
