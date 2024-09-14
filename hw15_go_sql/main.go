package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mussyaroslav/home_work_basic/hw15_go_sql/config"
	"github.com/mussyaroslav/home_work_basic/hw15_go_sql/db"
	"github.com/mussyaroslav/home_work_basic/hw15_go_sql/handlers"
)

func InitDB() (*db.DB, error) {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	database, err := db.NewDB(connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	log.Println("Connected to the database successfully")
	return database, nil
}

func main() {
	database, err := InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer func() {
		if cerr := database.Close(); cerr != nil {
			log.Printf("Error closing the database: %v", cerr)
		}
	}()

	handler := handlers.NewHandler(database)

	http.HandleFunc("/users", handler.UsersHandler)
	http.HandleFunc("/products", handler.ProductsHandler)
	http.HandleFunc("/orders", handler.OrdersHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Starting server on :8080")
	errCh := make(chan error, 1)

	go func() {
		errCh <- server.ListenAndServe()
	}()

	srvErr := <-errCh
	if srvErr != nil && !errors.Is(srvErr, http.ErrServerClosed) {
		log.Printf("Error starting server: %v", srvErr)
	}

	if srvErr != nil {
		log.Printf("Server error: %v", srvErr)
	}
}
