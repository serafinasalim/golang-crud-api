package dto

import "time"

type TaskRequest struct {
	Title       string    `json:"title" binding:"required,max=100"`
	Description string    `json:"description" binding:"required,min=5"`
	StartDate   time.Time `json:"startDate" binding:"required"`
	Deadline    time.Time `json:"deadline" binding:"required"`
}
