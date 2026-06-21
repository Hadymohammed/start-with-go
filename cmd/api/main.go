package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"start-with-go/internal/api"
	"start-with-go/internal/config"
	"start-with-go/internal/db"
	"start-with-go/internal/handler"
	"start-with-go/internal/middleware"
	"start-with-go/internal/repository"
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

	handler.RegisterDocsRoutes(r)

	authMiddleware := func(f api.StrictHandlerFunc, operationID string) api.StrictHandlerFunc {
		return func(ctx *gin.Context, request any) (any, error) {
			switch operationID {
			case "NotesCreate", "NotesUpdate", "NotesDelete":
				middleware.APIKeyAuth(cfg.APIKey)(ctx)
				if ctx.IsAborted() {
					return nil, nil
				}
			}
			return f(ctx, request)
		}
	}

	api.RegisterHandlers(r, api.NewStrictHandler(noteHandler, []api.StrictMiddlewareFunc{authMiddleware}))

	log.Printf("server listening on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
