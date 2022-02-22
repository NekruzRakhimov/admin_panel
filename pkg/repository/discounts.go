package repository

import (
	"admin_panel/db"
	"admin_panel/model"
	"log"
)

func GetAllContractDetailByBIN(bin, PeriodFrom, PeriodTo string) (contracts []model.ContractWithJsonB, err error) {
	if err = db.GetDBConn().Table("contracts").
		Where(`requisites ->> 'bin' = ? 
					AND contract_parameters ->> 'start_date' >= ?`, bin, PeriodFrom).
		Find(&contracts).Error; err != nil {
		log.Println("[repository][GetAllContractDetailByBIN] error is: ", err.Error())
		return nil, err
	}

	return contracts, nil
}
