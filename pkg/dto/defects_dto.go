package dto

import (
	"admin_panel/models"
	"net/url"
)

func ParseDefectsSearchParameters(values url.Values) (*models.DefectsSearchParameters, error) {
	dateFilter, err := ParseDateFilter(values, "date")
	if err != nil {
		return nil, err
	}

	return &models.DefectsSearchParameters{
		Pharmacy: ParseStringFilter(values, "pharmacy"),
		Date:     dateFilter,
	}, nil
}
