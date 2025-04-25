package connectiondb

import (
	"context"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v4"
	"time"
)

func dbconnect() error {

	err := godotenv.Load("/root/stocksapp/.env")
	if err != nil {
		return fmt.Errorf("error cargando .env: %w", err)
	}

	user := os.Getenv("USERDB")
	pass := os.Getenv("PASSDB")
	host := os.Getenv("HOSTDB")
	database := os.Getenv("DATABASE")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:26257/%s?sslmode=verify-full", user, pass, host, database)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	defer conn.Close(context.Background())	


	var now time.Time
	err = conn.QueryRow(ctx, "SELECT NOW()" ).Scan(&now)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	fmt.Println(now)
	return nil
}