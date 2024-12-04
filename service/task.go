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
		Completed:   params.Completed,
		StartDate:   params.StartDate,
		Deadline:    params.Deadline,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.repository.CreateTask(&task)
}
