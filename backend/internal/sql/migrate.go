package sql

import (
	"embed"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrations embed.FS

func MustMigrate(pool *pgxpool.Pool) {
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	db := stdlib.OpenDBFromPool(pool)
	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	if err := db.Close(); err != nil {
		panic(err)
	}
}
