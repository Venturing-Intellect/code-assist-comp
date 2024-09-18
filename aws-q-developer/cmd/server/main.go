package main

import (
	"ai-compare-app/configs"
	"ai-compare-app/internal/controllers"
	"ai-compare-app/internal/repositories"
	"ai-compare-app/internal/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := sql.Open("postgres", config.GetDatabaseURL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	repo := repositories.NewPostgresFeedbackRepository(db)
	service := services.NewFeedbackService(repo)
	controller := controllers.NewFeedbackController(service)

	r := chi.NewRouter()
	r.Post("/feedback", controller.CreateFeedback)

	log.Printf("Starting server on port %s", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
