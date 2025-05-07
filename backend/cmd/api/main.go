package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"stocksapp/backend/internal/db"
	"strconv"
)

func main() {
	app := fiber.New()
	app.Get("/api/companyinfo/:ticker", func(c *fiber.Ctx) error {
		ticker := c.Params("ticker")
		if ticker == "" {
			return c.Status(400).JSON(fiber.Map{"error": "ticker es requerido"})
		}
		conn, err := db.DbConnect()
		if err != nil {
			log.Printf("Error conectando a la base de datos: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "No se pudo conectar a la base de datos"})
		}
		defer conn.Close(context.Background())

		company, err := db.GetCompanyProfile(context.Background(), conn, ticker)
		if err != nil {
			log.Printf("Error obteniendo company info para %s: %v", ticker, err)
			return c.Status(404).JSON(fiber.Map{"error": "No se encontró información para ese ticker"})
		}
		return c.JSON(company)
	})

	distDir := "../frontend/dist"
	if _, err := os.Stat(distDir); os.IsNotExist(err) {
		app.Get("*", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"error": "server corriendo, dist no encontrado",
			})
		})
	} else {
		// Servir dist si el directorio existe
		app.Static("/", distDir)
	}

	app.Get("/api/tickers", func(c *fiber.Ctx) error {
		conn, err := db.DbConnect()
		if err != nil {
			log.Printf("Error conectando a la base de datos: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "No se pudo conectar a la base de datos"})
		}
		defer conn.Close(context.Background())

		tickers, err := db.GetAllTickers(context.Background(), conn)
		if err != nil {
			log.Printf("Error obteniendo tickers: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "No se pudieron obtener los tickers"})
		}
		return c.JSON(tickers)
	})
	//paginacion tickers
	app.Get("/api/tickers2/:page", func(c *fiber.Ctx) error {
		page := c.Params("page")
		conn, err := db.DbConnect()
		if err != nil {
			log.Printf("Error conectando a la base de datos: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "No se pudo conectar a la base de datos"})
		}
		defer conn.Close(context.Background())
		limit := 10
		intpage, err := strconv.Atoi(page)
		offset := 10 * (intpage - 1)

		tickers, err := db.GetTickersPagination(context.Background(), conn, offset, limit)
		if err != nil {
			log.Printf("Error obteniendo tickers: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "No se pudieron obtener los tickers"})
		}
		return c.JSON(tickers)
	})

	//paginacion quotes
	app.Get("/api/quotes2/:page", func(c *fiber.Ctx) error {
		page := c.Params("page")
		conn, err := db.DbConnect()
		if err != nil {
			log.Printf("Error conectando a la base de datos: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "No se pudo conectar a la base de datos"})
		}
		defer conn.Close(context.Background())
		limit := 10
		intpage, err := strconv.Atoi(page)
		offset := 10 * (intpage - 1)

		tickers, err := db.GetQuotesPagination(context.Background(), conn, offset, limit)
		if err != nil {
			log.Printf("Error obteniendo tickers: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "No se pudieron obtener los tickers"})
		}
		return c.JSON(tickers)
	})

	app.Get("/api/quotes", func(c *fiber.Ctx) error {
		conn, err := db.DbConnect()
		if err != nil {
			log.Printf("Error conectando a la base de datos: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "No se pudo conectar a la base de datos"})
		}
		defer conn.Close(context.Background())

		quotes, err := db.GetAllQuotes(context.Background(), conn)
		if err != nil {
			log.Printf("Error obteniendo quotes: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "No se pudieron obtener los quotes"})
		}
		return c.JSON(quotes)
	})

	// Si no hay un ticker en la URL se traen todos
	app.Get("/api/quotes/:ticker", func(c *fiber.Ctx) error {
		conn, err := db.DbConnect()
		if err != nil {
			log.Printf("Error conectando a la base de datos: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "No se pudo conectar a la base de datos"})
		}
		defer conn.Close(context.Background())

		ticker := c.Params("ticker")
		quotes, err := db.GetQuotesByTicker(context.Background(), conn, ticker)
		if err != nil {
			log.Printf("Error obteniendo quote para %s: %v", ticker, err)
			return c.Status(500).JSON(fiber.Map{"error": "Error consultando la base de datos"})
		}
		if len(quotes) == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "No se encontró información para ese ticker"})
		}
		return c.JSON(quotes)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
