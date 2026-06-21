package main

import (
	"database/sql"
	"log"
	"start-with-go/internal/config"
	"start-with-go/internal/db"
	"start-with-go/internal/handler"
	"start-with-go/internal/middleware"
	"start-with-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	sqlDB, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer sqlDB.Close()

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	} else {
		log.Println("Successfully connected to the database")
	}

	queries := db.New(sqlDB)
	repo := repository.NewNoteRepository(queries)
	noteHandler := handler.NewNoteHandler(repo)

	r := gin.Default()

	api := r.Group("/api/v1")
	api.Use(middleware.APIKeyAuth(cfg.APIKey))
	{
		notes := api.Group("/notes")
		notes.GET("", noteHandler.List)
		notes.POST("", noteHandler.Create)
		notes.GET("/:id", noteHandler.Get)
		notes.PUT("/:id", noteHandler.Update)
		notes.DELETE("/:id", noteHandler.Delete)
	}

	log.Printf("server listening on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
