package handlers

import "ToDoList/internal/models"

type taskHandler struct {
	repository models.TaskRepository
}

func NewTaskHandler(rep models.TaskRepository) *taskHandler {
	return &taskHandler{repository: rep}
}

type userHandler struct {
	repository models.UserRepository
}

func NewUserHandler(rep models.UserRepository) *userHandler {
	return &userHandler{repository: rep}
}
