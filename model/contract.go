package model

import "time"

type MarketingServicesAgreement struct {
	Requisites Requisites `json:"requisites"`
	SupplierCompanyManager SupplierCompanyManager `json:"supplier_company_manager"`
	Manager string `json:"manager"`
	KAM string `json:"kam"`
	ParamContract ParametersOfTheContract `json:"param_contract"`



}

// Ревезиты
type Requisites struct {
	Beneficiary       string `json:"beneficiary"`
	BankOfBeneficiary string `json:"bank_of_beneficiary"`
	BIN               int64  `json:"bin"`
	// индивидуальный идентификационный код
	IIC           int64  `json:"iic"`
	Phone         string `json:"phone"`
	AccountNumber string `json:"account_number"`
}

// Руководитель компании поставщика
type SupplierCompanyManager struct {
	WorkPhone string `json:"work_phone"`
	Email     string `json:"email"`
	Skype     string `json:"skype"`
	Phone     string `json:"phone"`
	// помоему  в этом случае ему нужен слайс стрингов
	Position string `json:"position"`
	// Основание - надо понять как правильно назвать это
	Base     string `json:"base"`
}

type ParametersOfTheContract struct {
	NumberOfContract string `json:"number_of_contract"`
	AmountContract int `json:"amount_contract"`
	CurrencyContract string `json:"currency_contract"`
	Prepayment int `json:"prepayment"`
	DateOfDelivery time.Time `json:"date_of_delivery"`
	DeliveryAddress string `json:"delivery_address"`
	//интервал времени поставки после поступления денежгых средств
	DeliveryTimeInterval int
	//время возврата при условии не поставки
	ReturnTimeDelivery int `json:"return_time_delivery"`
}



type DiscountPercent struct {
	Name string `json:"name"`
	Amount int `json:"amount"`
	IsActive bool `json:"is_active"`
}





type  Currency struct {
	ID int `json:"id"`
	Alpha3 string `json:"alpha_3"`
	Symbol string `json:"symbol"`
	Name string `json:"name"`
	ImageName string `json:"image_name"`
	CreatedAt time.Time `json:"created_at"`

}