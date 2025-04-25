package externalapi

import (
	"context"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

func initFinnhubClient() *finnhub.DefaultApiService {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println("Error cargando .env:", err)
	}
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", os.Getenv("FINNHUB_TOKEN"))
	return finnhub.NewAPIClient(cfg).DefaultApi
}

// GetCompanyProfile retorna la info del perfil de la empresa para un ticker dado
func GetCompanyProfile(symbol string) (*finnhub.CompanyProfile2, error) {
	finnhubClient := initFinnhubClient()
	profile, _, err := finnhubClient.CompanyProfile2(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

//Definir si GetStocksBasic lo voy a usar para las recomendaciones o para graficar

// GetStocksBasic retorna la info financiera básica para un ticker dado
// func GetStocksBasic(symbol string) (*finnhub.BasicFinancials, error) {
// 	finnhubClient := initFinnhubClient()
// 	basic, _, err := finnhubClient.CompanyBasicFinancials(context.Background()).Symbol(symbol).Metric("all").Execute()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &basic, nil
// }

// GetStocksQuote retorna el quote para un ticker dado
func GetStocksQuote(symbol string) (*finnhub.Quote, error) {
	finnhubClient := initFinnhubClient()
	quote, _, err := finnhubClient.Quote(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		return nil, err
	}
	return &quote, nil
}