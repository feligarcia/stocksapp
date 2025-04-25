package connectiondb

import (
	"testing"
)

//Para probar que dbconnect() funcione
func TestDbConnect(t *testing.T) {
	err := dbconnect()
	if err != nil {
		t.Fatalf("Fallo la conexion a la base de datos: %v", err)
	}
}
