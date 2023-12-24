package rest

import (
	"net/http"
	"strconv"
	"task-management/pkg/handlers"
	"task-management/pkg/models"
	"task-management/pkg/services"
	"task-management/utilities/standard"

	"github.com/labstack/echo/v4"
)

type TaskManagementHandler struct {
	taskManagementSrv services.TaskManagementService
}

func NewTaskManagementHandler(taskManagementSrv services.TaskManagementService) handlers.TaskManagementAPI {
	return TaskManagementHandler{taskManagementSrv: taskManagementSrv}
}

// @Summary CreateAndUpdateTask
// @Description Get a sample response from the server
// @ID post-task
// @Accept json
// @Produce  json
// @Param request body models.TaskRequest true "Request Body if send ID that will become UPDATE API"
// @Success 201 {object} []models.Task "Resource created successfully"
// @Router /task [POST]
func (tm TaskManagementHandler) CreateAndUpdateTask(c echo.Context) error {
	var bodyRequest models.Task
	if err := c.Bind(&bodyRequest); err != nil {
		return c.JSONPretty(http.StatusBadRequest, standard.BadRequest(err.Error()), "")
	}
	taskList, err := tm.taskManagementSrv.CreateAndUpdateTask(bodyRequest)
	if err != nil {
		return c.JSONPretty(http.StatusNotFound, standard.NotFound(err.Error()), "")
	}
	return c.JSONPretty(http.StatusCreated, models.TaskResponse{
		StandardReponse: standard.CreatedOrUpdatedStatus("Created"),
		Data:            taskList,
	}, "")
}

// @Summary UpdateTaskStatus
// @Description Get a sample response from the server
// @ID put-task-status
// @Produce  json
// @Param task_id query int true "Resource ID"
// @Success 201 {object} []models.Task "Resource updated successfully"
// @Router /task [PUT]
func (tm TaskManagementHandler) UpdateTaskStatus(c echo.Context) error {
	taskIDQ := c.QueryParam("task_id")
	taskID, err := strconv.Atoi(taskIDQ)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, standard.BadRequest(err.Error()), "")
	}
	taskList, err := tm.taskManagementSrv.UpdateTaskStatus(taskID)
	if err != nil {
		return c.JSONPretty(http.StatusNotFound, standard.NotFound(err.Error()), "")
	}
	return c.JSONPretty(http.StatusCreated, models.TaskResponse{
		StandardReponse: standard.CreatedOrUpdatedStatus("Updated"),
		Data:            taskList,
	}, "")
}

// @Summary GetTask
// @Description Get a sample response from the server
// @ID get-task
// @Produce  json
// @Param task_id query int true "Resource ID"
// @Success 200 {object} models.Task "Resource found"
// @Router /task [GET]
func (tm TaskManagementHandler) GetTask(c echo.Context) error {
	taskIDQ := c.QueryParam("task_id")
	taskID, err := strconv.Atoi(taskIDQ)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, standard.BadRequest(err.Error()), "")
	}
	task, err := tm.taskManagementSrv.GetTask(taskID)
	if err != nil {
		return c.JSONPretty(http.StatusNotFound, standard.NotFound(err.Error()), "")
	}
	return c.JSONPretty(http.StatusCreated, models.TaskResponse{
		StandardReponse: standard.OkStatus(),
		Data:            task,
	}, "")
}

// @Summary DeleteTask
// @Description Get a sample response from the server
// @ID put-task-status
// @Produce  json
// @Param task_id query int true "Resource ID"
// @Success 200 {object} []models.Task "Resource deleted successfully"
// @Router /task [DELETE]
func (tm TaskManagementHandler) DeleteTask(c echo.Context) error {
	taskIDQ := c.QueryParam("task_id")
	taskID, err := strconv.Atoi(taskIDQ)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, standard.BadRequest(err.Error()), "")
	}
	taskList, err := tm.taskManagementSrv.DeleteTask(taskID)
	if err != nil {
		return c.JSONPretty(http.StatusNotFound, standard.NotFound(err.Error()), "")
	}
	return c.JSONPretty(http.StatusCreated, models.TaskResponse{
		StandardReponse: standard.OkStatus(),
		Data:            taskList,
	}, "")
}

// @Summary GetTasksList
// @Description Get a sample response from the server
// @ID get-task-list
// @Produce  json
// @Success 200 {object} []models.Task "Resource found"
// @Router /task-list [GET]
func (tm TaskManagementHandler) GetTasksList(c echo.Context) error {
	taskList := tm.taskManagementSrv.GetTasksList()
	return c.JSONPretty(http.StatusCreated, models.TaskResponse{
		StandardReponse: standard.OkStatus(),
		Data:            taskList,
	}, "")
}
