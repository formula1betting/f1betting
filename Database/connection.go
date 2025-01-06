package dbms

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectToDB(ctx context.Context) (*pgx.Conn, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		return nil, fmt.Errorf("SUPABASE_DB_URL environment variable not set")
	}

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err) // Wrap error for context
	}
	return conn, nil
}

func CloseConnection(ctx context.Context, conn *pgx.Conn) {
	conn.Close(ctx)
}
