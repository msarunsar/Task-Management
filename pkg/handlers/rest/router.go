package rest

import (
	"sync"
	"task-management/pkg/services"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	main := e.Group("/task-management")
	api := main.Group("/api")
	g := api.Group("/v1")

	var mutex sync.Mutex
	taskManagementSrv := services.NewTaskManagementService(&mutex)
	initTaskManagementRouter(g, taskManagementSrv)
	return e
}

func initTaskManagementRouter(e *echo.Group, taskManagementSrv services.TaskManagementService) {
	handler := NewTaskManagementHandler(taskManagementSrv)

	e.POST("/task", handler.CreateAndUpdateTask)

	e.PUT("/task", handler.UpdateTaskStatus)

	e.GET("/task", handler.GetTask)

	e.DELETE("/task", handler.DeleteTask)

	e.GET("/task-list", handler.GetTasksList)
}
