package main

import "stocksapp/backend/internal/externalapi"

func main() {
    externalapi.GetStocksBasic()
    externalapi.GetStocksQuote()
    externalapi.GetCompanyProfile()
}
