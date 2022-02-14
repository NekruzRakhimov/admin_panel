package model

type RBRequest struct {
	BIN            string `json:"bin"`
	ContractorName string `json:"contractor_name"`
	PeriodFrom     string `json:"period_from"`
	PeriodTo       string `json:"period_to"`
}

type RbDTO struct {
	ID             int    `json:"id"`
	ContractNumber string `json:"contract_number"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	DiscountAmount int    `json:"discount_amount"`
}
