package models

import "time"

type Status string

const (
	Todo       = "Todo"
	InProgress = "in-progress"
	Done       = "done"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
