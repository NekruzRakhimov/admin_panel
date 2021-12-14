package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
)

func GetAllCurrencies() ([]model.Currency, error) {
	return repository.GetAllCurrencies()
}

func GetAllPositions() ([]model.Position, error) {
	return repository.GetAllPositions()
}

func GetAllAddresses() ([]model.Address, error) {
	return repository.GetAllAddresses()
}

func GetAllFrequencyDeferredDiscounts() ([]model.FrequencyDeferredDiscount, error) {
	return repository.GetAllFrequencyDeferredDiscounts()
}
