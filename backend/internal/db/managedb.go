package db

import (
    "context"
    "log"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
)

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

// Crea la tabla tickers
func CrearTablaTickers(ctx context.Context, tx pgx.Tx) error {
    log.Println("Eliminando la tabla 'tickers' si existe...")
    if _, err := tx.Exec(ctx, "DROP TABLE IF EXISTS tickers"); err != nil {
        return err
    }
    log.Println("Creando la tabla 'tickers'...")
    query := `CREATE TABLE tickers (
        ticker STRING PRIMARY KEY,
        nombre STRING UNIQUE,
        pais STRING,
        bolsa STRING,
        industria STRING,
        logo STRING
    )`
    if _, err := tx.Exec(ctx, query); err != nil {
        return err
    }
    log.Println("Tabla 'tickers' creada exitosamente.")
    return nil
}

// Crea la tabla basicfinancial
func CrearTablaBasicFinancial(ctx context.Context, tx pgx.Tx) error {
    log.Println("Eliminando la tabla 'basicfinancial' si existe...")
    if _, err := tx.Exec(ctx, "DROP TABLE IF EXISTS basicfinancial"); err != nil {
        return err
    }
    log.Println("Creando la tabla 'basicfinancial'...")
    query := `CREATE TABLE basicfinancial (
        ticker STRING PRIMARY KEY REFERENCES tickers(ticker),
        marketcap FLOAT8,
        acciones FLOAT8,
        web STRING UNIQUE,
        telefono STRING
    )`
    if _, err := tx.Exec(ctx, query); err != nil {
        return err
    }
    log.Println("Tabla 'basicfinancial' creada exitosamente.")
    return nil
}
// Funciones de ejemplo de https://www.cockroachlabs.com/docs/v25.1/build-a-go-app-with-cockroachdb?filters=local
func insertRows(ctx context.Context, tx pgx.Tx, accts [4]uuid.UUID) error {
    // Insert four rows into the "accounts" table.
    log.Println("Creating new rows...")
    if _, err := tx.Exec(ctx,
        "INSERT INTO accounts (id, balance) VALUES ($1, $2), ($3, $4), ($5, $6), ($7, $8)", accts[0], 250, accts[1], 100, accts[2], 500, accts[3], 300); err != nil {
        return err
    }
    return nil
}

func printBalances(conn *pgx.Conn) error {
    rows, err := conn.Query(context.Background(), "SELECT id, balance FROM accounts")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        var id uuid.UUID
        var balance int
        if err := rows.Scan(&id, &balance); err != nil {
            log.Fatal(err)
        }
        log.Printf("%s: %d\n", id, balance)
    }
    return nil
}

func transferFunds(ctx context.Context, tx pgx.Tx, from uuid.UUID, to uuid.UUID, amount int) error {
    // Read the balance.
    var fromBalance int
    if err := tx.QueryRow(ctx,
        "SELECT balance FROM accounts WHERE id = $1", from).Scan(&fromBalance); err != nil {
        return err
    }

    if fromBalance < amount {
        log.Println("insufficient funds")
    }

    // Perform the transfer.
    log.Printf("Transferring funds from account with ID %s to account with ID %s...", from, to)
    if _, err := tx.Exec(ctx,
        "UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, from); err != nil {
        return err
    }
    if _, err := tx.Exec(ctx,
        "UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, to); err != nil {
        return err
    }
    return nil
}

func deleteRows(ctx context.Context, tx pgx.Tx, one uuid.UUID, two uuid.UUID) error {
    // Delete two rows into the "accounts" table.
    log.Printf("Deleting rows with IDs %s and %s...", one, two)
    if _, err := tx.Exec(ctx,
        "DELETE FROM accounts WHERE id IN ($1, $2)", one, two); err != nil {
        return err
    }
    return nil
}

// (main eliminada, este archivo solo contiene funciones utilitarias)
// Para crear tablas, ejecuta setup.go con package main