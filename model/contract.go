package model

import "time"

type MarketingServicesAgreement struct {
	Requisites             Requisites              `json:"requisites"`
	SupplierCompanyManager SupplierCompanyManager  `json:"supplier_company_manager"`
	Manager                string                  `json:"manager"`
	KAM                    string                  `json:"kam"`
	ParamContract          ContractParameters `json:"param_contract"`
}

// Ревезиты
type Requisites struct {
	Beneficiary       string `json:"beneficiary"`
	BankOfBeneficiary string `json:"bank_of_beneficiary"`
	BIN               string `json:"bin"`
	// индивидуальный идентификационный код
	IIC           string `json:"iic"`
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
	Base string `json:"base"`
}

type ContractParameters struct {
	NumberOfContract          string    `json:"number_of_contract"`
	AmountContract            int       `json:"amount_contract"`
	CurrencyContract          string    `json:"currency_contract"`
	Prepayment                int       `json:"prepayment"`
	DateOfDelivery            time.Time `json:"date_of_delivery"`
	FrequencyDeferredDiscount string `json:"frequency_deferred_discount"`
	DeliveryAddress           []string `json:"delivery_address"`
	//интервал времени поставки после поступления денежгых средств
	DeliveryTimeInterval string `json:"delivery_time_interval"`
	//время возврата при условии не поставки
	ReturnTimeDelivery int `json:"return_time_delivery"`

	//TODO:
	ProductNumber int    `json:"product_number"`
	Tradename     string `json:"tradename"`
	Price         int    `json:"price"`
	Currency      string `json:"currency"`
}

type DiscountPercent struct {
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	IsActive bool   `json:"is_active"`
}

type ContractsAttachments struct {
	AttachmentTemplate string `json:"attachment_template"`
	Applications       string `json:"applications"`
}
