package model


type Notification struct {
	Bin string `json:"bin"`
	ContractDate string `json:"contract_date"`
	ContractNumber string `json:"contract_number"`
	Type string `json:"type"`
	Email string `json:"email"`
	
}

