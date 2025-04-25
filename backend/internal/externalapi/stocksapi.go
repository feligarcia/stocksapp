package externalapi

import (
	"context"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

func initFinnhubClient() *finnhub.DefaultApiService {
	err := godotenv.Load("/root/stocksapp/.env")
	if err != nil {
		fmt.Println("Error cargando .env:", err)
	}
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", os.Getenv("FINNHUB_TOKEN"))
	return finnhub.NewAPIClient(cfg).DefaultApi
}

func GetCompanyProfile() {
	finnhubClient := initFinnhubClient()
	companyprofile, _, err := finnhubClient.CompanyProfile2(context.Background()).Symbol("MSFT").Execute()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Info company:")
	if companyprofile.Name != nil {
		fmt.Println("  Nombre:", *companyprofile.Name)
	}
	if companyprofile.Ticker != nil {
		fmt.Println("  Ticker:", *companyprofile.Ticker)
	}
	if companyprofile.Country != nil {
		fmt.Println("  País:", *companyprofile.Country)
	}
	if companyprofile.Exchange != nil {
		fmt.Println("  Bolsa:", *companyprofile.Exchange)
	}
	if companyprofile.Currency != nil {
		fmt.Println("  Moneda:", *companyprofile.Currency)
	}
	if companyprofile.Ipo != nil {
		fmt.Println("  IPO:", *companyprofile.Ipo)
	}
	if companyprofile.MarketCapitalization != nil {
		fmt.Printf("  Market Cap: %.2f\n", *companyprofile.MarketCapitalization)
	}
	if companyprofile.ShareOutstanding != nil {
		fmt.Printf("  Acciones en circulación: %.2f\n", *companyprofile.ShareOutstanding)
	}
	if companyprofile.Weburl != nil {
		fmt.Println("  Web:", *companyprofile.Weburl)
	}
	if companyprofile.FinnhubIndustry != nil {
		fmt.Println("  Industria:", *companyprofile.FinnhubIndustry)
	}
	if companyprofile.Logo != nil {
		fmt.Println("  Logo URL:", *companyprofile.Logo)
	}
	if companyprofile.Phone != nil {
		fmt.Println("  Teléfono:", *companyprofile.Phone)
	}
}

func GetStocksBasic() {
	finnhubClient := initFinnhubClient()
	basicFinancials, _, err := finnhubClient.CompanyBasicFinancials(context.Background()).Symbol("MSFT").Metric("all").Execute()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Símbolo:", *basicFinancials.Symbol)
	fmt.Println("Tipo de Métrica:", *basicFinancials.MetricType)
	fmt.Println("Métricas:")
	if basicFinancials.Metric != nil {
		for k, v := range *basicFinancials.Metric {
			fmt.Printf("  %s: %v\n", k, v)
		}
	}
}

func GetStocksQuote() {
	finnhubClient := initFinnhubClient()
	quote, _, err := finnhubClient.Quote(context.Background()).Symbol("MSFT").Execute()
	if err != nil {
		fmt.Println("Error obteniendo quote:", err)
		return
	}
	fmt.Println("Quote de MSFT:")
	if quote.C != nil {
		fmt.Printf("  Precio actual: %.2f\n", *quote.C)
	} else {
		fmt.Println("  Precio actual: (sin datos)")
	}
	if quote.O != nil {
		fmt.Printf("  Precio de apertura: %.2f\n", *quote.O)
	} else {
		fmt.Println("  Precio de apertura: (sin datos)")
	}
	if quote.H != nil {
		fmt.Printf("  Precio más alto del día: %.2f\n", *quote.H)
	} else {
		fmt.Println("  Precio más alto del día: (sin datos)")
	}
	if quote.L != nil {
		fmt.Printf("  Precio más bajo del día: %.2f\n", *quote.L)
	} else {
		fmt.Println("  Precio más bajo del día: (sin datos)")
		
	}
	if quote.Pc != nil {
		fmt.Printf("  Precio al cierre del día anterior: %.2f\n", *quote.Pc)
	} else {
		fmt.Println("  Precio al cierre del día anterior: (sin datos)")
	}
}