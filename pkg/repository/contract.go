package repository

import (
	"admin_panel/db"
	"admin_panel/model"
	"fmt"
	"log"
)

func CreateContract(contractWithJson model.ContractWithJsonB) error {
	fmt.Printf(">>>> %+v", contractWithJson)
	if err := db.GetDBConn().Table("contracts").Omit("created_at", "updated_at").Create(&contractWithJson).Error; err != nil {
		log.Println("[repository.CreateContract]|[db.GetDBConn().Table(\"contracts\").Create(&contractWithJson).Error]| error is: ", err.Error())
		return err
	}
	return nil
}

func EditContract(contractWithJson model.ContractWithJsonB) error {
	if err := db.GetDBConn().Table("contracts").Omit("created_at", "updated_at").Save(&contractWithJson).Error; err != nil {
		log.Println("[repository.EditContract]|[db.GetDBConn().Table(\"contracts\").Save(&contractWithJson).Error]| error is: ", err.Error())
		return err
	}
	return nil
}

func GetAllContracts(contractStatus string) (contracts []model.ContractWithJsonB, err error) {
	var contractStatusRus = ""
	sqlQuery := "SELECT * FROM contracts WHERE id not in (select prev_contract_id from contracts) AND is_active = true "
	if contractStatus != "" {
		switch contractStatus {
		case "DRAFT":
			contractStatusRus = "черновик"
		case "ON_APPROVAL":
			contractStatusRus = "на согласовании"
		case "ACTIVE":
			contractStatusRus = "в работе"
		case "EXPIRED":
			contractStatusRus = "заверщённый"
		}
		sqlQuery += fmt.Sprintf(" AND status = '%s'", contractStatusRus)
	}

	sqlQuery += " ORDER BY created_at DESC"

	if err := db.GetDBConn().Raw(sqlQuery).Scan(&contracts).Error; err != nil {
		log.Println("[repository.GetAllContracts]|[db.GetDBConn().Raw(sqlQuery).Scan(&contracts).Error]| error is: ", err.Error())
		return nil, err
	}

	return contracts, nil
}

func GetContractDetails(contractId int) (contract model.ContractWithJsonB, err error) {
	contract.ID = contractId
	if err := db.GetDBConn().Table("contracts").Find(&contract).Error; err != nil {
		return model.ContractWithJsonB{}, err
	}

	return contract, nil
}

func ConformContract(contractId int, status string) error {
	sqlQuery := "UPDATE contracts SET status = $1 WHERE id = $2"
	if err := db.GetDBConn().Exec(sqlQuery, status, contractId).Error; err != nil {
		return err
	}

	return nil
}

func ChangeContractStatus(contractId int, status string) error {
	sqlQuery := "UPDATE contracts SET status = $1 WHERE id = $2"
	if err := db.GetDBConn().Exec(sqlQuery, status, contractId).Error; err != nil {
		return err
	}

	return nil
}

func DisActiveContract(contractId int) error {
	sqlQuery := "UPDATE contracts SET is_active = false, updated_at = now() WHERE id = $1"
	if err := db.GetDBConn().Exec(sqlQuery, contractId).Error; err != nil {
		return err
	}

	return nil
}

func FinishContract(contractId int) error {
	sqlQuery := "UPDATE contracts SET status = ?, updated_at = now() WHERE id = ?"
	if err := db.GetDBConn().Exec(sqlQuery, "заверщённый", contractId).Error; err != nil {
		return err
	}

	return nil
}

func RevisionContract(contractId int, comment string) error {
	sqlQuery := "UPDATE contracts SET status = ?, comment = ?, updated_at = now() WHERE id = ?"
	//TODO: добавить проверку статуса договора
	if err := db.GetDBConn().Raw(sqlQuery, "черновик", comment, contractId).Error; err != nil {
		return err
	}

	return nil
}
