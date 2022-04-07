package models

type StoredReport struct {
	ID                         int     `json:"id"`
	Bin                        string  `json:"bin"`
	ContractID                 int     `json:"contract_id"`
	ContractNumber             string  `json:"contract_number"`
	ContractDate               string  `json:"contract_date"`
	Beneficiary                string  `json:"beneficiary"`
	StartDate                  string  `json:"start_date"`
	EndDate                    string  `json:"end_date"`
	ContractAmount             float32 `json:"contract_amount"`
	DiscountAmount             float32 `json:"discount_amount"`
	ContractAmountWithDiscount float32 `json:"contract_amount_with_discount"`
}
