package service

import (
	"golang-crud-api/dto"
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

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	tasks, err := s.repository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return nil, nil
	}
	return tasks, nil
}

func (s *TaskService) CreateTask(params dto.TaskRequest) (*model.Task, error) {
	task := model.Task{
		Title:       params.Title,
		Description: params.Description,
		Completed:   false,
		StartDate:   params.StartDate,
		Deadline:    params.Deadline,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.repository.CreateTask(&task)
}

func (s *TaskService) GetTaskById(id string) (model.Task, error) {
	task, err := s.repository.GetTaskById(id)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (s *TaskService) UpdateTask(id string, taskUpdate dto.TaskUpdate) (*model.Task, error) {
	// Find the task by ID
	task, err := s.repository.GetTaskById(id)
	if err != nil {
		return nil, err
	}

	// Apply updates to the task
	if taskUpdate.Description != "" {
		task.Description = taskUpdate.Description
	}
	if taskUpdate.Completed != task.Completed {
		task.Completed = taskUpdate.Completed
	}
	if !taskUpdate.StartDate.IsZero() {
		task.StartDate = taskUpdate.StartDate
	}
	if !taskUpdate.Deadline.IsZero() {
		task.Deadline = taskUpdate.Deadline
	}

	task.UpdatedAt = time.Now()

	updatedTask, err := s.repository.UpdateTask(id, &task)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (s *TaskService) DeleteTask(id string) error {
	return s.repository.DeleteTask(id)
}
