package controller

import (
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
// @Success 200 {object} utils.APIResponse{data=[]model.Task} "Tasks fetched successfully"
// @Failure 500 {object} utils.HTTPError "Failed to fetch tasks"
// @Router /tasks [get]
func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.service.GetAllTasks()
	if err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, "Failed to fetch tasks", err)
		return
	}

	if len(tasks) == 0 {
		utils.RespondSuccess(ctx, http.StatusOK, "No tasks found", tasks)
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, "Tasks fetched successfully", tasks)
}

// @Summary Create Tasks
// @Description
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body dto.TaskRequest true "Task Request Body"
// @Success 201 {object} utils.APIResponse{status=string,message=string,data=model.Task} "Task created successfully"
// @Failure 400 {object} utils.APIResponse{status=string,message=string} "Invalid input"
// @Failure 500 {object} utils.APIResponse{status=string,message=string} "Failed to create task"
// @Router /tasks [post]
func (c *TaskController) CreateTask(ctx *gin.Context) {
	var params dto.TaskRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	createdTask, err := c.service.CreateTask(params)
	if err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, "Failed to create task", err)
		return
	}

	utils.RespondSuccess(ctx, http.StatusCreated, "Task created successfully", createdTask)
}
