package model

import (
	"time"
)

type MarketingServicesContract struct {
	ID                     int                    `json:"id,omitempty"`
	Requisites             Requisites             `json:"requisites,omitempty"`
	SupplierCompanyManager SupplierCompanyManager `json:"supplier_company_manager,omitempty"`

	ParamContract          ContractParams         `json:"param_contract,omitempty"`
	DiscountPercent        []DiscountPercent      `json:"discount_percent,omitempty"`
	Products               []Product              `json:"products,omitempty"`
}

// Ревезиты
type Requisites struct {
	//ID                int    `json:"id"`
	Beneficiary       string `json:"beneficiary,omitempty,omitempty"`
	BankOfBeneficiary string `json:"bank_of_beneficiary,omitempty"`
	BIN               string `json:"bin,omitempty"`
	// индивидуальный идентификационный код
	IIC           string `json:"iic,omitempty"`
	Phone         string `json:"phone,omitempty"`
	AccountNumber string `json:"account_number, omitempty"`
	Manager       string `json:"manager, omitempty"`
	KAM string `json:"kam,omitempty"`

}

// Руководитель компании поставщика
type SupplierCompanyManager struct {
	//ID        int    `json:"id"`
	WorkPhone string `json:"work_phone,omitempty"`
		Email     string `json:"email,omitempty"`
		Skype     string `json:"skype,omitempty"`
	Phone     string `json:"phone,omitempty"`
	// помоему  в этом случае ему нужен слайс стрингов
	Position string `json:"position,omitempty"`
	// Основание - надо понять как правильно назвать это
		Base string `json:"base,omitempty"`
}

//type ContractParameters struct {
//		NumberOfContract          string    `json:"number_of_contract"`
	//AmountContract            int       `json:"amount_contract"`
	//CurrencyContract          string    `json:"currency_contract"`
	//Prepayment                int       `json:"prepayment"`
	//DateOfDelivery            time.Time `json:"date_of_delivery"`
	//FrequencyDeferredDiscount string    `json:"frequency_deferred_discount"`
	//DeliveryAddress           []string  `json:"delivery_address"`
	////интервал времени поставки после поступления денежгых средств
	//DeliveryTimeInterval string `json:"delivery_time_interval"`
	////время возврата при условии не поставки
	//ReturnTimeDelivery int `json:"return_time_delivery"`
	//
	////TODO:
	//ProductNumber int    `json:"product_number"`
	//Tradename     string `json:"tradename"`
	//Price         int    `json:"price"`
	//Currency      string `json:"currency"`
//}

//type DiscountPercent struct {
//	Name     string `json:"name"`
//	Amount   int    `json:"amount"`
//	IsActive bool   `json:"is_active"`
//}

type ContractsAttachments struct {
	AttachmentTemplate string `json:"attachment_template"`
	Applications       string `json:"applications"`
}

//TODO: поменять запрос в репозитории Маркетинговых договорах

type ContractParams struct {
	NumberOfContract string `json:"number_of_contract,omitempty"`
	AmountContract   int    `json:"amount_contract,omitempty"`
	//CurrencyID       int    `json:"currency_id,omitempty"`
	Currency       string    `json:"currency,omitempty"`

	Prepayment     int       `json:"prepayment,omitempty"`
	DateOfDelivery time.Time `json:"date_of_delivery,omitempty"`
	//Кратность расчета отложенной скидки TODO: возможно нужно поменять
	FrequencyDeferredDiscount string `json:"frequency_deferred_discount,omitempty"`

	//DeliveryAddress pq.StringArray `json:"delivery_address,omitempty"`
	DeliveryAddress []string `json:"delivery_address,omitempty"`
	//интервал времени поставки после поступления денежгых средств
	DeliveryTimeInterval int `json:"delivery_time_interval,omitempty"`
	//время возврата при условии не поставки
	ReturnTimeDelivery int       `json:"return_time_delivery,omitempty"`
	ContractDate       time.Time `json:"contract_date,omitempty"`
}

type DiscountPercent struct {
	Type                string `json:"type,omitempty"`
	Name                string `json:"name,omitempty"`
	DiscountAmount      int    `json:"discount_amount,omitempty"`
	GraceDays           string `json:"grace_days,omitempty"`
	PaymentMultiplicity string `json:"payment_multiplicity,omitempty"`
	IsActive            bool   `json:"is_active,omitempty"`
	Amount              int    `json:"amount,omitempty"`
	Comments            string `json:"comments,omitempty"`
}

type Product struct {
	ProductNumber string  `json:"product_number,omitempty"`
	Price         float32 `json:"price,omitempty"`
	Currency      string  `json:"currency,omitempty"`
}

type Currency struct {
	ID        int       `json:"id"`
	Alpha3    string    `json:"alpha_3,omitempty"`
	Symbol    string    `json:"symbol,omitempty"`
	Name      string    `json:"name,omitempty"`
	ImageName string    `json:"image_name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	IsRemoved bool      `json:"is_removed"`
}
