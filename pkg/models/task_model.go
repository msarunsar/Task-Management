package models

import "task-management/utilities/standard"

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}

type TaskRequest struct {
	ID          int
	Title       string `faker:"sentence"`
	Description string `faker:"paragraph"`
}

type TaskResponse struct {
	standard.StandardReponse
	Data interface{}
}
