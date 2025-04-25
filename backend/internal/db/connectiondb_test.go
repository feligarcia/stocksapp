package db

import (
	"context"
	"testing"
	"time"
)

func TestDbConnect(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect()
	if err != nil {
		t.Fatalf("Fallo la conexion a la base de datos: %v", err)
	}
	defer conn.Close(ctx)

	var now time.Time
	query := "SELECT NOW()"
	err = conn.QueryRow(ctx, query).Scan(&now)
	if err != nil {
		t.Fatalf("No se pudo ejecutar SELECT NOW(): %v", err)
	}
	t.Logf("Conexion exitosa, hora actual: %s", now)
}
