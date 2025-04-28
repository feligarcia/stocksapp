package db

type CompanyProfile struct {
	Ticker           string  `json:"ticker"`
	Name             string  `json:"name"`
	Country          string  `json:"country"`
	Currency         string  `json:"currency"`
	Exchange         string  `json:"exchange"`
	FinnhubIndustry  string  `json:"finnhub_industry"`
	Ipo              string  `json:"ipo"`
	Logo             string  `json:"logo"`
	Marketcap        float64 `json:"marketcap"`
	Acciones         float64 `json:"acciones"`
	Web              string  `json:"web"`
	Telefono         string  `json:"telefono"`
}

type Quote struct {
	Ticker string  `json:"ticker"`
	C      float64 `json:"c"`
	D      float64 `json:"d"`
	Dp     float64 `json:"dp"`
	H      float64 `json:"h"`
	L      float64 `json:"l"`
	O      float64 `json:"o"`
	Pc     float64 `json:"pc"`
	T      int64   `json:"t"`
}
