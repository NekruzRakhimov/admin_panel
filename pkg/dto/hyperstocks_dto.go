package dto

import (
	"admin_panel/models"
	"net/url"
)

func ParseHyperstocksSearchParameters(values url.Values) (*models.HyperstocksSearchParameters, error) {
	dateFilter, err := ParseDateFilter(values, "date")
	if err != nil {
		return nil, err
	}

	return &models.HyperstocksSearchParameters{
		Pharmacy: ParseStringFilter(values, "pharmacy"),
		Date:     dateFilter,
	}, nil
}
