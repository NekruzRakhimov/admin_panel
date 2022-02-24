package model

type RBRequest struct {
	BIN            string `json:"bin"`
	ContractorName string `json:"contractor_name"`
	PeriodFrom     string `json:"period_from"`
	PeriodTo       string `json:"period_to"`
}

type RbDTO struct {
	ID              int     `json:"id"`
	ContractNumber  string  `json:"contract_number"`
	StartDate       string  `json:"start_date"`
	EndDate         string  `json:"end_date"`
	BrandName       string  `json:"brand_name,omitempty"`
	ProductCode     string  `json:"product_code,omitempty"`
	DiscountPercent float32 `json:"discount_percent"`
	DiscountAmount  float32 `json:"discount_amount"`
	LeasePlan       float32 `json:"lease_plan"`
}
