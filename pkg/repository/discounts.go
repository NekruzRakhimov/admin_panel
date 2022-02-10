package repository

import (
	"admin_panel/db"
	"admin_panel/model"
)

func GetAllContractDetailByBIN(bin string) (contracts []model.ContractWithJsonB, err error) {
	if err := db.GetDBConn().Table("contracts").Where("requisites ->> 'bin' = ?", bin).Find(&contracts).Error; err != nil {
		return nil, err
	}

	return contracts, nil
}
