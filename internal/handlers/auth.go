package handlers

import (
	"ToDoList/internal/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (uh UserHandler) SignUp(c echo.Context) {
	var user models.User

	err := json.NewDecoder(c.Request().Body()).Decode(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

}

func (uh UserHandler) SignIn(c echo.Context) {

}
