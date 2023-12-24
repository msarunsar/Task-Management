package handlers

import "github.com/labstack/echo/v4"

type TaskManagementAPI interface {
	CreateAndUpdateTask(c echo.Context) error
	UpdateTaskStatus(c echo.Context) error
	GetTask(c echo.Context) error
	DeleteTask(c echo.Context) error
	GetTasksList(c echo.Context) error
}
