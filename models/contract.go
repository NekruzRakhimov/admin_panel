package models

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
	AdditionalAgreementNumber int             `json:"additional_agreement_number"`
	ExtContractCode           string          `json:"ext_contract_code"`
	View                      string          `json:"view"`
	Regions                   string          `json:"regions"`
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
	AdditionalAgreementNumber int                    `json:"additional_agreement_number"`
	IsIndivid                 bool                   `json:"is_individ"`
	//	Brand           string `json:"brand"`
	//	DiscountPercent string `json:"discount_percent"`
	DiscountBrand   []DiscountBrand `json:"discount_brand"`
	ExtContractCode string          `json:"ext_contract_code"`
	View            string          `json:"view"`
	Regions         []Regions       `json:"regions"`
}

type PriceType struct {
	PriceTypeName     string `json:"pricetype_name"`
	PriceTypeCode     string `json:"pricetype_code"`
	PriceTypeCurrency string `json:"pricetype_currency"`
	ClientBin         string `json:"client_bin"`
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
	ID                        int     `json:"id"`
	PrevContractId            int     `json:"-" gorm:"-"`
	ContractorName            string  `json:"contractor_name"`
	Beneficiary               string  `json:"beneficiary,omitempty"`
	ContractNumber            string  `json:"contract_number"`
	ContractName              string  `json:"contract_name"`
	ContractType              string  `json:"contract_type"`
	ContractTypeEng           string  `json:"contract_type_eng"`
	Status                    string  `json:"status"`
	Author                    string  `json:"author"`
	Amount                    float32 `json:"amount"`
	CreatedAt                 string  `json:"created_at,omitempty"`
	UpdatedAt                 string  `json:"updated_at,omitempty"`
	IsExtendContract          bool    `json:"is_extend_contract"`
	ExtendDate                string  `json:"extend_date"`
	AdditionalAgreementNumber int     `json:"additional_agreement_number"`
	EndDate                   string  `json:"end_date"`
	StartDate                 string  `json:"start_date"`
	Bin                       string  `json:"bin"`
	View                      string  `json:"view"`
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
	ContractNumber string  `json:"contract_number"`
	ContractName   string  `json:"contract_name"`
	ContractAmount float32 `json:"contract_amount"`

	// update fields
	CurrencyName  string `json:"currency_name,omitempty"`
	CurrencyCode  string `json:"currency_code,omitempty"`
	PricetypeName string `json:"pricetype_name,omitempty"`
	PricetypeCode string `json:"pricetype_code,omitempty"`

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
	CurrencyName              string  `json:"currency_name,omitempty"`
	CurrencyCode              string  `json:"currency_code,omitempty"`
	PricetypeName             string  `json:"pricetype_name,omitempty"`
	PricetypeCode             string  `json:"pricetype_code,omitempty"`
	//PriceType                 string  `json:"price_type,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type Product struct {
	ProductNumber    string    `json:"product_number,omitempty"`
	ProductName      string    `json:"product_name"`
	Price            float64   `json:"price,omitempty"`
	Currency         string    `json:"currency,omitempty"`
	Substance        string    `json:"substance"`
	StorageCondition string    `json:"storage_condition"`
	Producer         string    `json:"producer"`
	Sku              string    `json:"sku"`
	SkuName          string    `json:"sku_name"`
	Plan             float32   `json:"plan"`
	DiscountPercent  float32   `json:"discount_percent"`
	PriceType        PriceType `json:"price_type"`
}

type DoubtedDiscountResponse struct {
	RBRequest       RBRequest   `json:"rb_request"`
	DoubtedDiscount []RBRequest `json:"doubted_discount"`
}

type DoubtedDiscount struct {
	ContractNumber string                   `json:"contract_number"`
	Discounts      []DoubtedDiscountDetails `json:"discounts"`
}

type DoubtedDiscountDetails struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	IsCompleted bool   `json:"is_completed"`
}

