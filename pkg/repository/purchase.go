package repository

import (
	"admin_panel/db"
	"admin_panel/model"
)

func GetPurchase(bin string) model.DataPurchase {

	var Purchase model.DataPurchase

	db.GetDBConn().Raw("SELECT id, contract_parameters ->> 'contract_number' AS contract_number, discounts ->> 'discount_amount' AS discount_amount, requisites ->> 'bin' AS bin FROM contracts WHERE requisites ->> 'bin' = $1", bin).Scan(&Purchase)

	return Purchase
}
