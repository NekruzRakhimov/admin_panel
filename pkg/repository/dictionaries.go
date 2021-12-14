package repository

import (
	"admin_panel/db"
	"admin_panel/model"
)

func GetAllCurrencies() (currencies []model.Currency, err error) {
	sqlQuery := "SELECT * FROM currencies"
	if err := db.GetDBConn().Raw(sqlQuery).Scan(&currencies).Error; err != nil {
		return nil, err
	}

	return currencies, nil
}

func GetAllPositions() (positions []model.Position, err error) {
	sqlQuery := "SELECT * FROM positions"
	if err := db.GetDBConn().Raw(sqlQuery).Scan(&positions).Error; err != nil {
		return nil, err
	}

	return positions, nil
}

func GetAllAddresses() (addresses []model.Address, err error) {
	sqlQuery := "SELECT * FROM addresses"
	if err := db.GetDBConn().Raw(sqlQuery).Scan(&addresses).Error; err != nil {
		return nil, err
	}

	return addresses, nil
}

func GetAllFrequencyDeferredDiscounts() (frequencyDeferredDiscounts []model.FrequencyDeferredDiscount, err error) {
	sqlQuery := "SELECT * FROM frequency_deferred_discount"
	if err := db.GetDBConn().Raw(sqlQuery).Scan(&frequencyDeferredDiscounts).Error; err != nil {
		return nil, err
	}

	return frequencyDeferredDiscounts, nil
}
