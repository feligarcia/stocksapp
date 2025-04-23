package connectiondb

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v4"
)

func dbconnect() {

	err := godotenv.Load("/root/stocksapp/.env")
	if err != nil {
		log.Fatalf("Error cargando .env: %v", err)
	}

	user := os.Getenv("USERDB")
	pass := os.Getenv("PASSDB")
	host := os.Getenv("HOSTDB")
	database := os.Getenv("DATABASE")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:26257/%s?sslmode=verify-full", user, pass, host, database)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer conn.Close(context.Background())	


	var now time.Time
	err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	fmt.Println(now)
}