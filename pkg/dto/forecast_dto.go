package dto

import (
	"admin_panel/models"
	"net/url"
)

func ParseForecastSearchParameters(values url.Values) (*models.ForecastSearchParameters, error) {
	return &models.ForecastSearchParameters{
		PharmacyCode: ParseStringFilter(values, "pharmacyCode"),
		ProductCode: ParseStringFilter(values,"productCode"),
	}, nil
}
