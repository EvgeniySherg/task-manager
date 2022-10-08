package handlers

import "ToDoList/internal/models"

type TaskHandler struct {
	repository models.TaskRepository
}

func NewHandler(rep models.TaskRepository) TaskHandler {
	return TaskHandler{repository: rep}
}
