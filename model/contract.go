package model

type ContractWithJsonB struct {
	ID                        int             `json:"id"`
	Type                      string          `json:"type"`
	PrevContractId            int             `json:"-"`
	Status                    string          `json:"status"` //вынести статус в отдельную таблицу
	Requisites                string          `json:"requisites"`
	Manager                   string          `json:"manager"`
	KAM                       string          `json:"kam"`
	SupplierCompanyManager    string          `json:"supplier_company_manager"`
	ContractParameters        string          `json:"contract_parameters"`
	WithTemperatureConditions bool            `json:"with_temperature_conditions"`
	Products                  string          `json:"products"`
	Discounts                 string          `json:"discounts"`
	Comment                   string          `json:"comment"`
	CreatedAt                 string          `json:"created_at,omitempty"`
	UpdatedAt                 string          `json:"updated_at,omitempty"`
	IsIndivid                 bool            `json:"is_individ"`
	IsExtendContract          bool            `json:"is_extend_contract"`
	ExtendDate                string          `json:"extend_date"`
	DiscountBrand             []DiscountBrand `json:"discount_brand"`

}

type Contract struct {
	ID                        int                    `json:"id"`
	Type                      string                 `json:"type"`
	PrevContractId            int                    `json:"-" gorm:"-"`
	Status                    string                 `json:"status"`
	Requisites                Requisites             `json:"requisites"`
	Manager                   string                 `json:"manager,omitempty"`
	KAM                       string                 `json:"kam,omitempty"`
	SupplierCompanyManager    SupplierCompanyManager `json:"supplier_company_manager"`
	ContractParameters        ContractParameters     `json:"contract_parameters"`
	WithTemperatureConditions bool                   `json:"with_temperature_conditions"`
	Products                  []Product              `json:"products"`
	Discounts                 []Discount             `json:"discounts"`
	Comment                   string                 `json:"comment"`
	CreatedAt                 string                 `json:"created_at,omitempty"`
	UpdatedAt                 string                 `json:"updated_at,omitempty"`
	IsExtendContract          bool                   `json:"is_extend_contract"`
	ExtendDate                string                 `json:"extend_date"`

	IsIndivid bool `json:"is_individ"`
	//	Brand           string `json:"brand"`
	//	DiscountPercent string `json:"discount_percent"`
	DiscountBrand []DiscountBrand `json:"discount_brand"`

}

type ContractDTOFor1C struct {
	ID                        int                        `json:"id"`
	Type                      string                     `json:"type"`
	PrevContractId            int                        `json:"-" gorm:"-"`
	Status                    string                     `json:"status"`
	Requisites                Requisites                 `json:"requisites"`
	Manager                   string                     `json:"manager,omitempty"`
	Country                   string                     `json:"country"`
	KAM                       string                     `json:"kam,omitempty"`
	SupplierCompanyManager    SupplierCompanyManager     `json:"supplier_company_manager"`
	ContractParameters        ContractParametersDTOFor1C `json:"contract_parameters"`
	WithTemperatureConditions bool                       `json:"with_temperature_conditions"`
	Products                  []Product                  `json:"products"`
	Discounts                 []Discount                 `json:"discounts"`
	Comment                   string                     `json:"comment"`
	CreatedAt                 string                     `json:"created_at,omitempty"`
	UpdatedAt                 string                     `json:"updated_at,omitempty"`
}

type ContractMiniInfo struct {
	ID               int     `json:"id"`
	PrevContractId   int     `json:"-" gorm:"-"`
	ContractorName   string  `json:"contractor_name"`
	Beneficiary      string  `json:"beneficiary,omitempty"`
	ContractNumber   string  `json:"contract_number"`
	ContractType     string  `json:"contract_type"`
	Status           string  `json:"status"`
	Author           string  `json:"author"`
	Amount           float32 `json:"amount"`
	CreatedAt        string  `json:"created_at,omitempty"`
	UpdatedAt        string  `json:"updated_at,omitempty"`
	IsExtendContract bool    `json:"is_extend_contract"`
	ExtendDate       string  `json:"extend_date"`
}

// Requisites Ревезиты
type Requisites struct {
	ContractorName         string `json:"contractor_name"`
	Beneficiary            string `json:"beneficiary,omitempty"`
	BankOfBeneficiary      string `json:"bank_of_beneficiary,omitempty"`
	BankBeneficiaryAddress string `json:"bank_beneficiary_address"`
	SwiftCode              string `json:"swift_code"`
	BIN                    string `json:"bin,omitempty"`
	IIC                    string `json:"iic,omitempty"`
	Phone                  string `json:"phone,omitempty"`
	AccountNumber          string `json:"account_number,omitempty"`
}

