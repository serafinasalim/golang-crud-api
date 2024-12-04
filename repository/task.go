package repository

import (
	"fmt"
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

func (r *TaskRepository) CreateTask(task *model.Task) (*model.Task, error) {
	task.Id = fmt.Sprintf("%d", len(tasks)+1) // Simple ID Generationo
	tasks = append(tasks, *task)
	return task, nil
}
