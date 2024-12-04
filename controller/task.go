package controller

import (
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
