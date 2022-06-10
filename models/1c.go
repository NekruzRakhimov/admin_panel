package models

type Regions struct {
	RegionName string `json:"region_name"`
	RegionCode string `json:"region_code"`
}

type PayloadProduct struct {
	TypeValue      string `json:"typeValue"`
	TypeParameters string `json:"typeParameters"`
}

//type Segment struct {
//	SegmentCode       string   `json:"segment_code"`
//	NameSegment       string   `json:"name_segment"`
//	ListsNomenclature []ListsNomenclature `json:"lists_nomenclature"`
//	Counterparty      string   `json:"counterparty"`
//	ForMarket         bool     `json:"for_market"`
//}

//type ListsNomenclature struct {
//	ProductCode string `json:"product_code"`
//	ProductName string `json:"product_name"`
//}
