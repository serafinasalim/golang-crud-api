package controller

import (
	"database/sql"
	"golang-crud-api/dto"
	"golang-crud-api/service"
	"golang-crud-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service *service.TaskService
}

func NewTaskController(service *service.TaskService) *TaskController {
	return &TaskController{service: service}
}

// @Summary Get All Tasks
// @Description Get a list of all tasks
// @Tags Tasks
// @Accept json
// @Produce json
// @Success 200 {object} utils.APIResponse{data=[]dto.TaskResponse} "Tasks fetched successfully"
// @Success 204 {object} utils.APIResponse{} "Tasks No Record"
// @Failure 500 {object} utils.HTTPError "Failed to fetch tasks"
// @Router /tasks [get]
func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.service.GetAllTasks()
	if err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, utils.ErrInternalServer, err)
		return
	}

	if len(tasks) == 0 {
		utils.RespondSuccess(ctx, http.StatusNoContent, utils.ErrTaskNoRecord, tasks)
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, utils.MsgOk, tasks)
}

// @Summary Create Tasks
// @Description Sample Payload: <br> `{ `<br>` "title": "Fix Bugs", `<br>` "description": "fix multiple bugs in dev", `<br>` "startDate": "2024-12-05T00:00:00Z", `<br>` "deadline": "2024-12-07T00:00:00Z" `<br>` }`
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body dto.TaskRequest true "Task Request Body"
// @Success 201 {object} utils.APIResponse{status=string,message=string,data=string} "Task created successfully"
// @Failure 400 {object} utils.APIResponse{status=string,message=string} "Invalid input"
// @Failure 500 {object} utils.HTTPError "Failed to create task"
// @Router /tasks [post]
func (c *TaskController) CreateTask(ctx *gin.Context) {
	var params dto.TaskRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, utils.ErrInvalidRequest, err)
		return
	}

	res, err := c.service.CreateTask(params)
	if err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, utils.ErrInternalServer, err)
		return
	}

	utils.RespondSuccess(ctx, http.StatusCreated, utils.MsgOk, res)
}

// @Summary Get Task by Uuid
// @Description
// @Tags Tasks
// @Accept json
// @Produce json
// @Param uuid path string true "Task Uuid"
// @Success 200 {object} utils.APIResponse{data=dto.TaskResponse} "Task fetched successfully"
// @Failure 404 {object} utils.APIResponse{status=string,message=string} "Task not found"
// @Failure 500 {object} utils.HTTPError "Failed to fetch task"
// @Router /tasks/{uuid} [get]
func (c *TaskController) GetTaskByUuid(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	task, err := c.service.GetTaskByUuid(uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondError(ctx, http.StatusNotFound, utils.ErrTaskNotFound, err)
			return
		}

		utils.RespondError(ctx, http.StatusInternalServerError, utils.ErrInternalServer, err)
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, utils.MsgOk, task)
}

// @Summary Update Task
// @Description Sample Payload (only send the ones you want to update): <br> `{ `<br>` "description": "fix lots of bugs", `<br>` "completed": true, `<br>` "startDate": "2024-12-05T00:00:00Z", `<br>` "deadline": "2024-12-07T00:00:00Z" `<br>` }`
// @Tags Tasks
// @Accept json
// @Produce json
// @Param uuid path string true "Task Uuid"
// @Param taskUpdate body dto.TaskUpdate true "Updated Task"
// @Success 200 {object} model.Task "Task updated successfully"
// @Failure 400 {object} utils.APIResponse{status=string,message=string} "Invalid input"
// @Failure 404 {object} utils.APIResponse{} "Task not found"
// @Router /tasks/{uuid} [patch]
func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("uuid")

	var taskUpdate dto.TaskUpdate
	if err := ctx.ShouldBindJSON(&taskUpdate); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, utils.ErrInvalidRequest, err)
		return
	}

	err := c.service.UpdateTask(id, taskUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondError(ctx, http.StatusNotFound, utils.ErrTaskNotFound, err)
			return
		}

		utils.RespondError(ctx, http.StatusInternalServerError, utils.ErrInternalServer, err)
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, utils.MsgOk, "Update Successful")
}

// @Summary Delete Task
// @Description
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param uuid path string true "Task Uuid"
// @Success 200 {object} utils.APIResponse "Task deleted successfully"
// @Failure 400 {object} utils.APIResponse{status=string,message=string} "Invalid request"
// @Failure 404 {object} utils.APIResponse{status=string,message=string} "Task not found"
// @Router /tasks/{uuid} [delete]
func (c *TaskController) DeleteTask(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	err := c.service.DeleteTask(uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondError(ctx, http.StatusNotFound, utils.ErrTaskNotFound, err)
			return
		}

		utils.RespondError(ctx, http.StatusInternalServerError, utils.ErrInternalServer, err)
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, utils.MsgOk, "Delete Successful")
}
