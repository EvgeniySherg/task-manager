package models

import (
	"context"
	"errors"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserRepository interface {
	GetUser(ctx context.Context, username, pass string) (*User, error)
	CreateUser(c context.Context, user *User) error
}

var ErrUnknownAccess = errors.New("access level not found")
