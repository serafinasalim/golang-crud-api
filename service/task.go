package service

import (
	"golang-crud-api/dto"
	"golang-crud-api/helper"
	"golang-crud-api/model"
	"golang-crud-api/repository"
	"time"
)

type TaskService struct {
	repository *repository.TaskRepository
}

func NewTaskService(repository *repository.TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (s *TaskService) GetAllTasks() (result []dto.TaskResponse, err error) {
	tasks, err := s.repository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return nil, nil
	}

	// Convert to Dto to avoid security risks by not returning id (guessable)
	for _, task := range tasks {
		result = append(result, helper.ConvertTaskToDto(task))
	}

	return result, nil
}

func (s *TaskService) CreateTask(params dto.TaskRequest) (result *string, err error) {
	task := model.Task{
		Title:       params.Title,
		Description: params.Description,
		Completed:   false,
		StartDate:   params.StartDate,
		Deadline:    params.Deadline,
		CreatedAt:   time.Now(),
		CreatedBy:   params.CreatedBy,
	}

	result, err = s.repository.CreateTask(task)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *TaskService) GetTaskByUuid(uuid string) (result *dto.TaskResponse, err error) {
	task, err := s.repository.GetTaskByUuid(uuid)
	if err != nil {
		return nil, err
	}

	result = &dto.TaskResponse{
		// Not returning id to avoid security risks (guessable)
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

	return result, nil
}

func (s *TaskService) UpdateTask(uuid string, params dto.TaskUpdate) (err error) {
	// Find the task by Uuid
	details, err := s.repository.GetTaskByUuid(uuid)
	if err != nil {
		return err
	}

	now := time.Now()
	updateValue := &model.TaskUpdate{
		Uuid:        uuid,
		Description: helper.CoalesceString(params.Description, details.Description),
		Completed:   helper.CoalesceBoolPtr(params.Completed, &details.Completed),
		StartDate:   helper.CoalesceTime(params.StartDate, details.StartDate),
		Deadline:    helper.CoalesceTime(params.Deadline, details.Deadline),
		UpdatedAt:   &now,
		UpdatedBy:   params.UpdatedBy,
	}

	err = s.repository.UpdateTask(updateValue)
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) DeleteTask(uuid string) (err error) {
	// Find the task by Uuid
	_, err = s.repository.GetTaskByUuid(uuid)
	if err != nil {
		return err
	}

	err = s.repository.DeleteTask(uuid)
	if err != nil {
		return err
	}

	return nil
}
