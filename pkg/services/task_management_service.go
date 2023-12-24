package services

import (
	"fmt"
	"sort"
	"sync"
	"task-management/pkg/models"
	"task-management/pkg/repositories"
)

var (
	TaskStatus = map[int]string{
		1: "To Do",
		2: "In Progress",
		3: "Done",
	}
)

type TaskManagement struct {
	mu *sync.Mutex
}

func NewTaskManagementService(mutex *sync.Mutex) TaskManagementService {
	return TaskManagement{mu: mutex}
}

func (tm TaskManagement) CreateAndUpdateTask(task models.Task) ([]models.Task, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if repositories.Task[task.ID] != (models.Task{}) {
		repositories.Task[task.ID] = models.Task{
			ID:          repositories.Task[task.ID].ID,
			Description: task.Description,
			Title:       task.Title,
			Status:      repositories.Task[task.ID].Status,
		}
	} else if task.ID != 0 {
		return nil, fmt.Errorf("Task ID:%d NOT FOUND", task.ID)
	} else {
		var id = repositories.ID
		repositories.Task[repositories.ID] = models.Task{
			ID:          id,
			Description: task.Description,
			Title:       task.Title,
			Status:      TaskStatus[1],
		}
		repositories.ID++
	}
	return mapStructToArray(repositories.Task), nil
}

func (tm TaskManagement) UpdateTaskStatus(taskID int) ([]models.Task, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if repositories.Task[taskID] != (models.Task{}) {
		status := ""
		switch repositories.Task[taskID].Status {
		case TaskStatus[1]:
			status = TaskStatus[2]
		case TaskStatus[2]:
			status = TaskStatus[3]
		default:
			return mapStructToArray(repositories.Task), nil
		}
		repositories.Task[taskID] = models.Task{
			ID:          repositories.Task[taskID].ID,
			Description: repositories.Task[taskID].Description,
			Title:       repositories.Task[taskID].Title,
			Status:      status,
		}
		return mapStructToArray(repositories.Task), nil
	} else {
		return nil, fmt.Errorf("Task ID:%d NOT FOUND", taskID)
	}
}

func (tm TaskManagement) GetTask(taskID int) (models.Task, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if repositories.Task[taskID] != (models.Task{}) {
		return repositories.Task[taskID], nil
	} else {
		return models.Task{}, fmt.Errorf("Task ID:%d NOT FOUND", taskID)
	}
}

func (tm TaskManagement) DeleteTask(taskID int) ([]models.Task, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if repositories.Task[taskID] != (models.Task{}) {
		delete(repositories.Task, taskID)
		return mapStructToArray(repositories.Task), nil
	} else {
		return nil, fmt.Errorf("Task ID:%d NOT FOUND", taskID)
	}
}

func (tm TaskManagement) GetTasksList() []models.Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	return mapStructToArray(repositories.Task)
}

func mapStructToArray(mapStruct map[int]models.Task) []models.Task {
	var arrayFromMap []models.Task
	for _, value := range mapStruct {
		arrayFromMap = append(arrayFromMap, value)
	}
	array := make([]models.Task, len(arrayFromMap))
	copy(array, arrayFromMap)
	sort.SliceStable(array, func(i, j int) bool {
		return array[i].ID < array[j].ID
	})
	return array
}
