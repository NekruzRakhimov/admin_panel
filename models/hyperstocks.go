package models

type HyperstocksSearchParameters struct {
	Date     *DateFilter
	Pharmacy *string
}

type HyperstocksFile struct {
	FileName string
	File     string
}
