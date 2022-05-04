package models

type RBRequest struct {
	BIN              string            `json:"bin"`
	Type             string            `json:"type"`
	ContractorName   string            `json:"contractor_name"`
	PeriodFrom       string            `json:"period_from"`
	PeriodTo         string            `json:"period_to"`
	ClientCode     string `json:"client_code"`
	DoubtedDiscounts []DoubtedDiscount `json:"doubted_discounts"`
}

type RbDTO struct {
	ID                   int     `json:"id"`
	ContractNumber       string  `json:"contract_number"`
	StartDate            string  `json:"start_date"`
	EndDate              string  `json:"end_date"`
	TypePeriod           string  `json:"type_period"`
	BrandName            string  `json:"brand_name,omitempty"`
	ProductCode          string  `json:"product_code,omitempty"`
	DiscountPercent      float32 `json:"discount_percent"`
	DiscountAmount       float32 `json:"discount_amount"`
	TotalWithoutDiscount float32 `json:"TotalWithoutDiscount"`
	LeasePlan            float32 `json:"lease_plan"`
	RewardAmount         float32 `json:"reward_amount"`
	DiscountType         string  `json:"discount_type"`
	Status               string  `json:"status"`
	RegionName           string  `json:"region_name"`
	RegionCode           string  `json:"region_code"`
}

type Block struct {
	TotalBlockNum string         `json:"total_block_num"`
	StartDate     string         `json:"start_date"`
	EndDate       string         `json:"end_date"`
	ClientBin     string         `json:"client_bin"`
	BlockNum      string         `json:"block_num"`
	ReqBody       []BlockProduct `json:"req_body"`
}
type BlockProduct struct {
	ProductName string `json:"product_name"`
	ProductCode string `json:"product_code"`
	Total       int    `json:"total"`
	QntTotal    int    `json:"qnt_total"`
	Date        string `json:"date"`
	StoreCode   string `json:"store_code"`
	StoreName   string `json:"store_name"`
	BrandCode   string `json:"brand_code"`
	BrandName   string `json:"brand_name"`
	Control     string `json:"control"`
}
