package config

import (
	"ToDoList/internal/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DBPostgres   postgres.PostgresConfig
}

func InitConfigFile() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func InitConfig() *Config {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}
	cfg := &Config{
		Port:         viper.GetString("port"),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		DBPostgres: postgres.PostgresConfig{
			Host:         viper.GetString("db.host"),
			Port:         viper.GetString("db.port"),
			User:         viper.GetString("db.user"),
			Password:     os.Getenv("DB_PASSWORD"), //  viper.GetString("db.password")
			Sslmode:      viper.GetString("db.sslmode"),
			DatabaseName: viper.GetString("db.databaseName"),
		},
	}
	return cfg
}