type Discount struct {
	Name            string           `json:"name,omitempty"`
	Code            string           `json:"code"`
	DiscountAmount  int              `json:"discount_amount,omitempty"`
	IsSelected      bool             `json:"is_selected"`
	PeriodFrom      string           `json:"period_from"`
	IsSale          bool             `json:"is_sale"`
	PeriodTo        string           `json:"period_to"`
	DiscountPercent float32          `json:"discount_percent"`
	GrowthPercent   float32          `json:"growth_percent"`
	Periods         []DiscountPeriod `json:"periods,omitempty"`
	DiscountBrands  []DiscountBrands `json:"discount_brands"`
	Products        []Product        `json:"products"`
}

type ResponseDiscount struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	IsSale  bool   `json:"is_sale"`
	Periods []struct {
		PeriodTo        string `json:"period_to"`
		PeriodFrom      string `json:"period_from"`
		TotalAmount     int    `json:"total_amount"`
		RewardAmount    int    `json:"reward_amount"`
		GrowthPercent   int    `json:"growth_percent"`
		DiscountPercent int    `json:"discount_percent"`
	} `json:"periods"`
	PeriodTo        string      `json:"period_to"`
	IsSelected      bool        `json:"is_selected"`
	PeriodFrom      string      `json:"period_from"`
	GrowthPercent   int         `json:"growth_percent"`
	DiscountBrands  interface{} `json:"discount_brands"`
	DiscountPercent int         `json:"discount_percent"`
}

type DiscountBrands struct {
	PeriodFrom string     `json:"period_from"`
	PeriodTo   string     `json:"period_to"`
	Brands     []BrandDTO `json:"brands"`
}

type BrandDTO struct {
	DiscountPercent float32 `json:"discount_percent"`
	PurchaseAmount  float32 `json:"purchase_amount"`
	SalesAmount     float32 `json:"sales_amount"`
	BrandName       string  `json:"brand_name"`
	BrandCode       string  `json:"brand_code"`
}

type DiscountPeriod struct {
	PeriodFrom      string  `json:"period_from"`
	PeriodTo        string  `json:"period_to"`
	TotalAmount     float32 `json:"total_amount"`
	RewardAmount    int     `json:"reward_amount"`
	DiscountPercent float32 `json:"discount_percent"`
	Type            string  `json:"type,omitempty"`
	Name            string  `json:"name,omitempty"`
	PurchaseAmount  float32 `json:"purchase_amount,omitempty"`
	GrowthPercent   float32 `json:"growth_percent,omitempty"`
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
	ID             int    `json:"id"`
	ContractNumber string `json:"contract_number"`
	ContractType   string `json:"contract_type"`
	Author         string `json:"author"`
	Status         string `json:"status"`
	StartDate      string `json:"start_date" gorm:"created_at"`
	EndDate        string `json:"end_date"`
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

type DataPurchase struct {
	Bin            string `json:"bin"`
	ContractNumber string `json:"contract_number"`
	DiscountAmount string `json:"discount_amount"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
}

type ResponseContractFrom1C struct {
	Bin         string `json:"bin"`
	ContractArr []struct {
		ContractName            string `json:"contract_name"`
		ContractCode            string `json:"contract_code"`
		ContractCurrencyName    string `json:"contract_currency_name"`
		ContractCurrencyCode    string `json:"contract_currency_code"`
		ContractPricetypeName   string `json:"contract_pricetype_name"`
		ContractPricetypeCode   string `json:"contract_pricetype_code"`
		ContractTotal           string `json:"contract_total"`
		ContractDate            string `json:"contract_date"`
		ContractNumber          string `json:"contract_number"`
		ContractAddress         string `json:"contract_address"`
		ContractDuration        string `json:"contract_duration"`
		ContractAllowableAmount string `json:"contract_allowable_amount"`
		ContractTimeLimit       string `json:"contract_time_limit"`
		ContractType            string `json:"contract_type"`
	} `json:"contract_arr"`
}
