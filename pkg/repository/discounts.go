package repository

import (
	"admin_panel/db"
	"admin_panel/models"
	"encoding/json"
	"log"
)

func SaveDataFrom1C(block models.Block) error {
	blockJson, err := json.Marshal(&block)
	if err != nil {
		return err
	}

	sqlQuery := "INSERT INTO sales_batch (content) VALUES($1)"
	if err := db.GetDBConn().Exec(sqlQuery, blockJson).Error; err != nil {
		return err
	}

	return nil
}

func SaveDoubtedDiscounts(bin, periodFrom, periodTo, contractNumber string, discounts []models.DoubtedDiscountDetails) error {
	for _, discount := range discounts {
		if err := AddOrUpdateDoubtedDiscount(bin, periodFrom, periodTo, contractNumber, discount.Code, discount.Name, discount.IsCompleted); err != nil {
			log.Println("[repository.SaveDoubtedDiscounts]|[repository.AddOrUpdateDoubtedDiscount] error is: ", err.Error())
			return err
		}
	}

	return nil
}

func AddOrUpdateDoubtedDiscount(bin, periodFrom, periodTo, contractNumber, discountCode, DiscountName string, isCompleted bool) error {
	sqlQuery := "UPDATE doubted_discounts SET is_completed = ? WHERE bin = ? AND contract_number = ? AND code = ? AND period_from = ? AND period_to = ?"
	result := db.GetDBConn().Exec(sqlQuery, isCompleted, bin, contractNumber, discountCode, periodFrom, periodTo)
	if result.RowsAffected == 0 {
		sqlQuery = "INSERT INTO doubted_discounts (code, name, bin, contract_number, period_from, period_to) VALUES (?, ?, ?, ?, ?, ?)"
		if err := db.GetDBConn().Exec(sqlQuery, discountCode, DiscountName, bin, contractNumber, periodFrom, periodTo).Error; err != nil {
			return err
		}
	}

	return nil
}

func GetAllContractDetailByBIN(clientCode, PeriodFrom, PeriodTo string) (contracts []models.ContractWithJsonB, err error) {
	sqlQuery := `SELECT *
          FROM (SELECT *
          FROM "contracts"
          WHERE (requisites ->> 'client_code' = ? AND
          status = 'в работе')) as sub_query
          WHERE to_date(contract_parameters ->> 'start_date', 'DD.MM.YYYY') <= to_date(?, 'DD.MM.YYYY')
          AND to_date(contract_parameters ->> 'end_date', 'DD.MM.YYYY') >= to_date(?, 'DD.MM.YYYY')`

	if err = db.GetDBConn().Raw(sqlQuery, clientCode, PeriodFrom, PeriodTo).Scan(&contracts).Error; err != nil {
		log.Println("[repository][GetAllContractDetailByBIN] error is: ", err.Error())
		return nil, err
	}

	//var brands []models.DiscountBrand
	for i, contract := range contracts {
		if err = db.GetDBConn().Raw("SELECT id, brand as brand_name, brand_code, discount_percent FROM  brands  WHERE  contract_id = ?", contract.ID).Scan(&contracts[i].DiscountBrand).Error; err != nil {
			return nil, err
		}

	}

	return contracts, nil
}

func GetDiscountPeriod(bin string) ([]models.Discount, error) {
	//var discounts []models.Discount
	var discount []models.Discount

	//db.GetDBConn().Raw("SELECT jsonb_array_elements(discounts) FROM contracts WHERE  requisites ->> bin = $1", bin).Scan(&discounts)
	err := db.GetDBConn().Raw("SELECT discounts::text as discount FROM contracts WHERE requisites ->> 'bin' = $1", bin).Scan(&discount).Error
	if err != nil {
		return nil, err
	}

	return discount, nil

}

func DoubtedDiscountExecutionCheck(request models.RBRequest, contractNumber, discountCode string) (isCompleted bool) {
	var isCompletedArr []bool
	sqlQuery := "SELECT is_completed FROM doubted_discounts WHERE bin = ? AND contract_number = ? AND code = ? AND period_from = ? AND period_to = ?"
	_ = db.GetDBConn().Raw(sqlQuery, request.BIN, contractNumber, discountCode, request.PeriodFrom, request.PeriodTo).Pluck("is_completed", &isCompletedArr)

	if len(isCompletedArr) > 0 {
		isCompleted = isCompletedArr[0]
	}

	return isCompleted
}
