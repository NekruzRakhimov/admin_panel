package models

import "github.com/lib/pq"

type StoreRegion struct {
	StoreName  string `json:"store_name"`
	StoreCode  string `json:"store_code"`
	RegionName string `json:"region_name"`
	RegionCode string `json:"region_code"`
	OrgCode    string `json:"org_code"`
	OrgName    string `json:"org_name"`
	DrugStore  string `json:"drug_store"`
}

type Organize struct {
	OrgCode string `json:"org_code"`
	OrgName string `json:"org_name"`
}

type Pharmacy struct {
	StoreName string `json:"store_name"`
	StoreCode string `json:"store_code"`
	DrugStore string `json:"drug_store"`
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
	KmType            string         `json:"km_type"`
}

type AutoOrder struct {
	ID                int    `json:"id"`
	GraphicName       string `json:"graphic_name"`
	Formula           string `json:"formula"`
	ByMatrix          bool   `json:"by_matrix"`
	Schedule          string `json:"schedule"`
	FormedOrdersCount int    `json:"formed_orders_count"`
	Organization      string `json:"organization"`
	Status            string `json:"status"`
	Store             string `json:"store"`
	CreatedAt         string `json:"created_at"`
	FormedAt          string `json:"formed_at"`
	SentAt            string `json:"sent_at"`
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

type FormedGraphic struct {
	ID                      int                    `json:"id"`
	FormulaID               int                    `json:"-"`
	GraphicID               int                    `json:"-"`
	GraphicName             string                 `json:"graphic_name"`
	Supplier                string                 `json:"supplier"` // поставщика
	Store                   string                 `json:"store"`
	ByMatrix                bool                   `json:"by_matrix"`
	Schedule                string                 `json:"schedule"`
	ProductAvailabilityDays int                    `json:"product_availability_days"`
	DisterDays              int                    `json:"dister_days"`
	StoreDays               int                    `json:"store_days"`
	Status                  string                 `json:"status"`
	Products                []FormedGraphicProduct `json:"-"`
	IsLetter                bool                   `json:"is_letter"`
	CreatedAt               string                 `json:"created_at"`
	ExtDocumentNumber       string                 `json:"ext_document_number"`
}

type FormedGraphicProduct struct {
	ID                      int     `json:"id"`
	GraphicID               int     `json:"graphic_id"`
	ProductCode             string  `json:"product_code"`
	ProductName             string  `json:"product_name"`
	OrderQnt                float64 `json:"order_qnt"`
	Days                    int     `json:"days"`
	Remainder               float64 `json:"remainder"`
	ProductAvailabilityDays int     `json:"product_availability_days"`
	SalesCount              float64 `json:"sales_count"`
	SalesDayCount           float64 `json:"sales_day_count"`
	Koef                    int     `json:"koef"`
	TotalStoreCount         float64 `json:"total_store_count"`
	Min                     float64 `json:"min"`
	Max                     float64 `json:"max"`
	StoreCode               string  `json:"store_code"`
}

type Order1С struct {
	OrderId      string `json:"order_id"`
	SupplierCode string `json:"supplier_code"`
	StoreCode    string `json:"store_code"`
	Products     []struct {
		ProductCode string `json:"product_code"`
		SalesCount  string `json:"sales_count"`
		Price       string `json:"price"`
	} `json:"products"`
}

type Data struct {
	OrderId      string         `json:"order_id"`
	SupplierCode string         `json:"supplier_code"`
	StoreCode    string         `json:"store_code"`
	Products     []DataProducts `json:"products"`
}

type DataProducts struct {
	ProductCode string `json:"product_code"`
	SalesCount  string `json:"sales_count"`
	Price       string `json:"price"`
}
