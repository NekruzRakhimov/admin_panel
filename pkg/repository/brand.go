package repository

import (
	"admin_panel/db"
	"admin_panel/model"
	"fmt"
	"log"
)

func CreateContractss(contractWithJson model.ContractWithJsonB) error {
	fmt.Printf(">>>> %+v", contractWithJson)
	if err := db.GetDBConn().Table("contracts").Omit("created_at", "updated_at", "is_extend_contract", "extend_date").Create(&contractWithJson).Error; err != nil {
		log.Println("[repository.CreateContract]|[db.GetDBConn().Table(\"contracts\").Create(&contractWithJson).Error]| error is: ", err.Error())
		return err
	}

	for _, value := range contractWithJson.DiscountBrand {
		err := db.GetDBConn().Exec("INSERT INTO brands(brand, discount_percent, contract_id) VALUES ($1, $2, $3)", value.BrandName, value.DiscountPercent, contractWithJson.ID).Error
		if err != nil {
			return err
		}

	}

	return nil
}

func GetBrandInfo(bin string) ([]model.BrandInfo, error) {
	var brandsInfo []model.BrandInfo
	err := db.GetDBConn().Raw("SELECT c.id, b.contract_id, c.contract_parameters ->> 'contract_number' AS contract_number, b.discount_percent, b.brand FROM brands b "+
		"JOIN contracts  c ON b.contract_id = c.id WHERE c.requisites ->> 'bin' = $1", bin).Scan(&brandsInfo).Error
	if err != nil {
		return nil, err
	}

	return brandsInfo, nil
}
