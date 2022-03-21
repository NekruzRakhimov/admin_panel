package repository

import (
	"admin_panel/db"
	"admin_panel/models"
	"fmt"
	"log"
)

func CreateContractss(contractWithJson models.ContractWithJsonB) error {
	fmt.Printf(">>>> %+v", contractWithJson)
	if err := db.GetDBConn().Table("contracts").Omit("created_at", "updated_at", "is_extend_contract", "extend_date").Create(&contractWithJson).Error; err != nil {
		log.Println("[repository.CreateContract]|[db.GetDBConn().Table(\"contracts\").Create(&contractWithJson).Error]| error is: ", err.Error())
		return err
	}

	for _, value := range contractWithJson.DiscountBrand {
		err := db.GetDBConn().Exec("INSERT INTO brands(brand as brand_name, discount_percent, contract_id) VALUES ($1, $2, $3)", value.BrandName, value.DiscountPercent, contractWithJson.ID).Error
		if err != nil {
			return err
		}

	}

	return nil
}

func GetBrandInfo(bin string) ([]models.BrandInfo, error) {
	var brandsInfo []models.BrandInfo
	err := db.GetDBConn().Raw("SELECT c.id, b.contract_id, c.contract_parameters ->> 'contract_number' AS contract_number, b.discount_percent, b.brand FROM brands b "+
		"JOIN contracts  c ON b.contract_id = c.id WHERE c.requisites ->> 'bin' = $1", bin).Scan(&brandsInfo).Error
	if err != nil {
		return nil, err
	}

	return brandsInfo, nil
}

func GetIDByBIN(bin string) ([]models.BrandAndPercent, string) {
	var BrandsAndDiscount []models.BrandAndPercent
	var ContractsID models.ContractID

	// тут по БИНу получаю номер договора
	// ID Договора я ему возвращаю тут получается
	db.GetDBConn().Raw("SELECT id, contract_parameters ->> 'contract_number' AS contract_number FROM contracts WHERE requisites ->> 'bin' = $1", bin).Scan(&ContractsID)

	log.Println("ID CONTRACT", ContractsID)

	db.GetDBConn().Raw("SELECT brand AS brand_name, discount_percent, contract_id FROM  brands WHERE contract_id = $1", ContractsID.Id).Scan(&BrandsAndDiscount)

	//TODO: я тут не возвращаю ID договора
	return BrandsAndDiscount, ContractsID.ContractNumber
}
