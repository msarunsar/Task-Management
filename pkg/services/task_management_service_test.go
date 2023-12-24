package services

import (
	"sync"
	"task-management/pkg/models"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/suite"
)

type TaskManagementTestSuit struct {
	suite.Suite
	underTest TaskManagementService
}

func (s *TaskManagementTestSuit) SetupTest() {
	var mutex sync.Mutex
	s.underTest = NewTaskManagementService(&mutex)
}

func TestTaskManagementTestSuite(t *testing.T) {
	suite.Run(t, new(TaskManagementTestSuit))
}

func (s *TaskManagementTestSuit) TestTaskManagement() {
	createRequest := models.Task{
		Title:       faker.Sentence(),
		Description: faker.Paragraph(),
	}
	var testCreate = struct {
		mock   models.Task
		expect models.Task
	}{
		mock: models.Task{
			Title:       createRequest.Title,
			Description: createRequest.Description,
		},
		expect: models.Task{
			ID:          1,
			Title:       createRequest.Title,
			Description: createRequest.Description,
			Status:      "To Do",
		},
	}
	s.Run("create", func() {
		result, err := s.underTest.CreateAndUpdateTask(testCreate.mock)
		s.NoError(err)
		s.Equal(testCreate.expect, result[0])
	})

	s.Run("find_by_id", func() {
		result, err := s.underTest.GetTask(1)
		s.NoError(err)
		s.Equal(testCreate.expect, result)
	})

	s.Run("find_by_id_error", func() {
		_, err := s.underTest.GetTask(100)
		s.Error(err)
	})

	s.Run("find", func() {
		result := s.underTest.GetTasksList()
		s.Equal(testCreate.expect, result[0])
	})

	updateRequest := models.Task{
		ID:          1,
		Title:       faker.Sentence(),
		Description: faker.Paragraph(),
	}

	var testUpdate = struct {
		mock   models.Task
		expect models.Task
	}{
		mock: models.Task{
			ID:          updateRequest.ID,
			Title:       updateRequest.Title,
			Description: updateRequest.Description,
		},
		expect: models.Task{
			ID:          1,
			Title:       updateRequest.Title,
			Description: updateRequest.Description,
			Status:      "To Do",
		},
	}

	s.Run("update", func() {
		result, err := s.underTest.CreateAndUpdateTask(testUpdate.mock)
		s.NoError(err)
		s.Equal(testUpdate.expect, result[0])
	})

	s.Run("update_error", func() {
		testUpdate.mock.ID = 100
		_, err := s.underTest.CreateAndUpdateTask(testUpdate.mock)
		s.Error(err)
	})

	var testUpdateStatus = struct {
		mock   models.Task
		expect models.Task
	}{
		mock: models.Task{
			ID: 1,
		},
		expect: models.Task{
			ID:          1,
			Title:       updateRequest.Title,
			Description: updateRequest.Description,
			Status:      "In Progress",
		},
	}

	s.Run("update_status", func() {
		result, err := s.underTest.UpdateTaskStatus(testUpdateStatus.mock.ID)
		s.NoError(err)
		s.Equal(testUpdateStatus.expect, result[0])
	})

	s.Run("update_status_error", func() {
		_, err := s.underTest.UpdateTaskStatus(100)
		s.Error(err)
	})

	var testDelete = struct {
		mock   models.Task
		expect []models.Task
	}{
		mock: models.Task{
			ID: 1,
		},
		expect: []models.Task{},
	}

	s.Run("delete", func() {
		result, err := s.underTest.DeleteTask(testDelete.mock.ID)
		s.NoError(err)
		s.Equal(testDelete.expect, result)
	})

	s.Run("delete_error", func() {
		_, err := s.underTest.DeleteTask(100)
		s.Error(err)
	})

}
