package repository

import (
	"ToDoList/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type database struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) models.TaskRepository {
	return &database{
		DB: db,
	}
}

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskNotCreated = errors.New("task not created")

func (db *database) GetById(ctx context.Context, id int) (*models.Task, error) {
	var task models.Task
	query := fmt.Sprint("SELECT task_name, task_description, status, created_at, update_at FROM task WHERE id=$1 ")

	err := db.DB.QueryRowContext(ctx, query, id).Scan(&task.Name, &task.Description, &task.Status, &task.Date, &task.LastUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}
	return &task, nil
}

func (db *database) GetAllTask(ctx context.Context, userId int) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)

	query := "SELECT task_name, task_description, status, created_at, update_at FROM task WHERE owner_id=$1"

	rows, err := db.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		task := models.Task{}
		err := rows.Scan(&task.Name, &task.Description, &task.Status, &task.Date, &task.LastUpdate)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (db *database) GetTaskFilterByDate(ctx context.Context, date string) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	query := "SELECT task_name, task_description, status, created_at, update_at FROM task WHERE created_at > $1"

	rows, err := db.DB.QueryContext(ctx, query, t)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		task := models.Task{}
		err := rows.Scan(&task.Name, &task.Description, &task.Status, &task.Date, &task.LastUpdate)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (db *database) CreateTask(ctx context.Context, task *models.Task) error {

	query := fmt.Sprint("INSERT INTO task (task_name, task_description, status, owner_id) VALUES ($1, $2, $3, $4)  RETURNING id, task_name")

	var createdTask models.Task

	err := db.DB.QueryRowContext(ctx, query, task.Name, task.Description, task.Status, task.OwnerID).Scan(&createdTask.Id, &createdTask.Name)

	switch {
	case err == sql.ErrNoRows:
		return ErrTaskNotCreated
	case err != nil:
		return fmt.Errorf("create book err -> %v", err)
	default:
		logrus.Printf("task with id - %v  created,  title - %s\n", createdTask.Id, createdTask.Name)
	}
	return nil
}

func (db *database) UpdateTask(ctx context.Context, task *models.Task) error {
	changeTime := time.Now()

	query := fmt.Sprint("UPDATE task SET task_name = $1, task_description = $2, status = $3, update_at = $4 WHERE id = $5;")

	res, err := db.DB.ExecContext(ctx, query, task.Name, task.Description, task.Status, changeTime, task.Id)
	if err != nil {
		return fmt.Errorf("exec err -> %v", err)
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New(fmt.Sprintf("-> task %v not updated. No rows change ", task.Name))
	}
	return nil
}

func (db *database) DeleteTask(ctx context.Context, id int) error {
	query := fmt.Sprint(`DELETE FROM task WHERE id = $1`)

	res, err := db.DB.ExecContext(ctx, query, id)

	if err != nil {
		return fmt.Errorf("exec err -> %v", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected while delete err -> %v", err)
	}
	if rows == 0 {
		return ErrTaskNotFound
	}

	return nil
}
