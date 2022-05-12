package models

type Regions struct {
	RegionName string `json:"region_name"`
	RegionCode string `json:"region_code"`
}

type PayloadProduct struct {
		TypeValue string `json:"typeValue"`
		TypeParameters string `json:"typeParameters"`

}