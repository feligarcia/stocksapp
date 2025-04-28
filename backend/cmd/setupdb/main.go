package main

import (
	"context"
	"log"
	"stocksapp/backend/internal/db"
	"stocksapp/backend/internal/externalapi"
	"time"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/jackc/pgx/v5"
)
//Este archivo se usa para alimentar la base de datos de la información de la API
func main() {

	var listtickers = []string{"AAPL", "MSFT", "GOOGL", "AMZN", "TSLA", "META", "BRK.B", "NVDA", "UNH", "JNJ",
		"V", "XOM", "WMT", "JPM", "PG", "MA", "CVX", "LLY", "HD", "ABBV",
		"MRK", "PEP", "KO", "BAC", "AVGO", "COST", "MCD", "DIS", "ADBE", "CSCO",
		"PFE", "TMO", "NKE", "DHR", "WFC", "ABT", "ACN", "TXN", "CRM", "LIN",
		"AMD", "CMCSA", "NEE", "UPS", "MS", "INTC", "ORCL", "PM", "QCOM", "UNP",
		"AMGN", "RTX", "HON", "BMY", "IBM", "GE", "BA", "MDT", "LOW", "SBUX",
		"CAT", "BLK", "AXP", "GS", "CVS", "SPGI", "ISRG", "GILD", "LMT", "T",
		"PLD", "C", "NOW", "ADP", "ZTS", "INTU", "MU", "DE", "CI", "TJX",
		"MDLZ", "MO", "MMC", "AMAT", "PNC", "SYK", "BDX", "SCHW", "ETN", "ELV",
		"ADI", "FI", "CB", "BSX", "APD", "BKNG", "REGN", "CL", "GM", "F",
		"VRTX", "EOG", "ITW", "EMR", "ROST", "AON", "PSX", "HUM", "CSX", "MAR"}

	conn, err := db.DbConnect()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer conn.Close(context.Background())

	// Crear tabla 'tickers'
	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return db.CrearTablaTickers(context.Background(), tx)
	})
	if err != nil {
		log.Fatal("Error creando la tabla 'tickers': ", err)
	}

	// Crear tabla 'companyprofile'
	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return db.CrearTablaCompanyProfile(context.Background(), tx)
	})
	if err != nil {
		log.Fatal("Error creando la tabla 'companyprofile': ", err)
	}

	// Crear tabla 'quote'
	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return db.CrearTablaQuote(context.Background(), tx)
	})
	if err != nil {
		log.Fatal("Error creando la tabla 'quote': ", err)
	}

	// Crear tabla 'basicfinancial' si no existe
	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return db.CrearTablaBasicFinancial(context.Background(), tx)
	})
	if err != nil {
		log.Fatal("Error creando la tabla 'basicfinancial': ", err)
	}

	log.Println("Tablas creadas exitosamente.")

	// Alimentar las tablas con datos de la api, por el momento son 3 tablas
	for _, ticker := range listtickers {
		// Obtener datos de la API
		profile, err := externalapi.GetCompanyProfile(ticker)
		if err != nil {
			log.Printf("Error obteniendo perfil para %s: %v", ticker, err)
			continue
		}
		quote, err := externalapi.GetStocksQuote(ticker)
		if err != nil {
			log.Printf("Error obteniendo quote para %s: %v", ticker, err)
			continue
		}
		err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			if err := db.InsertTicker(context.Background(), tx, ticker); err != nil {
				return err
			}
			log.Printf("%s insertado exitosamente en tickers", ticker)
			if err := db.InsertCompanyProfile(context.Background(), tx, ticker, profile); err != nil {
				return err
			}
			log.Printf("%s insertado exitosamente en companyprofile", ticker)
			if err := db.InsertQuote(context.Background(), tx, ticker, quote); err != nil {
				return err
			}
			log.Printf("%s insertado exitosamente en quote", ticker)
			return nil
		})
		if err != nil {
			log.Printf("Error insertando datos para %s: %v", ticker, err)
		}
		time.Sleep(time.Second) //Esperar 1 segundo para cumplir el limite de la api free que son 60api call/min
	}

	log.Println("TODOS los Datos insertados exitosamente.")

}
