package helper

import (
	"golang-crud-api/dto"
	"golang-crud-api/model"
)

func ConvertTaskToDto(task model.Task) dto.TaskResponse {
	return dto.TaskResponse{
		Uuid:        task.Uuid,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		StartDate:   task.StartDate,
		Deadline:    task.Deadline,
		CreatedAt:   task.CreatedAt,
		CreatedBy:   task.CreatedBy,
		UpdatedAt:   task.UpdatedAt,
		UpdatedBy:   task.UpdatedBy,
	}
}
