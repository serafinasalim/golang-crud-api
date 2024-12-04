package repository

import (
	"golang-crud-api/model"
)

var tasks []model.Task

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) GetAllTasks() ([]model.Task, error) {
	if len(tasks) == 0 {
		return []model.Task{}, nil
	}
	return tasks, nil
}
