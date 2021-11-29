package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
)

func GetAllCurrency() (rights []model.Currency, err error) {
	return repository.GetAllCurrency()
}

