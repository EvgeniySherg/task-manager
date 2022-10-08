package config

import (
	"ToDoList/internal/postgres"
	"time"
)

type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DBPostgres   postgres.PostgresConfig
}

func InitConfig() *Config {
	cfg := &Config{
		Port:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		DBPostgres: postgres.PostgresConfig{
			Host:         "localhost",
			Port:         "5432",
			User:         "postgres",
			Password:     "admin",
			Sslmode:      "disable",
			DatabaseName: "postgres",
		},
	}
	return cfg
}
