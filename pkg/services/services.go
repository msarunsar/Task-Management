package services

import "task-management/pkg/models"

type TaskManagementService interface {
	CreateAndUpdateTask(task models.Task) ([]models.Task, error)
	UpdateTaskStatus(taskID int) ([]models.Task, error)
	GetTask(taskID int) (models.Task, error)
	DeleteTask(taskID int) ([]models.Task, error)
	GetTasksList() []models.Task
}
