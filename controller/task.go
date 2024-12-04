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
// @Description Sample Payload: <br> `{ `<br>` "title": "Fix Bugs", `<br>` "description": "fix multiple bugs in dev", `<br>` "startDate": "2024-12-05T00:00:00Z", `<br>` "deadline": "2024-12-07T00:00:00Z" `<br>` }`
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body dto.TaskRequest true "Task Request Body"
// @Success 201 {object} utils.APIResponse{status=string,message=string,data=model.Task} "Task created successfully"
// @Failure 400 {object} utils.APIResponse{status=string,message=string} "Invalid input"
// @Failure 500 {object} utils.HTTPError "Failed to create task"
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

// @Summary Get Task by Id
// @Description
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "Task Id"
// @Success 200 {object} utils.APIResponse{data=model.Task} "Task fetched successfully"
// @Failure 404 {object} utils.APIResponse{status=string,message=string} "Task not found"
// @Failure 500 {object} utils.HTTPError "Failed to fetch task"
// @Router /tasks/{id} [get]
func (c *TaskController) GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := c.service.GetTaskById(id)
	if err != nil {
		if err.Error() == "task not found" {
			utils.RespondError(ctx, http.StatusNotFound, "Task not found", err)
		} else {
			utils.RespondError(ctx, http.StatusInternalServerError, "Failed to fetch task", err)
		}
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, "Task fetched successfully", task)
}

// @Summary Update Task
// @Description Sample Payload (only send the ones you want to update): <br> `{ `<br>` "description": "fix lots of bugs", `<br>` "completed": true, `<br>` "startDate": "2024-12-05T00:00:00Z", `<br>` "deadline": "2024-12-07T00:00:00Z" `<br>` }`
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "Task Id"
// @Param taskUpdate body dto.TaskUpdate true "Updated Task"
// @Success 200 {object} model.Task "Task updated successfully"
// @Failure 400 {object} utils.APIResponse{status=string,message=string} "Invalid input"
// @Failure 404 {object} utils.APIResponse{status=string,message=string} "Task not found"
// @Router /tasks/{id} [patch]
func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var taskUpdate dto.TaskUpdate
	if err := ctx.ShouldBindJSON(&taskUpdate); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	updatedTask, err := c.service.UpdateTask(id, taskUpdate)
	if err != nil {
		utils.RespondError(ctx, http.StatusNotFound, "Task not found", err)
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, "Task updated successfully", updatedTask)
}

// @Summary Delete Task
// @Description
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task Id"
// @Success 200 {object} utils.APIResponse "Task deleted successfully"
// @Failure 400 {object} utils.APIResponse{status=string,message=string} "Invalid request"
// @Failure 404 {object} utils.APIResponse{status=string,message=string} "Task not found"
// @Router /tasks/{id} [delete]
func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.DeleteTask(id)
	if err != nil {
		utils.RespondError(ctx, http.StatusNotFound, "Task not found", err)
		return
	}

	// Respond with a success message if deletion is successful
	utils.RespondSuccess(ctx, http.StatusOK, "Task deleted successfully", nil)
}
