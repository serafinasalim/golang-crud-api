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

	var res []dto.TaskResponse

	// Convert to Dto to avoid security risks by not returning id (guessable)
	for _, task := range tasks {
		res = append(res, helper.ConvertTaskToDto(task))
	}

	return res, nil
}

func (s *TaskService) CreateTask(params dto.TaskRequest) (result *string, err error) {
	task := model.Task{
		Title:       params.Title,
		Description: params.Description,
		Completed:   false,
		StartDate:   params.StartDate,
		Deadline:    params.Deadline,
		CreatedAt:   time.Now(),
		CreatedBy:   "placeholder",
	}
	// placeholder before jwt authentication integration

	res, err := s.repository.CreateTask(task)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TaskService) GetTaskByUuid(uuid string) (result *dto.TaskResponse, err error) {
	task, err := s.repository.GetTaskByUuid(uuid)
	if err != nil {
		return nil, err
	}

	res := &dto.TaskResponse{
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

	return res, nil
}

func (s *TaskService) UpdateTask(uuid string, params dto.TaskUpdate) (err error) {
	// Find the task by Uuid
	_, err = s.repository.GetTaskByUuid(uuid)
	if err != nil {
		return err
	}

	updateValue := &model.Task{
		Uuid:        uuid,
		Description: params.Description,
		Completed:   params.Completed,
		StartDate:   params.StartDate,
		Deadline:    params.Deadline,
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
