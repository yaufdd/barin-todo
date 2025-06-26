package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"connectrpc.com/connect"
	_ "github.com/go-sql-driver/mysql"

	todov1connect "github.com/ericbamba/barin-todo/gen/proto/todo/v1/todov1connect"
	"github.com/ericbamba/barin-todo/internal/handler"
	"github.com/ericbamba/barin-todo/internal/middleware"
	"github.com/ericbamba/barin-todo/internal/repo"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "todo"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "todopass"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "todo_db"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println("DSN:", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	repo := &repo.TodoRepo{DB: db}
	h := handler.NewHandler(repo)

	mux := http.NewServeMux()

	path, todoHandler := todov1connect.NewTodoServiceHandler(h, connect.WithInterceptors(middleware.LoggingInterceptor()))
	mux.Handle(path, todoHandler)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("listen :8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
