package repository

import (
	"admin_panel/db"
	"admin_panel/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"strings"
)

func CreateContract(contractWithJson models.ContractWithJsonB) error {
	//fmt.Printf(">>>> %+v", contractWithJson)
	if err := db.GetDBConn().Table("contracts").Exec("UPDATE contracts SET status = ? WHERE id = ?", "заверщённый", contractWithJson.PrevContractId).Error; err != nil {
		return err
	}

	err := db.GetDBConn().Table("contracts").Omit("created_at", "updated_at", "is_extend_contract", "extend_date", "brand_name", "brand_code", "discount_percent", "contract_id").Create(&contractWithJson).Error
	fmt.Println(contractWithJson.ID, "ContractParam")

	if err != nil {
		log.Println("[repository.CreateContract]|[db.GetDBConn().Table(\"contracts\").Create(&contractWithJson).Error]| error is: ", err.Error())
		return err
	}

	//TODO: Нурсу сказать, чтобы он убрал
	for _, value := range contractWithJson.DiscountBrand {
		err := db.GetDBConn().Exec("INSERT INTO brands(brand, brand_code, discount_percent, contract_id) VALUES ($1, $2, $3, $4)", value.BrandName, value.BrandCode, value.DiscountPercent, contractWithJson.ID).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func EditContract(contractWithJson models.ContractWithJsonB) error {
	if err := db.GetDBConn().Table("contracts").Omit("created_at", "updated_at", "is_extend_contract", "extend_date", "contract_id").Save(&contractWithJson).Error; err != nil {
		log.Println("[repository.EditContract]|[db.GetDBConn().Table(\"contracts\").Save(&contractWithJson).Error]| error is: ", err.Error())
		return err
	}

	if len(contractWithJson.Discounts) > 0 {
		if err := db.GetDBConn().Exec("DELETE FROM brands WHERE contract_id = ?", contractWithJson.ID).Error; err != nil {
			return err
		}
	}

	for _, value := range contractWithJson.DiscountBrand {
		err := db.GetDBConn().Exec("INSERT INTO brands(brand, brand_code, discount_percent, contract_id) VALUES ($1, $2, $3, $4)", value.BrandName, value.BrandCode, value.DiscountPercent, contractWithJson.ID).Error
		if err != nil {
			return err
		}
	}

	if err := RecordContractStatusChange(contractWithJson.ID, contractWithJson.Status); err != nil {
		return err
	}

	return nil
}

func GetAllContracts(contractStatus string) (contracts []models.ContractWithJsonB, err error) {
	fmt.Println("GetALlContract Calling---------------------------")

	var contractStatusRus = ""
	sqlQuery := "SELECT * FROM contracts WHERE id not in (select prev_contract_id from contracts) AND is_active = true"

	if contractStatus != "" && contractStatus != "ACTIVE_AND_EXPIRED" {
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

	if contractStatus == "ACTIVE_AND_EXPIRED" {
		sqlQuery += fmt.Sprintf(" AND status in ('%s', '%s')", "в работе", "заверщённый")
	}

	//sqlQuery += " ORDER BY created_at DESC"
	sqlQuery += " ORDER BY id desc"

	if err := db.GetDBConn().Raw(sqlQuery).Scan(&contracts).Error; err != nil {
		log.Println("[repository.GetAllContracts]|[db.GetDBConn().Raw(sqlQuery).Scan(&contracts).Error]| error is: ", err.Error())
		return nil, err
	}

	return contracts, nil
}

func GetContractDetails(contractId int) (contract models.ContractWithJsonB, err error) {
	contract.ID = contractId
	var brands []models.DiscountBrand
	if err := db.GetDBConn().Table("contracts").Find(&contract).Error; err != nil {
		return models.ContractWithJsonB{}, err
	}

	if err = db.GetDBConn().Raw("SELECT id, brand as brand_name, brand_code, discount_percent FROM  brands  WHERE  contract_id = ?", contract.ID).Scan(&brands).Error; err != nil {
		return models.ContractWithJsonB{}, err
	}
	log.Println("BRANDS", brands)
	contract.DiscountBrand = brands

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

func GetContractStatusChangesHistory(contractId int) (history []models.ContractStatusHistory, err error) {
	sqlQuery := `SELECT id,
      		 contract_id,
       		status,
       		to_char(created_at::date, 'DD.MM.YYYY'),
       		author
		FROM status_changes_history WHERE contract_id = ?`
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

	fmt.Println(contractId, contractCode, "ДАННЫЕ с 1С")
	if err := db.GetDBConn().Exec("UPDATE  contracts set ext_contract_code = $1  WHERE id = $2", contractCode, contractId).Error; err != nil {
		return err
	}

	return nil
}

func SaveContractExternalCodeByBIN(contractFor1C models.ContractDTOFor1C, contractCode string) error {
	if err := db.GetDBConn().Exec("UPDATE  contracts set ext_contract_code = $1  WHERE requisites ->> 'bin' = $2 AND contract_parameters ->> 'contract_number' = $3", contractCode, contractFor1C.Requisites.BIN, contractFor1C.ContractParameters.ContractNumber).Error; err != nil {
		return err
	}

	return nil
}

func SearchContractByNumber(param string, status string) ([]models.SearchContract, error) {
	var search []models.SearchContract
	query := fmt.Sprint("SELECT id, status, requisites ->> 'beneficiary' AS  beneficiary,  contract_parameters ->> 'contract_number' AS contract_number," +
		"type AS contract_type,  created_at, updated_at, manager AS author, contract_parameters ->> 'contract_amount' AS amount FROM  contracts " +
		"WHERE  contract_parameters ->> 'contract_number' like  $1 ")
	if status == "" {
		err := db.GetDBConn().Raw(query, "%"+param+"%").Scan(&search).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return search, err
		}

	}
	if status != "" {
		query += fmt.Sprintf("AND status =  $2")
		err := db.GetDBConn().Raw(query, "%"+param+"%", status).Scan(&search).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return search, err
		}

	}

	for i := range search {
		if search[i].ContractType == "supply" {
			search[i].ContractType = "Договор поставок"
		} else if search[i].ContractType == "marketing_services" {
			search[i].ContractType = "Договор маркетинговых услуг"
		}

	}

	return search, nil

}

func SearchContractHistory(field string, param string) ([]models.SearchContract, error) {
	var search []models.SearchContract

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
	var endDate models.Date

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

func SearchHistoryExecution(field string, param string) ([]models.SearchContract, error) {
	var search []models.SearchContract
	if field == "author" {
		query := fmt.Sprintf("SELECT id, manager AS author, status," +
			"created_at, contract_parameters ->> 'end_date' AS end_date, comment FROM  contracts " +
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
	query := fmt.Sprintf("SELECT id, manager AS author, status,"+
		"created_at, contract_parameters ->> 'end_date' AS end_date, comment FROM  contracts"+
		"WHERE  %s ->> $1 like  $2", jsonBTable)

	err := db.GetDBConn().Raw(query, field, "%"+param+"%").Scan(&search).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return search, err
	}

	return search, nil

}

func SaveSuppliers(suppliers []models.DataClient) error {
	for _, supplier := range suppliers {
		if err := db.GetDBConn().Table("suppliers").Create(&supplier).Error; err != nil {
			return err
		}

	}

	return nil
}

func GetSuppliers() (suppliers []models.DataClient, err error) {
	sqlQuery := "SELECT * FROM suppliers"
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&suppliers).Error; err != nil {
		return nil, err
	}

	return suppliers, nil
}

func GetSuppliersByParameter(filed string, value string)  (suppliers []models.DataClient, err error) {
	//query := fmt.Sprintf("SELECT *FROM stored_reports WHERE %s LIKE $1", field)
	sqlQuery := fmt.Sprintf("SELECT *FROM suppliers WHERE  %s  LIKE ?", filed)
	//sqlQuery := "SELECT * FROM suppliers WHERE  client_name $1"

	err  = db.GetDBConn().Raw(sqlQuery, "%"+value+"%").Scan(&suppliers).Error
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}
