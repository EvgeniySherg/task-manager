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
	salt       = "qwertyuiop"
	signingKey = "asdfghjkl"
)

func (uh *userHandler) SignUp(c echo.Context) error {
	var user *models.User

	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user.Password = generatePasswordHash(user.Password)

	err = uh.repository.CreateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "user create successfully")
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

func (uh *userHandler) SignIn(c echo.Context) error {
	var sign signInInput

	err := json.NewDecoder(c.Request().Body).Decode(&sign)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := uh.repository.GetUser(c.Request().Context(), sign.Username, generatePasswordHash(sign.Password))

	if err != nil {
		return c.JSON(http.StatusNonAuthoritativeInfo, err.Error()) //TODO найти код ошибки
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

func generatePasswordHash(pass string) string {
	hash := sha1.New()

	hash.Write([]byte(pass))

	hashCode := fmt.Sprintf("%v", hash.Sum([]byte(salt)))
	return base64.StdEncoding.EncodeToString([]byte(hashCode))
}
