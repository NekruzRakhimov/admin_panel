package models

type Segment struct {
	ID          int    `json:"id"`
	SegmentCode string `json:"segment_code"`
	NameSegment string `json:"name_segment"`
	Beneficiary string `json:"beneficiary"`
	Bin         string `json:"bin"`
	ClientCode  string `json:"client_code"`
	Email       string `json:"email"`
	//Counterparty      string   `json:"counterparty"`
	ForMarket  bool                `json:"for_market"`
	Products   []ListsNomenclature `json:"product"`
	ProductStr string              `json:"-" gorm:"column:product"`
	Region     []Region            `json:"region" gorm:"-"`
	RegionStr  string              `json:"-" gorm:"column:region"`
}

type ListsNomenclature struct {
	ProductCode string `json:"product_code"`
	ProductName string `json:"product_name"`
}

type Region struct {
	Counterparty string `json:"counterparty"`
	Region       string `json:"region"`
	RegionCode   string `json:"region_code"`
	Email        string `json:"email"`
}
