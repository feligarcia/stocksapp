package db

import (
	"context"
	"testing"
	"time"
)

// Para probar que DbConnect funcione y la conexión esté viva
func TestDbConnect(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect()
	if err != nil {
		t.Fatalf("Fallo la conexion a la base de datos: %v", err)
	}
	defer conn.Close(ctx)

	// Verificar que se puede hacer una consulta
	var now time.Time
	err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	if err != nil {
		t.Fatalf("No se pudo ejecutar SELECT NOW(): %v", err)
	}
	t.Logf("Conexion exitosa, hora actual: %s", now)
}
