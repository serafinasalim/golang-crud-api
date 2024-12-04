package repository

import (
	"errors"
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

func (r *TaskRepository) GetTaskById(id string) (model.Task, error) {
	for _, task := range tasks {
		if task.Id == id {
			return task, nil
		}
	}
	return model.Task{}, errors.New("task not found")
}

func (r *TaskRepository) UpdateTask(id string, updatedTask *model.Task) (*model.Task, error) {
	for i, task := range tasks {
		if task.Id == id {
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Completed != tasks[i].Completed {
				tasks[i].Completed = updatedTask.Completed
			}
			if !updatedTask.StartDate.IsZero() {
				tasks[i].StartDate = updatedTask.StartDate
			}
			if !updatedTask.Deadline.IsZero() {
				tasks[i].Deadline = updatedTask.Deadline
			}

			tasks[i].UpdatedAt = updatedTask.UpdatedAt
			return &tasks[i], nil
		}
	}
	return nil, fmt.Errorf("task with Id %s not found", id)
}

func (r *TaskRepository) DeleteTask(id string) error {
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with Id %s not found", id)
}
