package repository

import (
	"admin_panel/db"
	"admin_panel/model"
	"fmt"
)

func GetAllCurrency() (currency []model.Currency, err error) {
	db.GetDBConn().Find(&currency)
	if currency == nil{
		return nil, err
	}

	return currency, err


}


func CreateMarketingContract(contract model.MarketingServicesContract) (model.MarketingServicesContract, error)  {
	// Структура Requisites
	req := contract.Requisites
	//struct supplier_company_manager
	SCM := contract.SupplierCompanyManager
	// Параметры Контракта
	paramСont := contract.ParamContract
	discount := contract.DiscountPercent
	//	sqlQuery := "INSERT INTO roles_rights (role_id, right_id) VALUES(?, ?)"
	sqlQueryRequisites := fmt.Sprintf("INSERT INTO %s (beneficiary, bank_of_beneficiary,  bin,  iic,  phone, account_number) " +
		"VALUES($1, $2, $3, $4, $5, $6) RETURNING id", "requisites")
	//db.GetDBConn().Exec(sqlQueryRequisites, req.Beneficiary, req.BankOfBeneficiary, req.BIN, req.IIC, req.Phone, req.AccountNumber)
	db.GetDBConn().Raw(sqlQueryRequisites, req.Beneficiary, req.BankOfBeneficiary, req.BIN, req.IIC, req.Phone, req.AccountNumber).Scan(&req.ID)

	// supplier_company_manager
	sqlReqSCM := fmt.Sprintf("INSERTRT INTO %s (work_phone, email, skype, phone, position, base)" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", "supplier_company_manager")

	//err := db.GetDBConn().Exec(sqlReqSCM, SCM.WorkPhone, SCM.Email, SCM.Skype, SCM.Phone, SCM.Position, SCM.Base).Error
	//if err != nil {
	//	return contract, err
	//}
	db.GetDBConn().Raw(sqlReqSCM, SCM.WorkPhone, SCM.Email, SCM.Skype, SCM.Phone, SCM.Position, SCM.Base).Scan(&SCM.ID)



	sqlReqPC := fmt.Sprintf("INSERT INTO %s (number_of_contract, amount_contract, currency_id,  prepayment, date_of_delivery, frequency_deferred_discount, delivery_address, return_time_delivery) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", "param_contract")


	db.GetDBConn().Raw(sqlReqPC, paramСont.NumberOfContract, paramСont.AmountContract,
		paramСont.CurrencyID, paramСont.Prepayment, paramСont.DateOfDelivery, paramСont.FrequencyDeferredDiscount,
		paramСont.DeliveryAddress, paramСont.ReturnTimeDelivery).Scan(&paramСont.ID)

	sqlDisc := fmt.Sprintf("INSERT INTO %s(name, amount) VALUES($1, $2) RETURNING id", "discount_percent")
	db.GetDBConn().Raw(sqlDisc, discount.Name, discount.Amount).Scan(&discount.ID)















	return contract, nil
}
