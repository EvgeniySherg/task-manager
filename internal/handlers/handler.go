package handlers

import "ToDoList/internal/models"

type TaskHandler struct {
	repository models.TaskRepository
}

func NewHandler(rep models.TaskRepository) TaskHandler {
	return TaskHandler{repository: rep}
}

type UserHandler struct {
	repository models.UserRepository
}

func NewUserHandler(rep models.UserRepository) UserHandler {
	return UserHandler{repository: rep}
}
