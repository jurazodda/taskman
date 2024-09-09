package repository

import (
	"context"
	"fmt"
	"taskman/internal/models"
	"time"
)

func (r *Repository) CreateTask(ctx context.Context, params models.CreateTaskParams) (task *models.Task, err error) {
	task = &models.Task{
		Title:       &params.Title,
		Description: &params.Description,
	}
	err = r.db.WithContext(ctx).Create(&task).Error
	return
}

func (r *Repository) GetTaskByID(ctx context.Context, taskID int) (task *models.Task, err error) {
	err = r.db.WithContext(ctx).First(&task, taskID).Error
	return
}

func (r *Repository) GetTasks(ctx context.Context) (tasks []*models.Task, err error) {
	err = r.db.WithContext(ctx).Find(&tasks).Error
	return
}

func (r *Repository) UpdateTask(ctx context.Context, taskID int, params models.UpdateTaskParams) (task *models.Task, err error) {
	task = &models.Task{
		ID:          taskID,
		Title:       &params.Title,
		Description: &params.Description,
		UpdatedAt:   time.Now(),
	}

	err = r.db.WithContext(ctx).Model(&models.Task{}).Where("id = ?", taskID).Updates(&task).Error
	return
}

func (r *Repository) DeleteTask(ctx context.Context, taskID int) (err error) {
	err = r.db.WithContext(ctx).Delete(&models.Task{}, taskID).Error
	return
}

func (r *Repository) MarkTask(ctx context.Context, taskID int, params models.MarkTaskParams) error {
	err := r.db.WithContext(ctx).Model(&models.Task{}).Where("id = ?", taskID).Update("done", params.IsDone).Error
	if err != nil {
		return fmt.Errorf("failed to mark task: %w", err)
	}
	return nil
}
