package handlers

import (
	"ToDoList/internal/models"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	salt       = "qwertyuiop" //
	signingKey = "asdfghjkl"  //
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

// регистрация пользователя с хешированием пароля
func (uh *userHandler) SignUp(c echo.Context) error {
	var user *models.User

	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user.Password = uh.generatePasswordHash(user.Password)

	err = uh.repository.CreateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "user create successfully")
}

// SignIn аутентификация пользователя по Header запроса
// функция возвращает токен, который потом можно использовать вручную добавив Header c заголовком Authorization
// при наличии такого заголовка запрос GetTaskById выдаст нужный результат
func (uh *userHandler) SignIn(c echo.Context) error {
	var signUser models.User

	err := json.NewDecoder(c.Request().Body).Decode(&signUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := uh.repository.GetUser(c.Request().Context(), signUser.Name, uh.generatePasswordHash(signUser.Password))

	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error()) //
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, user.Id})

	myToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"token": myToken})
}

func (uh *userHandler) generatePasswordHash(pass string) string {
	hash := sha1.New()

	hash.Write([]byte(pass))

	hashCode := fmt.Sprintf("%v", hash.Sum([]byte(salt)))
	return base64.StdEncoding.EncodeToString([]byte(hashCode))
}
