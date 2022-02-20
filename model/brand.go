package model

type Brand struct {
	BrandArr []struct {
		BrandName string `json:"brand_name"`
		BrandCode string `json:"brand_code"`
	} `json:"brand_arr"`
}

type Sales struct {
	SalesArr []struct {
		ProductName string `json:"product_name"`
		ProductCode string `json:"product_code"`
		Total       int    `json:"total"`
		QntTotal    int    `json:"qnt_total"`
		Date        string `json:"date"`
		BrandCode   string `json:"brand_code"`
	} `json:"sales_arr"`
}

type DateSales struct {
	Datestart string `json:"datestart"`
	Dateend   string `json:"dateend"`
	ClientBin string `json:"client_bin"`
}

type AddBrand struct {
	BrandName string `json:"brand_name"`
	BrandCode string `json:"brand_code,omitempty"`
}

type DiscountBrand struct {
	Id              int    `json:"id"`
	Brand           string `json:"brand"`
	DiscountPercent int    `json:"discount_percent"`
	ContractId      int    `json:"contract_id"`
}

type BrandInfo struct {
	Id              int     `json:"id"`
	ContractInfo    int     `json:"contract_info"`
	Brand           string  `json:"brand"`
	DiscountPercent int     `json:"discount_percent"`
	ContractId      int     `json:"contract_id"`
	Total           float64 `json:"total"`
	DiscountSum     float64 `json:"discount_sum"`
}

type ReqBrand struct {
	ClientBin      string   `json:"client_bin"`
	Beneficiary    string   `json:"beneficiary"`
	DateStart      string   `json:"datestart"`
	DateEnd        string   `json:"dateend"`
	Type           string   `json:"type"`
	TypeValue      string   `json:"typeValue"`
	TypeParameters []string `json:"typeParameters"`
}

type T struct {
	Datestart      string   `json:"datestart"`
	Dateend        string   `json:"dateend"`
	ClientBin      string   `json:"client_bin"`
	Type           string   `json:"type"`
	TypeValue      string   `json:"typeValue"`
	TypeParameters []string `json:"typeParameters"`
}
