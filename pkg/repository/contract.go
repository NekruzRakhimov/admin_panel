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
	req := model.MarketingServicesContract{}.Requisites
	//struct supplier_company_manager
	SCM := model.MarketingServicesContract{}.SupplierCompanyManager
	// Параметры Контракта
	paramСont := model.MarketingServicesContract{}.ParamContract



	//	sqlQuery := "INSERT INTO roles_rights (role_id, right_id) VALUES(?, ?)"
	sqlQueryRequisites := fmt.Sprintf("INSERT INTO %s (beneficiary, bank_of_beneficiary,  bin,  iic,  phone, account_number) " +
		"VALUES($1, $2, $3, $4, $5, $6)", "requisites")
	db.GetDBConn().Exec(sqlQueryRequisites, req.Beneficiary, req.BankOfBeneficiary, req.BIN, req.IIC, req.Phone, req.AccountNumber)


	// supplier_company_manager
	sqlReqSCM := fmt.Sprintf("INSERTRT INTO %s (work_phone, email, skype, phone, position, base)" +
		" VALUES ($1, $2, $3, $4, $5, $6, %7)", "supplier_company_manager")

	err := db.GetDBConn().Exec(sqlReqSCM, SCM.WorkPhone, SCM.Email, SCM.Skype, SCM.Phone, SCM.Position, SCM.Base).Error
	if err != nil {
		return contract, err
	}


	sqlReqPC := fmt.Sprintf("INSERT INTO %s (number_of_contract, amount_contract, currency_contract,  prepayment, date_of_delivery, frequency_deferred_discount, delivery_address, return_time_delivery) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", "param_contract")

	db.GetDBConn().Exec(sqlReqPC, paramСont.NumberOfContract, paramСont.AmountContract,
		paramСont.CurrencyContract, paramСont.Prepayment, paramСont.DateOfDelivery, paramСont.FrequencyDeferredDiscount,
		paramСont.DeliveryAddress, paramСont.ReturnTimeDelivery )





	return contract, nil
}
