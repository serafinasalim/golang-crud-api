package dto

import "time"

type TaskRequest struct {
	Title       string    `json:"title" binding:"required,max=100"`
	Description string    `json:"description" binding:"required,min=5"`
	StartDate   time.Time `json:"startDate" binding:"required"`
	Deadline    time.Time `json:"deadline" binding:"required"`
}

type TaskUpdate struct {
	Description string    `json:"description" binding:"omitempty,min=5"`
	Completed   bool      `json:"completed"`
	StartDate   time.Time `json:"startDate"`
	Deadline    time.Time `json:"deadline"`
}

type TaskResponse struct {
	Uuid        string     `json:"uuid"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	StartDate   time.Time  `json:"startDate"`
	Deadline    time.Time  `json:"deadline"`
	CreatedAt   time.Time  `json:"createdAt"`
	CreatedBy   string     `json:"createdBy"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	UpdatedBy   *string    `json:"updatedBy"`
}