// SupplierCompanyManager Руководитель компании поставщика
type SupplierCompanyManager struct {
	Country   string `json:"country"`
	WorkPhone string `json:"work_phone,omitempty"`
	Email     string `json:"email,omitempty"`
	Skype     string `json:"skype,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Position  string `json:"position,omitempty"`
	// Основание - надо понять как правильно назвать это
	Base     string `json:"base,omitempty"`
	FullName string `json:"full_name"`
}

type ContractParameters struct {
	ContractNumber            string   `json:"contract_number"`
	ContractAmount            float32  `json:"contract_amount"`
	Currency                  string   `json:"currency,omitempty"`
	Prepayment                float32  `json:"prepayment,omitempty"`
	DateOfDelivery            string   `json:"date_of_delivery,omitempty"`
	FrequencyDeferredDiscount string   `json:"frequency_deferred_discount,omitempty"` //Кратность расчета отложенной скидки TODO: возможно нужно поменять
	DeliveryAddress           []string `json:"delivery_address,omitempty"`
	DeliveryTimeInterval      int      `json:"delivery_time_interval,omitempty"` //интервал времени поставки после поступления денежых средств
	ReturnTimeDelivery        int      `json:"return_time_delivery,omitempty"`   //время возврата при условии не поставки
	ContractDate              string   `json:"contract_date,omitempty"`
	StartDate                 string   `json:"start_date"`
	EndDate                   string   `json:"end_date"`
	IsExtendContract          bool     `json:"is_extend_contract"`
	ExtendDate                string   `json:"extend_date"`
}

type ContractParametersDTOFor1C struct {
	ContractNumber            string  `json:"contract_number"`
	ContractAmount            float32 `json:"contract_amount"`
	Currency                  string  `json:"currency,omitempty"`
	Prepayment                float32 `json:"prepayment,omitempty"`
	DateOfDelivery            string  `json:"date_of_delivery,omitempty"`
	FrequencyDeferredDiscount string  `json:"frequency_deferred_discount,omitempty"` //Кратность расчета отложенной скидки TODO: возможно нужно поменять
	DeliveryAddress           string  `json:"delivery_address,omitempty"`
	DeliveryTimeInterval      int     `json:"delivery_time_interval,omitempty"` //интервал времени поставки после поступления денежых средств
	ReturnTimeDelivery        int     `json:"return_time_delivery,omitempty"`   //время возврата при условии не поставки
	PriceType                 string  `json:"price_type,omitempty"`
	StartDate                 string  `json:"start_date,omitempty"`
	EndDate                   string  `json:"end_date,omitempty"`
}

type Product struct {
	ProductNumber    string  `json:"product_number,omitempty"`
	ProductName      string  `json:"product_name"`
	Price            float64 `json:"price,omitempty"`
	Currency         string  `json:"currency,omitempty"`
	Substance        string  `json:"substance"`
	StorageCondition string  `json:"storage_condition"`
	Producer         string  `json:"producer"`
}

type Discount struct {
	Name    string           `json:"name,omitempty"`
	Code    string           `json:"code"`
	Periods []DiscountPeriod `json:"periods"`
}

type DiscountPeriod struct {
	PeriodFrom   string `json:"period_from"`
	PeriodTo     string `json:"period_to"`
	TotalAmount  int    `json:"total_amount"`
	RewardAmount int    `json:"reward_amount"`
	//DiscountAmount      float32 `json:"discount_amount,omitempty"`
	//GraceDays           string  `json:"grace_days,omitempty"`
	//PaymentMultiplicity string  `json:"payment_multiplicity,omitempty"`
	//Amount              float32 `json:"amount,omitempty"`
	//Site                string  `json:"site,omitempty"`
	//Other               string  `json:"other"`
	//Comments            string  `json:"comments,omitempty"`
}

//Discount struct {
//	//Type                string `json:"type,omitempty"`
//	Name                string  `json:"name,omitempty"`
//	DiscountAmount      float32 `json:"discount_amount,omitempty"`
//	GraceDays           string  `json:"grace_days,omitempty"`
//	PaymentMultiplicity string  `json:"payment_multiplicity,omitempty"`
//	Amount              float32 `json:"amount,omitempty"`
//	Site                string  `json:"site,omitempty"`
//	Other               string  `json:"other"`
//	Comments            string  `json:"comments,omitempty"`
//}

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

//type Currency struct {
//	ID        int       `json:"id"`
//	Alpha3    string    `json:"alpha_3,omitempty"`
//	Symbol    string    `json:"symbol,omitempty"`
//	Name      string    `json:"name,omitempty"`
//	ImageName string    `json:"image_name,omitempty"`
//	CreatedAt time.Time `json:"created_at"`
//	UpdatedAt time.Time `json:"updated_at"`
//	DeletedAt time.Time `json:"deleted_at"`
//	IsRemoved bool      `json:"is_removed"`
//}

type MarketingServicesContract struct {
	ID                     int                    `json:"id,omitempty"`
	Requisites             Requisites             `json:"requisites,omitempty"`
	SupplierCompanyManager SupplierCompanyManager `json:"supplier_company_manager,omitempty"`

	ParamContract   ContractParameters `json:"param_contract,omitempty"`
	DiscountPercent []Discount         `json:"discount_percent,omitempty"`
	Products        []Product          `json:"products,omitempty"`
}

type Currency struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type FrequencyDeferredDiscount struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Address struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description,omitempty"`
}

type Position struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ContractStatusHistory struct {
	ID        int    `json:"id"`
	Author    string `json:"author"`
	Status    string `json:"status"`
	StartDate string `json:"start_date" gorm:"created_at"`
	EndDate   string `json:"end_date"`
}

type Client struct {
	Name          string  `json:"name,omitempty"`
	Bank          string  `json:"bank,omitempty"`
	AccountNumber string  `json:"account_number,omitempty"`
	Id1C          string  `json:"id_1C,omitempty"`
	Bin           string  `json:"bin,omitempty"`
	Country       string  `json:"country,omitempty"`
	Reason        *string `json:"reason,omitempty"`
}

type Date struct {
	EndDate string `json:"end_date"`
}
