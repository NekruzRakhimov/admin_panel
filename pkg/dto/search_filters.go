package dto

import (
	"admin_panel/models"
	"errors"
	"net/url"
	"time"
)

func parseDateFilterValue(value string) (*time.Time, error) {
	if value == "" {
		return nil, nil
	}

	fromValue, err := time.Parse("02.01.2006", value)
	if err != nil {
		return nil, errors.New("invalid date filter")
	}

	return &fromValue, nil
}

func ParseDateFilter(values url.Values, key string) (*models.DateFilter, error) {
	dates, ok := values[key]
	if !ok {
		return nil, nil
	}

	if len(dates) != 2 {
		return nil, errors.New("invalid date filter")
	}

	fromValue, err := parseDateFilterValue(dates[0])
	if err != nil {
		return nil, err
	}

	toValue, err := parseDateFilterValue(dates[1])
	if err != nil {
		return nil, err
	}

	return models.NewDateFilter(fromValue, toValue), nil
}

func ParseStringFilter(values url.Values, key string) *string {
	statusFilterStr := values.Get(key)
	if statusFilterStr == "" {
		return nil
	}

	return &statusFilterStr
}
