package models

type DefectsSearchParameters struct {
	Date     *DateFilter
	Pharmacy *string
}

type DefectsFile struct {
	FileName string
	File     string
}
