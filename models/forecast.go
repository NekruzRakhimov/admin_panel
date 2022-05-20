package models

type ForecastSearchParameters struct {
	ProductCode  *string
	PharmacyCode *string
}

type Forecast struct {
	Sales []Sale
}

type ForecastSales struct {
	SalesArr []Sale
}

type HistoricalSales struct {
	SalesArr []Sale `json:"sales_arr"`
}

type Sale struct {
	QntTotal float64 `json:"qnt_total"`
	Date     string  `json:"date"`
	Category string  `json:"category"`
}
