package models

import (
	"context"
)

type Task struct {
	Id          int    `json:"Id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Status      string `json:"status"`
	LastUpdate  string `json:"lastUpdate"`
	OwnerID     int    `json:"owner_id"`
}

type TaskRepository interface {
	GetById(ctx context.Context, task *Task) (*Task, error)
	GetAllTask(ctx context.Context, task *Task) ([]*Task, error)
	GetTaskFilterByDate(ctx context.Context, task *Task) ([]*Task, error)
	CreateTask(ctx context.Context, task *Task) error
	UpdateTask(ctx context.Context, task *Task) error
	DeleteTask(ctx context.Context, task *Task) error
}
