package repository

import (
	"admin_panel/db"
	"admin_panel/model"
	"encoding/json"
	"log"
)

func GetAllCurrency() (currency []model.Currency, err error) {
	db.GetDBConn().Find(&currency)
	if currency == nil {
		return nil, err
	}

	return currency, err

}

func CreateMarketingContract(contract model.MarketingServicesContract)  error {
	// Структура Requisites
	req := contract.Requisites // DONE
	//struct supplier_company_manager
	SCM := contract.SupplierCompanyManager
	// Параметры Контракта
	paramСont := contract.ParamContract
	discount := contract.DiscountPercent
	products := contract.Products
	reqMarshall, err := json.Marshal(req)
	if err != nil {
		return err
	}
	paramСontMarshall, err := json.Marshal(paramСont)
	if err != nil {

	}

	SCMMarshall, err := json.Marshal(SCM)
	if err != nil {
		return err
	}

	productsMarshall, err := json.Marshal(products)
	if err != nil {
		return err
	}
	discountMarshall, err := json.Marshal(discount)
	if err != nil {
		return err
	}

	err = db.GetDBConn().Exec("INSERT INTO marketing_services_contract(requisites, supplier_company_manager, contract_parameters, products,discount_percent)"+
		"VALUES(?,?,?,?,?)", reqMarshall, SCMMarshall, paramСontMarshall,productsMarshall, discountMarshall).Error
	if err != nil {
		log.Println(err)
		return err

	}

	////	sqlQuery := "INSERT INTO roles_rights (role_id, right_id) VALUES(?, ?)"
	//sqlQueryRequisites := fmt.Sprintf("INSERT INTO %s (beneficiary, bank_of_beneficiary,  bin,  iic,  phone, account_number) " +
	//	"VALUES($1, $2, $3, $4, $5, $6) RETURNING id", "requisites")
	//
	//db.GetDBConn().Raw(sqlQueryRequisites, req.Beneficiary, req.BankOfBeneficiary, req.BIN, req.IIC, req.Phone, req.AccountNumber).Scan(&req.ID)
	//
	//// supplier_company_manager
	//sqlReqSCM := fmt.Sprintf("INSERTRT INTO %s (work_phone, email, skype, phone, position, base)" +
	//	"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", "supplier_company_manager")
	//
	////err := db.GetDBConn().Exec(sqlReqSCM, SCM.WorkPhone, SCM.Email, SCM.Skype, SCM.Phone, SCM.Position, SCM.Base).Error
	////if err != nil {
	////	return contract, err
	////}
	//db.GetDBConn().Raw(sqlReqSCM, SCM.WorkPhone, SCM.Email, SCM.Skype, SCM.Phone, SCM.Position, SCM.Base).Scan(&SCM.ID)
	//
	//
	//
	//sqlReqPC := fmt.Sprintf("INSERT INTO %s (number_of_contract, amount_contract, currency_id,  prepayment, date_of_delivery, frequency_deferred_discount, delivery_address, return_time_delivery) " +
	//	"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", "param_contract")
	//
	//
	//db.GetDBConn().Raw(sqlReqPC, paramСont.NumberOfContract, paramСont.AmountContract,
	//	paramСont.CurrencyID, paramСont.Prepayment, paramСont.DateOfDelivery, paramСont.FrequencyDeferredDiscount,
	//	paramСont.DeliveryAddress, paramСont.ReturnTimeDelivery).Scan(&paramСont.ID)
	//
	//
	//
	//for _, dis := range discount{
	//	sqlDisc := fmt.Sprintf("INSERT INTO %s(type, name, discount_amount, grace_days, payment_multiplicity, amount, comments) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id", "discount_percent")
	//	db.GetDBConn().Raw(sqlDisc, dis.Type, dis.Name, dis.DiscountAmount, dis.GraceDays, dis.PaymentMultiplicity, dis.Amount, dis.Comments).Scan(&dis.ID)
	//
	//}
	//
	//
	//for _, product := range products{
	//	sqlProd := fmt.Sprintf("INSERT INTO %s (product_number, price, currency) VALUES ($1, $2, $3)", "products")
	//
	//
	//	db.GetDBConn().Raw(sqlProd, product.ProductNumber, product.Price, product.Currency).Scan(&product.ID)
	//}


	return  nil
}