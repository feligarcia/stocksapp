package db

import (
	"context"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"
)

// DbConnect abre una conexión a la base de datos CockroachDB usando las variables de entorno y .env en backend.
// Retorna *pgx.Conn y error, y usa sslmode=require por defecto.
func DbConnect() (*pgx.Conn, error) {
	err := godotenv.Load("/root/stocksapp/backend/.env") // busca .env en el cwd, que debe ser backend
	if err != nil {
		// No es fatal, puede que las variables ya estén en el entorno
	}
	user := os.Getenv("USERDB")
	pass := os.Getenv("PASSDB")
	host := os.Getenv("HOSTDB")
	database := os.Getenv("DATABASE")
	if user == "" || pass == "" || host == "" || database == "" {
		return nil, fmt.Errorf("faltan variables de entorno para la conexión: USERDB, PASSDB, HOSTDB, DATABASE")
	}
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:26257/%s?sslmode=require", user, pass, host, database)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	return conn, nil
}