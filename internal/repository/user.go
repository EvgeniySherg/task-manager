package repository

import (
	"ToDoList/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type userDatabase struct {
	db *sql.DB
}

var ErrUserNotFound = errors.New("user not found")
var ErrUserNotCreated = errors.New("user not created")

func NewUserRepository(db *sql.DB) models.UserRepository {
	return &userDatabase{
		db: db,
	}
}

func (ud *userDatabase) CreateUser(ctx context.Context, user *models.User) error {

	query := fmt.Sprint("INSERT INTO users (name, pass) VALUES ($1, $2) RETURNING id")

	var createUser models.User

	err := ud.db.QueryRowContext(ctx, query, user.Name, user.Password).Scan(&createUser.Id)

	switch {
	case err == sql.ErrNoRows:
		return ErrUserNotCreated
	case err != nil:
		return fmt.Errorf("create user err -> %v", err)
	default:
		logrus.Printf("user with id - %v  created\n", createUser.Id)
	}
	return nil
}

func (ud *userDatabase) GetUser(ctx context.Context, username string, pass string) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id, name, pass  FROM users WHERE name = $1 AND pass = $2")

	err := ud.db.QueryRowContext(ctx, query, username, pass).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
