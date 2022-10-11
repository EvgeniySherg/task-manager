package config

import (
	"ToDoList/internal/postgres"
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
	//if err := godotenv.Load(); err != nil {    // <---- для использования в докер образе пришлось закоментить
	//	logrus.Fatal(err)						// <--- расскоментить для запуска из IDE
	//}
	cfg := &Config{
		Port:         viper.GetString("port"),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		DBPostgres: postgres.PostgresConfig{
			Host:         os.Getenv("DB_HOST"),     //viper.GetString("db.host"),
			Port:         os.Getenv("DB_PORT"),     //viper.GetString("db.port"),
			User:         os.Getenv("DB_USER"),     //viper.GetString("db.user"),
			Password:     os.Getenv("DB_PASSWORD"), //  viper.GetString("db.password")
			Sslmode:      viper.GetString("db.sslmode"),
			DatabaseName: os.Getenv("DB_NAME"), //viper.GetString("db.databaseName"),
		},
	}
	return cfg
}
