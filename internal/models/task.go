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

type UserRepository interface {
	GetUser(ctx context.Context, username, pass string) (*User, error)
	CreateUser(c context.Context, user *User) error
}
