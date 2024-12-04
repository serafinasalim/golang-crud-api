package service

import (
	"golang-crud-api/model"
	"golang-crud-api/repository"
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
