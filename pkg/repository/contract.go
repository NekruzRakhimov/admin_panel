package repository

import (
	"admin_panel/db"
	"admin_panel/model"
)

func GetAllCurrency() (currency []model.Currency, err error) {
	db.GetDBConn().Find(&currency)
	if currency == nil{
		return nil, err
	}

	return currency, err


}

