package models

import "time"

type ForecastSearchParameters struct {
	ProductCode  *string
	PharmacyCode *string
}

type Forecast struct {
	HistorySales  *HistoricalSales
	ForecastSales *ForecastSales
}

type ForecastSales struct {
	SalesArr []Sale
}

type HistoricalSales struct {
	SalesArr []Sale `json:"sales_arr"`
}

type Sale struct {
	QntTotal float64   `json:"qnt_total"`
	Date     time.Time `json:"date"`
}
