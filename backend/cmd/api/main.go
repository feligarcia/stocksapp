package main

import (
	"context"
	"os"
	"time"
	"log"
	"github.com/gofiber/fiber/v2"
	"stocksapp/backend/internal/db"
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hora_actual": time.Now().Format(time.RFC3339),
		})
	})

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