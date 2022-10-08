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
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func initHandlers(app *echo.Echo, db *sql.DB) {
	taskRep := repository.NewRepository(db)
	taskHandl := handlers.NewHandler(taskRep)
	app.GET("/task/:id", taskHandl.GetTaskById)
	app.GET("/tasks/:date", taskHandl.GetTasksFilterByDate)
	app.GET("/tasksByUser/:id", taskHandl.GetAllTasksByUserId)
	app.POST("/taskCreate", taskHandl.CreateTask)
	app.PUT("/taskUpdate", taskHandl.UpdateTask)
	app.DELETE("/taskDelete/:id", taskHandl.DeleteTask)
}

func main() {
	cnf := config.InitConfig()

	db, err := postgres.InitDB(&cnf.DBPostgres)
	if err != nil {
		log.Fatalf("failed init config: %v", err)
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
			log.Fatalf("failed start http server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err = db.Close(); err != nil {
		log.Printf("db close failed: %v", err)
	}

	shutdownCtx, forceShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer forceShutdown()

	if err = httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("shutdown http server err: %v", err)
	}
}
