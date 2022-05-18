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
	ContractAmount             float64 `json:"contract_amount"`
	DiscountAmount             float64 `json:"discount_amount"`
	Content                    []byte  `json:"content"`
	ContractAmountWithDiscount float64 `json:"contract_amount_with_discount"`
	CreatedAt                  string  `json:"created_at"`
	Author                     string  `json:"author"`
}
