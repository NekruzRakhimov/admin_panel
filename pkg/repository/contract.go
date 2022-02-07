package repository

import (
	"admin_panel/db"
	"admin_panel/model"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"strings"
)

func CreateContract(contractWithJson model.ContractWithJsonB) error {
	fmt.Printf(">>>> %+v", contractWithJson)
	if err := db.GetDBConn().Table("contracts").Omit("created_at", "updated_at", "is_extend_contract").Create(&contractWithJson).Error; err != nil {
		log.Println("[repository.CreateContract]|[db.GetDBConn().Table(\"contracts\").Create(&contractWithJson).Error]| error is: ", err.Error())
		return err
	}
	return nil
}

func EditContract(contractWithJson model.ContractWithJsonB) error {
	if err := db.GetDBConn().Table("contracts").Omit("created_at", "updated_at", "is_extend_contract").Save(&contractWithJson).Error; err != nil {
		log.Println("[repository.EditContract]|[db.GetDBConn().Table(\"contracts\").Save(&contractWithJson).Error]| error is: ", err.Error())
		return err
	}

	if err := RecordContractStatusChange(contractWithJson.ID, contractWithJson.Status); err != nil {
		return err
	}

	return nil
}

func GetAllContracts(contractStatus string) (contracts []model.ContractWithJsonB, err error) {
	var contractStatusRus = ""
	sqlQuery := "SELECT * FROM contracts WHERE id not in (select prev_contract_id from contracts) AND is_active = true "
	if contractStatus == "ACTIVE_AND_EXPIRED" {
		contractStatus = ""
	}

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
		case "CANCELED":
			contractStatusRus = "отменен"
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

	if err := RecordContractStatusChange(contractId, status); err != nil {
		return err
	}

	return nil
}

func ChangeContractStatus(contractId int, status string) error {
	sqlQuery := "UPDATE contracts SET status = $1 WHERE id = $2"
	if err := db.GetDBConn().Exec(sqlQuery, status, contractId).Error; err != nil {
		return err
	}

	if err := RecordContractStatusChange(contractId, status); err != nil {
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

	if err := RecordContractStatusChange(contractId, "заверщённый"); err != nil {
		return err
	}

	return nil
}

func RevisionContract(contractId int, comment string) error {
	fmt.Println("[repository.RevisionContract]|[BEGIN]")
	sqlQuery := "UPDATE contracts SET status = ?, comment = ?, updated_at = now() WHERE id = ?"
	//TODO: добавить проверку статуса договора
	err := db.GetDBConn().Exec(sqlQuery, "черновик", comment, contractId).Error
	if err != nil {
		return err
	}

	if err := RecordContractStatusChange(contractId, "черновик"); err != nil {
		return err
	}

	fmt.Println("[repository.RevisionContract]|[END]")
	return nil
}

func GetContractStatusChangesHistory(contractId int) (history []model.ContractStatusHistory, err error) {
	sqlQuery := "SELECT * FROM status_changes_history WHERE contract_id = ?"
	if err := db.GetDBConn().Raw(sqlQuery, contractId).Scan(&history).Error; err != nil {
		return nil, err
	}

	return history, nil
}

func RecordContractStatusChange(contractId int, status string) error {
	sqlQuery := "INSERT INTO status_changes_history (contract_id, status) VALUES (?, ?)"

	if err := db.GetDBConn().Exec(sqlQuery, contractId, status).Error; err != nil {
		return err
	}

	return nil
}

func SaveContractExternalCode(contractId int, contractCode string) error {

	return db.GetDBConn().Raw("UPDATE  contracts set ext_contract_code = $1  WHERE id = $2", contractCode, contractId).Error

}

func SearchContractByNumber(param string, status string) ([]model.SearchContract, error) {
	var search []model.SearchContract

	query := fmt.Sprintf("SELECT id, status, requisites ->> 'beneficiary' AS  beneficiary,  contract_parameters ->> 'contract_number' AS contract_number," +
		"type AS contract_type,  created_at, updated_at, manager AS author, contract_parameters ->> 'contract_amount' AS amount FROM  contracts " +
		"WHERE  contract_parameters ->> 'contract_number' like  $1 AND status =  $2")

	err := db.GetDBConn().Raw(query, "%"+param+"%", status).Scan(&search).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return search, err
	}

	return search, nil

}

func SearchContractHistory(field string, param string) ([]model.SearchContract, error) {
	var search []model.SearchContract

	if field == "author" {
		query := fmt.Sprintf("SELECT id, requisites ->> 'beneficiary' AS  beneficiary,  contract_parameters ->> 'contract_number' AS contract_number," +
			"type AS contract_type,  created_at, updated_at, manager AS author, contract_parameters ->> 'contract_amount' AS amount FROM  contracts " +
			"WHERE  manager  like  $1")
		err := db.GetDBConn().Raw(query, "%"+param+"%").Scan(&search).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return search, err
		}
		return search, nil

	}
	//это чтобы понять из какого объекта будем доставать поля из JSONB
	var jsonBTable string
	if field == "contract_number" {
		jsonBTable = "contract_parameters"
	} else if field == "beneficiary" {
		jsonBTable = "requisites"

	}

	query := fmt.Sprintf("SELECT id, requisites ->> 'beneficiary' AS  beneficiary,  contract_parameters ->> 'contract_number' AS contract_number,"+
		"type AS contract_type,  created_at, updated_at, manager AS author, contract_parameters ->> 'contract_amount' AS amount FROM  contracts "+
		"WHERE  %s ->> $1 like  $2", jsonBTable)

	err := db.GetDBConn().Raw(query, field, "%"+param+"%").Scan(&search).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return search, err
	}

	return search, nil

}

func ChangeDataContract(id int) error {
	var endDate model.Date

	onWork := "в работе"

	chekingExist := fmt.Sprint("SELECT contract_parameters ->> 'end_date' AS end_date FROM contracts WHERE  id = $1")
	err := db.GetDBConn().Raw(chekingExist, id).Scan(&endDate).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	splitDate := strings.Split(endDate.EndDate, ".")
	if len(splitDate) != 3 {
		return errors.New("длина даты должны быть 3")
	}

	year, err := strconv.Atoi(splitDate[2])
	if err != nil {
		return err
	}
	year += 1
	strYear := strconv.Itoa(year)
	extendDate := fmt.Sprintf("%s.%s.%s", splitDate[0], splitDate[1], strYear)

	sqlUpdate := fmt.Sprint(`UPDATE contracts  SET contract_parameters = jsonb_set("contract_parameters", '{"extend_date"}', to_jsonb($1::text), true) WHERE id = $2 AND status = $3`)

	err = db.GetDBConn().Exec(sqlUpdate, extendDate, id, onWork).Error
	if err != nil {
		return err
	}

	return nil

}
