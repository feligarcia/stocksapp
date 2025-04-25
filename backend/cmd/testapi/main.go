package main

import (
	"encoding/json"
	"fmt"
	"log"
	"stocksapp/backend/internal/externalapi"
)

func main() {
    //Probando GetStocksQuote y GetCompanyProfile para un ticker
	ticker := "AAPL"
	quote, err := externalapi.GetStocksQuote(ticker)
	if err != nil {
		log.Fatalf("Error obteniendo quote para %s: %v", ticker, err)
	}
	quoteJSON, err := json.MarshalIndent(quote, "", "  ")
	if err != nil {
		log.Fatalf("Error serializando quote a JSON: %v", err)
	}
	fmt.Printf("Quote para %s (JSON):\n%s\n", ticker, quoteJSON)

	profile, err := externalapi.GetCompanyProfile(ticker)
	if err != nil {
		log.Fatalf("Error obteniendo company profile para %s: %v", ticker, err)
	}
	profileJSON, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		log.Fatalf("Error serializando company profile a JSON: %v", err)
	}
	fmt.Printf("Profile para %s (JSON):\n%s\n", ticker, profileJSON)
}
