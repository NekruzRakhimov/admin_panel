package models

import (
	"time"
)

type DateFilter struct {
	From *time.Time
	To   *time.Time
}

func (d DateFilter) WithTime() *DateFilter {
	from := time.Date(d.From.Year(), d.From.Month(), d.From.Day(), 0, 0, 0, 0, d.From.Location())
	to := time.Date(d.To.Year(), d.To.Month(), d.To.Day(), 23, 59, 59, 999999999, d.To.Location())
	return NewDateFilter(&from, &to)
}

func NewDateFilter(from *time.Time, to *time.Time) *DateFilter {
	return &DateFilter{From: from, To: to}
}
