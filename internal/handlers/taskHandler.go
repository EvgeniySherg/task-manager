package handlers

import (
	"ToDoList/internal/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (th *taskHandler) GetTaskById(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Printf("cannot strconv.Atoi: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect ID num for get task")
	}

	task, err := th.repository.GetById(c.Request().Context(), ID)
	if err != nil {
		logrus.Printf("Get task by ID error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "database error, incorrect id")
	}

	return c.JSON(http.StatusOK, json.NewEncoder(c.Response()).Encode(task))
}

func (th *taskHandler) GetAllTasksByUserId(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		logrus.Printf("cannot strconv.Atoi: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect user ID num, for get tasks")
	}

	tasks, err := th.repository.GetAllTask(c.Request().Context(), userID)
	if err != nil {
		logrus.Printf("Get tasks by user ID error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "database error, incorrect user id")
	}

	return c.JSON(http.StatusOK, json.NewEncoder(c.Response()).Encode(tasks))
}

func (th *taskHandler) GetTasksFilterByDate(c echo.Context) error {
	date := c.Param("date")

	tasks, err := th.repository.GetTaskFilterByDate(c.Request().Context(), date)
	if err != nil {
		logrus.Printf("Get tasks filtered by date error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "database error, incorrect date information")
	}

	return c.JSON(http.StatusOK, json.NewEncoder(c.Response()).Encode(tasks))
}

func (th *taskHandler) CreateTask(c echo.Context) error {

	var newTask models.Task

	err := json.NewDecoder(c.Request().Body).Decode(&newTask)
	if err != nil {
		logrus.Println("entered incorrect data for create task")
		return c.String(http.StatusBadRequest, "incorrect task data")
	}

	err = th.repository.CreateTask(c.Request().Context(), &newTask)
	if err != nil {
		logrus.Printf("error while create new task %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "task create successfully")
}

func (th *taskHandler) UpdateTask(c echo.Context) error {
	var updateTask models.Task

	err := json.NewDecoder(c.Request().Body).Decode(&updateTask)
	if err != nil {
		logrus.Println("entered incorrect data for updating task")
		return c.String(http.StatusBadRequest, "incorrect task data")
	}

	err = th.repository.UpdateTask(c.Request().Context(), &updateTask)
	if err != nil {
		logrus.Printf("error while update task %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "task update successfully")
}

func (th *taskHandler) DeleteTask(c echo.Context) error {
	deleteID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Printf("cannot strconv.Atoi: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect ID num for delete task")
	}

	err = th.repository.DeleteTask(c.Request().Context(), deleteID)
	if err != nil {
		logrus.Printf("delete task error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "database error, incorrect id")
	}

	return c.JSON(http.StatusOK, "task delete")
}
