package handlers

import (
	"ToDoList/internal/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (uh *userHandler) RegistrationNewUser(c echo.Context) error {
	var newUser models.User

	err := json.NewDecoder(c.Request().Body).Decode(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = uh.repository.CreateUser(c.Request().Context(), &newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "user create successfully")
}
