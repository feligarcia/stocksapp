package db

import (
	"context"
	"time"
	"github.com/jackc/pgx/v5"
)

//Querys para cada funcion de la api, seria como los service en nestjs
//Traer perfil de una empresa
func GetCompanyProfile(ctx context.Context, conn *pgx.Conn, ticker string) (*CompanyProfile, error) {
	row := conn.QueryRow(ctx, `SELECT ticker, name, country, currency, exchange, finnhub_industry, ipo, logo, marketcap, acciones, web, telefono FROM companyprofile WHERE ticker = $1`, ticker)
	var cp CompanyProfile
	var ipo time.Time
	err := row.Scan(&cp.Ticker, &cp.Name, &cp.Country, &cp.Currency, &cp.Exchange, &cp.FinnhubIndustry, &ipo, &cp.Logo, &cp.Marketcap, &cp.Acciones, &cp.Web, &cp.Telefono)
	if err != nil {
		return nil, err
	}
	cp.Ipo = ipo.Format("2025-12-30") // formato YYYY-MM-DD
	return &cp, nil
}

//Traer todos los tickers
func GetAllTickers(ctx context.Context, conn *pgx.Conn) ([]string, error) {
	rows, err := conn.Query(ctx, "SELECT ticker FROM tickers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickers []string
	for rows.Next() {
		var ticker string
		if err := rows.Scan(&ticker); err != nil {
			return nil, err
		}
		tickers = append(tickers, ticker)
	}
	return tickers, nil
}

func GetTickersPagination(ctx context.Context, conn *pgx.Conn, offset int, limit int) ([]string, error) {
	rows, err := conn.Query(ctx, "SELECT ticker FROM tickers OFFSET $1 LIMIT $2", offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickers []string
	for rows.Next() {
		var ticker string
		if err := rows.Scan(&ticker); err != nil {
			return nil, err
		}
		tickers = append(tickers, ticker)
	}
	return tickers, nil
}

func GetQuotesPagination(ctx context.Context, conn *pgx.Conn, offset int, limit int) ([]Quote, error) {
    rows, err := conn.Query(ctx, "SELECT ticker, c, d, dp, h, l, o, pc, t FROM quote OFFSET $1 LIMIT $2", offset, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var quotes []Quote
    for rows.Next() {
        var q Quote
        err := rows.Scan(&q.Ticker, &q.C, &q.D, &q.Dp, &q.H, &q.L, &q.O, &q.Pc, &q.T)
        if err != nil {
            return nil, err
        }
        quotes = append(quotes, q)
    }
    return quotes, nil
}
//Traer todas las quotes
func GetAllQuotes(ctx context.Context, conn *pgx.Conn) ([]Quote, error) {
	rows, err := conn.Query(ctx, "SELECT ticker, c, d, dp, h, l, o, pc, t FROM quote")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []Quote
	for rows.Next() {
		var q Quote
		if err := rows.Scan(&q.Ticker, &q.C, &q.D, &q.Dp, &q.H, &q.L, &q.O, &q.Pc, &q.T); err != nil {
			return nil, err
		}
		quotes = append(quotes, q)
	}
	return quotes, nil
}

//Traer todas las quotes de un ticker
func GetQuotesByTicker(ctx context.Context, conn *pgx.Conn, ticker string) ([]Quote, error) {
	rows, err := conn.Query(ctx, "SELECT ticker, c, d, dp, h, l, o, pc, t FROM quote WHERE ticker = $1", ticker)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []Quote
	for rows.Next() {
		var q Quote
		if err := rows.Scan(&q.Ticker, &q.C, &q.D, &q.Dp, &q.H, &q.L, &q.O, &q.Pc, &q.T); err != nil {
			return nil, err
		}
		quotes = append(quotes, q)
	}
	return quotes, nil
}
