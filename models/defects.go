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
	ProductCode      string `json:"product_code"`
	ProductName      string `json:"product_name"`
	RegionCode       string `json:"region_code"`
	RegionName       string `json:"region_name"`
	StoreCode        string `json:"store_code"`
	StoreName        string `json:"store_name"`
	DefectQnt        string `json:"defect_qnt"`
	StoreSaldoQnt    string `json:"store_saldo_qnt"`
	StoreSaldoTotal  string `json:"store_saldo_total"`
	DefectTotal      string `json:"defect_total"`
	DefectTotalQnt   string `json:"defect_total_qnt"`
	MatrixTotalSales string `json:"matrix_total_sales"`
	DifPercent       string `json:"dif_percent"`
}

type DefectsFiltered struct {
	StoreCode  string   `json:"store_code"`
	StoreName  string   `json:"store_name"`
	SubDefects []Defect `json:"sub_defects"`
}

type DefectsRequest struct {
	Startdate string `json:"startdate"`
	Enddate   string `json:"enddate"`
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
