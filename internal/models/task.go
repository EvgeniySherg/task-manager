package models

import "context"

type Task struct {
	Id          int    `json:"Id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Status      string `json:"status"`
	LastUpdate  string `json:"lastUpdate"`
	User        User   `json:"user"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type TaskRepository interface {
	GetById(ctx context.Context, id int) (*Task, error)
	GetAllTask(ctx context.Context, userId int) ([]*Task, error)
	GetTaskFilterByDate(ctx context.Context, date string) ([]*Task, error)
	CreateTask(ctx context.Context, task *Task) error
	UpdateTask(ctx context.Context, task *Task) error
	DeleteTask(ctx context.Context, id int) error
}
