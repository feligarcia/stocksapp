package db

import (
	"context"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"
)
//Conexion a la base de datos
//sslmode=require por defecto, en proudcicon sslmode=verify-full.
func DbConnect() (*pgx.Conn, error) {
	err := godotenv.Load("/root/stocksapp/backend/.env") // busca .env en el cwd, que debe ser backend
	if err != nil {
		return nil, fmt.Errorf("fallo al cargar el archivo .env: %w", err)
	}
	user := os.Getenv("USERDB")
	pass := os.Getenv("PASSDB")
	host := os.Getenv("HOSTDB")
	database := os.Getenv("DATABASE")
	port := os.Getenv("PORTDB")
	if port == "" {
		port = "26257" // valor por defecto CockroachDB
	}
	if user == "" || pass == "" || host == "" || database == "" {
		return nil, fmt.Errorf("faltan variables de entorno para la conexión: USERDB, PASSDB, HOSTDB, DATABASE")
	}
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=require", user, pass, host, port, database)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("fallo en la conexion a la base de datos: %w", err)
	}
	return conn, nil
}