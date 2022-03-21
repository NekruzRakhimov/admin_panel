package repository

import (
	"admin_panel/db"
	"admin_panel/models"
)

func GetPurchase(bin string) []models.DataPurchase {

	var Purchase []models.DataPurchase

	db.GetDBConn().Raw("SELECT id, contract_parameters ->> 'contract_number' AS contract_number, discounts ->> 'discount_amount' "+
		"AS discount_amount, requisites ->> 'bin' AS bin, contract_parameters ->> start_date as start_date, "+
		"contract_parameters ->> end_date as end_date"+
		"FROM contracts WHERE requisites ->> 'bin' = $1", bin).Scan(&Purchase)

	return Purchase
}
