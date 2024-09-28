package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Pertsaa/pokkne/internal/handler"
	"github.com/Pertsaa/pokkne/internal/middleware"
	"github.com/Pertsaa/pokkne/internal/sql"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "host=pokkne-db port=5432 user=postgresdev password=postgresdev dbname=postgresdev sslmode=disable")
	if err != nil {
		log.Fatal("unable to create connection pool: %w", err)
	}
	defer pool.Close()

	sql.MustMigrate(pool)

	r := http.NewServeMux()

	h := handler.NewAPIHandler(ctx, pool)

	// Words
	r.HandleFunc("POST /v1/sentences", handler.Make(h.SentenceCreateHandler))
	r.HandleFunc("POST /v1/sentences/batch", handler.Make(h.SentenceBatchCreateHandler))
	r.HandleFunc("GET /v1/sentences", handler.Make(h.SentenceListHandler))
	r.HandleFunc("POST /v1/seed", handler.Make(h.SeedHandler))

	// Generate
	r.HandleFunc("POST /v1/chat", handler.Make(h.ChatHandler))

	stack := middleware.CreateStack(
		middleware.Log,
		middleware.CORS,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(r),
	}

	fmt.Println("API listening on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
