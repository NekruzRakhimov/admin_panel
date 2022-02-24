package repository

import (
	"admin_panel/db"
	"admin_panel/model"
	"log"
)

func GetAllContractDetailByBIN(bin, PeriodFrom, PeriodTo string) (contracts []model.ContractWithJsonB, err error) {
	if err = db.GetDBConn().Table("contracts").
		Where(`requisites ->> 'bin' = ? 	
					AND contract_parameters ->> 'start_date' >= ? AND contract_parameters ->> 'end_date' <= ?`, bin, PeriodFrom, PeriodTo).
		Find(&contracts).Error; err != nil {
		log.Println("[repository][GetAllContractDetailByBIN] error is: ", err.Error())
		return nil, err
	}

	//var brands []model.DiscountBrand
	for i, contract := range contracts {
		if err = db.GetDBConn().Raw("SELECT id, brand as brand_name, brand_code, discount_percent FROM  brands  WHERE  contract_id = ?", contract.ID).Scan(&contracts[i].DiscountBrand).Error; err != nil {
			return nil, err
		}

		log.Println("BRANDS", contracts[i].DiscountBrand)
	}

	return contracts, nil
}
