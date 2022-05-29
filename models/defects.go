package models

type DefectsSearchParameters struct {
	Date     *DateFilter
	Pharmacy *string
}

type DefectsFile struct {
	FileName string
	File     string
}

type Defect struct {
	MainDrugStore    string `json:"main_drug_store"`
	PF               string `json:"PF"`
	ProductCode      string `json:"product_code"`
	ProductName      string `json:"product_name"`
	StoreCode        string `json:"store_code"`
	StoreName        string `json:"store_name"`
	DefectQnt        string `json:"defect_qnt"`
	StoreSaldoQnt    string `json:"store_saldo_qnt"`
	MatrixSales      string `json:"matrix_sales"`
	MatrixProductQnt string `json:"matrix_product_qnt"`
	DefectPrice      string `json:"defect_price"`
}

type DefectsFiltered struct {
	StoreCode  string   `json:"store_code"`
	StoreName  string   `json:"store_name"`
	SubDefects []Defect `json:"sub_defects"`
}

type DefectsRequest struct {
	Startdate string `json:"startdate"`
	Enddate   string `json:"enddate"`
	QueryType string `json:"queryType"` //warehouse_defect - по складам / drugstore_defect - по аптекам
	DaysCount int    `json:"days_count"`
	IsPF      bool   `json:"isPF"`
}

type SalesCountRequest struct {
	Startdate string `json:"startdate"`
	Enddate   string `json:"enddate"`
	StoreCode string `json:"store_code"`
}

type SalesCount struct {
	ProductCode        string `json:"product_code"`
	ProductName        string `json:"product_name"`
	TotalSalesDayCount string `json:"total_sales_day_count"`
	SalesDayCount      string `json:"sales_day_count"`
	SalesCount         string `json:"sales_count"`
	TotalStoreCount    string `json:"total_store_count"`
}
