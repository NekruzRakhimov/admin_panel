package models

import "github.com/lib/pq"

type StoreRegion struct {
	StoreName  string `json:"store_name"`
	StoreCode  string `json:"store_code"`
	RegionName string `json:"region_name"`
	RegionCode string `json:"region_code"`
}

type Matrix struct {
	StoreName    string `json:"store_name"`
	StoreCode    string `json:"store_code"`
	RegionName   string `json:"region_name"`
	RegionCode   string `json:"region_code"`
	ProductName  string `json:"product_name"`
	ProductCode  string `json:"product_code"`
	SupplierName string `json:"supplier_name"`
	SupplierCode string `json:"supplier_code"`
	Format       string `json:"format"`
	Min          string `json:"min"`
	Max          string `json:"max"`
	Import       string `json:"import"`
	Defect       string `json:"defect"`
}

type Graphic struct {
	ID                int            `json:"id"`
	Number            string         `json:"number"`
	Author            string         `json:"author"`
	SupplierName      string         `json:"supplier_name"`
	SupplierCode      string         `json:"supplier_code"`
	RegionName        string         `json:"region_name"`
	RegionCode        string         `json:"region_code"`
	StoreName         string         `json:"store_name"`
	StoreCode         string         `json:"store_code"`
	OnceAMonth        bool           `json:"once_a_month"`
	TwiceAMonth       bool           `json:"twice_a_month"`
	IsOn              bool           `json:"is_on"`
	NomenclatureGroup string         `json:"nomenclature_group"`
	AutoOrderDate     string         `json:"auto_order_date"`
	ApplicationDay    pq.StringArray `json:"application_day"`
	ExecutionPeriod   string         `json:"execution_period"`
}
