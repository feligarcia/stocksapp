package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

// Crea la tabla basicfinancial si no existe
func CrearTablaBasicFinancial(ctx context.Context, tx pgx.Tx) error {
	log.Println("Creando la tabla 'basicfinancial' si no existe...")
	query := `CREATE TABLE IF NOT EXISTS basicfinancial (
		-- Por definir
	)`
	if _, err := tx.Exec(ctx, query); err != nil {
		return err
	}
	log.Println("Tabla 'basicfinancial' creada exitosamente (si no existía).")
	return nil
}

// Crea la tabla tickers
func CrearTablaTickers(ctx context.Context, tx pgx.Tx) error {
	log.Println("Creando la tabla 'tickers' si no existe...")
	query := `CREATE TABLE IF NOT EXISTS tickers (
        ticker STRING PRIMARY KEY 
    )`
	if _, err := tx.Exec(ctx, query); err != nil {
		return err
	}
	log.Println("Tabla 'tickers' creada exitosamente.")
	return nil
}

// Crea la tabla quote para almacenar datos de GetStocksQuote
func CrearTablaQuote(ctx context.Context, tx pgx.Tx) error {
	log.Println("Creando la tabla 'quote' si no existe...")
	query := `CREATE TABLE IF NOT EXISTS quote (
        ticker STRING PRIMARY KEY REFERENCES tickers(ticker),
        c FLOAT8,    -- Precio actual
        d FLOAT8,    -- Cambio absoluto
        dp FLOAT8,   -- Cambio porcentual
        h FLOAT8,    -- Máximo del día
        l FLOAT8,    -- Mínimo del día
        o FLOAT8,    -- Precio de apertura
        pc FLOAT8,   -- Precio de cierre anterior
        t BIGINT     -- Timestamp
    )`
	if _, err := tx.Exec(ctx, query); err != nil {
		return err
	}
	log.Println("Tabla 'quote' creada exitosamente.")
	return nil
}

// Crea la tabla basicfinancial
func CrearTablaCompanyProfile(ctx context.Context, tx pgx.Tx) error {
	log.Println("Creando la tabla 'companyprofile' si no existe...")
	query := `CREATE TABLE IF NOT EXISTS companyprofile (
        ticker STRING PRIMARY KEY REFERENCES tickers(ticker),
        name STRING,
        country STRING,
        currency STRING,
        exchange STRING,
        finnhub_industry STRING,
        ipo DATE,
        logo STRING,
        marketcap FLOAT8,
        acciones FLOAT8,
        web STRING UNIQUE,
        telefono STRING
    )`
	if _, err := tx.Exec(ctx, query); err != nil {
		return err
	}
	log.Println("Tabla 'companyprofile' creada exitosamente.")
	return nil
}

// Inserta fila de ticker
func InsertTicker(ctx context.Context, tx pgx.Tx, ticker string) error {
	_, err := tx.Exec(ctx, `INSERT INTO tickers (ticker) VALUES ($1) ON CONFLICT DO NOTHING`, ticker)
	return err
}

// Inserta fila de perfil de compañía
func InsertCompanyProfile(ctx context.Context, tx pgx.Tx, ticker string, profile *finnhub.CompanyProfile2) error {
	_, err := tx.Exec(ctx, `INSERT INTO companyprofile (
		ticker, name, country, currency, exchange, finnhub_industry, ipo, logo, marketcap, acciones, web, telefono)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		ON CONFLICT (ticker) DO UPDATE SET
			name=$2, country=$3, currency=$4, exchange=$5, finnhub_industry=$6, ipo=$7, logo=$8, marketcap=$9, acciones=$10, web=$11, telefono=$12`,
		ticker,
		profile.Name,
		profile.Country,
		profile.Currency,
		profile.Exchange,
		profile.FinnhubIndustry,
		profile.Ipo,
		profile.Logo,
		profile.MarketCapitalization,
		profile.ShareOutstanding,
		profile.Weburl,
		profile.Phone,
	)
	return err
}

// Inserta fila de quote
func InsertQuote(ctx context.Context, tx pgx.Tx, ticker string, quote *finnhub.Quote) error {
	timestamp := time.Now().Unix()
	_, err := tx.Exec(ctx, `INSERT INTO quote (ticker, c, d, dp, h, l, o, pc, t)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (ticker) DO UPDATE SET c=$2, d=$3, dp=$4, h=$5, l=$6, o=$7, pc=$8, t=$9`,
		ticker,
		quote.C, quote.D, quote.Dp, quote.H, quote.L, quote.O, quote.Pc, timestamp,
	)
	return err
}




// func deleteRows(ctx context.Context, tx pgx.Tx, one uuid.UUID, two uuid.UUID) error {
// 	// Delete two rows into the "accounts" table.
// 	log.Printf("Deleting rows with IDs %s and %s...", one, two)
// 	if _, err := tx.Exec(ctx,
// 		"DELETE FROM accounts WHERE id IN ($1, $2)", one, two); err != nil {
// 		return err
// 	}
// 	return nil
// }
