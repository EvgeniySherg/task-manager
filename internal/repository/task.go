package repository

import (
	"ToDoList/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type taskDatabase struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) models.TaskRepository {
	return &taskDatabase{
		db: db,
	}
}

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskNotCreated = errors.New("task not created")
var ErrInvalidUser = errors.New("this task belongs to another user")

func (db *taskDatabase) GetById(ctx context.Context, task *models.Task) (*models.Task, error) {
	var newTask models.Task
	query := fmt.Sprint("SELECT name, description, status, created_at, update_at FROM task WHERE id=$1 ")

	err := db.db.QueryRowContext(ctx, query, task.Id).Scan(&newTask.Name, &newTask.Description, &newTask.Status, &newTask.Date, &newTask.LastUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}
	return task, nil
}

func (db *taskDatabase) GetAllTask(ctx context.Context, respTask *models.Task) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)

	query := "SELECT name, description, status, created_at, update_at FROM task WHERE owner_id=$1"

	rows, err := db.db.QueryContext(ctx, query, respTask.OwnerID)
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

func (db *taskDatabase) GetTaskFilterByDate(ctx context.Context, respTask *models.Task) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)
	log.Println(respTask.LastUpdate)
	t, err := time.Parse("2006-01-02", respTask.LastUpdate)
	if err != nil {
		return nil, err
	}

	query := "SELECT name, description, status, created_at, update_at FROM task WHERE created_at > $1 AND owner_id = $2"

	rows, err := db.db.QueryContext(ctx, query, t)
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

func (db *taskDatabase) CreateTask(ctx context.Context, task *models.Task) error {

	query := fmt.Sprint("INSERT INTO task (name, description, status, owner_id) VALUES ($1, $2, $3, $4)  RETURNING id, name")

	var createdTask models.Task

	err := db.db.QueryRowContext(ctx, query, task.Name, task.Description, task.Status, task.OwnerID).Scan(&createdTask.Id, &createdTask.Name)

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

func (db *taskDatabase) UpdateTask(ctx context.Context, task *models.Task) error {
	query := fmt.Sprintf("SELECT owner_id FROM task WHERE id=$1")

	var examTask models.Task

	err := db.db.QueryRowContext(ctx, query, task.Id).Scan(&examTask.OwnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrTaskNotFound
		}
		return err
	}

	if task.OwnerID != examTask.OwnerID {
		return ErrInvalidUser
	}

	changeTime := time.Now()

	updateQuery := fmt.Sprint("UPDATE task SET name = $1, description = $2, status = $3, update_at = $4 WHERE id = $5;")

	res, err := db.db.ExecContext(ctx, updateQuery, task.Name, task.Description, task.Status, changeTime, task.Id)
	if err != nil {
		return fmt.Errorf("exec err -> %v", err)
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New(fmt.Sprintf("-> task %v not updated. No rows change ", task.Name))
	}
	return nil
}

func (db *taskDatabase) DeleteTask(ctx context.Context, task *models.Task) error {
	query := fmt.Sprintf("SELECT owner_id FROM task WHERE id=$1")

	var examTask models.Task

	err := db.db.QueryRowContext(ctx, query, task.Id).Scan(&examTask.OwnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrTaskNotFound
		}
		return err
	}

	if task.OwnerID != examTask.OwnerID {
		return ErrInvalidUser
	}

	deleteQuery := fmt.Sprint(`DELETE FROM task WHERE id = $1`)

	res, err := db.db.ExecContext(ctx, deleteQuery, task.Id)

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
