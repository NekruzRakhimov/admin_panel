package models

type Counterparty struct {
	ID1C                 string                 `json:"id_1C"`
	Name                 string                 `json:"name"`
	Organization         string                 `json:"organization"`
	BIN                  string                 `json:"bin"`
	ContractCounterparty []ContractCounterparty `json:"contracts"`
}

type ContractCounterparty struct {
	Discount  int    `json:"discount"`
	Name      string `json:"name"`
	ID1C      string `json:"id_1C"`
	PriceType string `json:"price_type"`
}

//
type RespContract struct {
	ContractCode string `json:"contract_code"`
	Status       string `json:"status"`
}
