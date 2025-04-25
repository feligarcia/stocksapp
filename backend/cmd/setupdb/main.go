package main

import (
	"context"
	crdbpgxv5 "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/jackc/pgx/v5"
	"log"
	"stocksapp/backend/internal/db"
)

func main() {
	conn, err := db.DbConnect()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer conn.Close(context.Background())

	// Crear tabla 'tickers'
	err = crdbpgxv5.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return db.CrearTablaTickers(context.Background(), tx)
	})
	if err != nil {
		log.Fatal("Error creando la tabla 'tickers': ", err)
	}

	// Crear tabla 'basicfinancial'
	err = crdbpgxv5.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return db.CrearTablaBasicFinancial(context.Background(), tx)
	})
	if err != nil {
		log.Fatal("Error creando la tabla 'basicfinancial': ", err)
	}

	log.Println("Tablas creadas exitosamente.")
}
