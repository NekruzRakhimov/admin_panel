package model

import (
	"github.com/lib/pq"
	"time"
)

<<<<<<< HEAD
type MarketingServicesAgreement struct {
	Requisites             Requisites              `json:"requisites"`
	SupplierCompanyManager SupplierCompanyManager  `json:"supplier_company_manager"`
	Manager                string                  `json:"manager"`
	KAM                    string                  `json:"kam"`
	ParamContract          ContractParameters `json:"param_contract"`
=======
type MarketingServicesContract struct {
	ID                     int                    `json:"id"`
	Requisites             Requisites             `json:"requisites"`
	SupplierCompanyManager SupplierCompanyManager `json:"supplier_company_manager"`
	Manager                string                 `json:"manager"`
	KAM                    string                 `json:"kam"`
	ParamContract          ContractParams         `json:"param_contract"`
	DiscountPercent        []DiscountPercent      `json:"discount_percent"`
	Products               []Product              `json:"products"`
>>>>>>> d28815d5cb5b9b046088904e60cacff419744789
}

// Ревезиты
type Requisites struct {
	ID                int    `json:"id"`
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
	ID        int    `json:"id"`
	WorkPhone string `json:"work_phone"`
	Email     string `json:"email"`
	Skype     string `json:"skype"`
	Phone     string `json:"phone"`
	// помоему  в этом случае ему нужен слайс стрингов
	Position string `json:"position"`
	// Основание - надо понять как правильно назвать это
	Base string `json:"base"`
}

<<<<<<< HEAD
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
=======
//TODO: поменять запрос в репозитории Маркетинговых договорах
type ContractParams struct {
	ID               int    `json:"id"`
	NumberOfContract string `json:"number_of_contract"`
	AmountContract   int    `json:"amount_contract"`
	CurrencyID       int    `json:"currency_id"`

	Prepayment     int       `json:"prepayment"`
	DateOfDelivery time.Time `json:"date_of_delivery"`
	//Кратность расчета отложенной скидки TODO: возможно нужно поменять
	FrequencyDeferredDiscount string `json:"frequency_deferred_discount"`

	DeliveryAddress pq.StringArray `json:"delivery_address"`
	//интервал времени поставки после поступления денежгых средств
	DeliveryTimeInterval int `json:"delivery_time_interval"`
	//время возврата при условии не поставки
	ReturnTimeDelivery int       `json:"return_time_delivery"`
	ContractDate       time.Time `json:"contract_date"`
}

type DiscountPercent struct {
	ID                  int    `json:"id"`
	Type                string `json:"type"`
	Name                string `json:"name"`
	DiscountAmount      int    `json:"discount_amount"`
	GraceDays           string `json:"grace_days"`
	PaymentMultiplicity string `json:"payment_multiplicity"`
	IsActive            bool   `json:"is_active"`
	Amount              int    `json:"amount"`
	Comments            string `json:"comments"`
}

type Product struct {
	ID            int `json:"id"`
	ProductNumber string `json:"product_number"`
	Price         float32 `json:"price"`
	Currency      string  `json:"currency"`
}

type Currency struct {
	ID        int       `json:"id"`
	Alpha3    string    `json:"alpha_3"`
	Symbol    string    `json:"symbol"`
	Name      string    `json:"name"`
	ImageName string    `json:"image_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	IsRemoved bool      `json:"is_removed"`
>>>>>>> d28815d5cb5b9b046088904e60cacff419744789
}
