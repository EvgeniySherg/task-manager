package models

import "context"

type Task struct {
	TaskId      int    `json:"taskId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Status      string `json:"status"`
	LastUpdate  string `json:"lastUpdate"`
	TaskUser    User   `json:"taskUser"`
}

type User struct {
	UserId       int    `json:"userId"`
	UserName     string `json:"userName"`
	UserPassword string `json:"userPassword"`
}

type TaskRepository interface {
	GetById(ctx context.Context, id int) (*Task, error)
	GetAllTask(ctx context.Context, userId int) ([]*Task, error)
	GetTaskFilterByDate(ctx context.Context, date string) ([]*Task, error)
	CreateTask(ctx context.Context, task *Task) error
	UpdateTask(ctx context.Context, task *Task) error
	DeleteTask(ctx context.Context, id int) error
}
