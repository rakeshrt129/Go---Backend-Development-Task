package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"user-age-api/db/sqlc"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewDB creates and returns a database connection
func NewDB(dbURL string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Set basic connection settings (good practice)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	// Check DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// NewQueries returns SQLC Queries instance
func NewQueries(db *sql.DB) *sqlc.Queries {
	return sqlc.New(db)
}
