package postgres

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

type PostgresConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	Sslmode      string
	DatabaseName string
}

func InitDB(cnf *PostgresConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cnf.Host, cnf.Port, cnf.User, cnf.Password, cnf.DatabaseName, cnf.Sslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logrus.Printf("connection to database not created")
		return nil, err
	}
	logrus.Printf("connection to database create successfully")
	return db, nil
}
