package repository

import (
	"database/sql"
	"golang-crud-api/model"
	"log"

	_ "github.com/lib/pq"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	if db == nil {
		log.Fatal("Received nil database connection")
	}

	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAllTasks() (result []model.Task, err error) {
	var task model.Task
	const query = `SELECT 
						id, 
						uuid, 
						title, 
						description, 
						completed, 
						start_date, 
						deadline, 
						created_at, 
						created_by, 
						updated_at, 
						updated_by 
					FROM public.tasks`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&task.Id,
			&task.Uuid,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.StartDate,
			&task.Deadline,
			&task.CreatedAt,
			&task.CreatedBy,
			&task.UpdatedAt,
			&task.UpdatedBy,
		); err != nil {
			return nil, err
		}
		result = append(result, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *TaskRepository) CreateTask(params model.Task) (result *string, err error) {
	const query = `INSERT INTO public.tasks (
						title, 
						description, 
						completed, 
						start_date, 
						deadline, 
						created_at,
						created_by
					)
					VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = r.db.Exec(
		query,
		params.Title,
		params.Description,
		params.Completed,
		params.StartDate,
		params.Deadline,
		params.CreatedAt,
		params.CreatedBy,
	)
	if err != nil {
		return nil, err
	}
	result = &params.Title

	return result, nil
}

func (r *TaskRepository) GetTaskByUuid(uuid string) (result *model.Task, err error) {
	var task model.Task
	const query = `SELECT 
						id, 
						uuid, 
						title, 
						description, 
						completed, 
						start_date, 
						deadline, 
						created_at, 
						created_by, 
						updated_at, 
						updated_by 
					FROM public.tasks 
					WHERE uuid = $1`

	err = r.db.QueryRow(query, uuid).Scan(
		&task.Id,
		&task.Uuid,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.StartDate,
		&task.Deadline,
		&task.CreatedAt,
		&task.CreatedBy,
		&task.UpdatedAt,
		&task.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) UpdateTask(params *model.Task) (err error) {
	// Query untuk memperbarui task berdasarkan Uuid
	const query = `UPDATE public.tasks SET
						description = COALESCE($2, description),
						completed = COALESCE($3, completed),
						start_date = COALESCE($4, start_date),
						deadline = COALESCE($5, deadline),
						updated_at = $6,
						updated_by = $7
					WHERE uuid = $1`

	_, err = r.db.Exec(
		query,
		params.Uuid,
		params.Description,
		params.Completed,
		params.StartDate,
		params.Deadline,
		params.UpdatedAt,
		params.UpdatedBy,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) DeleteTask(uuid string) (err error) {
	const query = `DELETE FROM public.tasks WHERE uuid = $1`

	result, err := r.db.Exec(query, uuid)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}
