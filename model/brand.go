package model

type Brand struct {
	BrandArr []struct {
		BrandName string `json:"brand_name"`
		BrandCode string `json:"brand_code"`
	} `json:"brand_arr"`
}

type Sales struct {
	SalesArr []struct {
		ProductName string  `json:"product_name"`
		ProductCode string  `json:"product_code"`
		Total       int     `json:"total"`
		QntTotal    float64 `json:"qnt_total"`
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