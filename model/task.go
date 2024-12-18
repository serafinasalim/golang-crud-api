package model

import "time"

type Task struct {
	Id          int        `json:"id"`
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

type TaskUpdate struct {
	Id          int        `json:"id"`
	Uuid        string     `json:"uuid"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   *bool      `json:"completed"`
	StartDate   time.Time  `json:"startDate"`
	Deadline    time.Time  `json:"deadline"`
	CreatedAt   time.Time  `json:"createdAt"`
	CreatedBy   string     `json:"createdBy"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	UpdatedBy   *string    `json:"updatedBy"`
}
