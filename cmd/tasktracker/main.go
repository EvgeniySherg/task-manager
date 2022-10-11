package main

import (
	"ToDoList/internal/config"
	"ToDoList/internal/handlers"
	"ToDoList/internal/postgres"
	"ToDoList/internal/repository"
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func initHandlers(app *echo.Echo, db *sql.DB) {
	userRep := repository.NewUserRepository(db)
	userHandl := handlers.NewUserHandler(userRep)

	taskRep := repository.NewRepository(db)
	taskHandl := handlers.NewTaskHandler(taskRep)

	//middleware версия запроса для получения конкретной таски. Попытка сделать утентификацию по JWT токену
	app.POST("/auth/sign-up", userHandl.SignUp) // регистрация с присвоением JWT токена
	app.POST("/auth/sign-in", userHandl.SignIn)
	app.GET("/taskJWT", taskHandl.GetTaskById, userHandl.UserIdentity)

	app.GET("/taskById", taskHandl.GetTaskById)
	app.GET("/tasksByDate", taskHandl.GetTasksFilterByDate)
	app.GET("/tasksByUser", taskHandl.GetAllTasksByUserId)
	app.POST("/taskCreate", taskHandl.CreateTask)
	app.PUT("/taskUpdate", taskHandl.UpdateTask)
	app.DELETE("/taskDelete", taskHandl.DeleteTask)

	app.POST("/userCreate", userHandl.RegistrationNewUser)

}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := config.InitConfigFile(); err != nil {
		logrus.Fatalf("failed init config: %v", err)
	}

	cnf := config.InitConfig()

	db, err := postgres.InitDB(&cnf.DBPostgres)
	if err != nil {
		logrus.Fatalf("failed init config: %v", err)
	}

	app := echo.New()

	initHandlers(app, db)

	httpServer := &http.Server{
		Addr:         cnf.Port,
		Handler:      app,
		ReadTimeout:  cnf.ReadTimeout,
		WriteTimeout: cnf.WriteTimeout,
	}

	go func() {
		if err := app.StartServer(httpServer); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("failed start http server -> %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err = db.Close(); err != nil {
		logrus.Printf("db close failed: %v", err)
	}

	shutdownCtx, forceShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer forceShutdown()

	if err = httpServer.Shutdown(shutdownCtx); err != nil {
		logrus.Fatalf("shutdown http server err: %v", err)
	}
}
