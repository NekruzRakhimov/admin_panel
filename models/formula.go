package models

type Formula struct {
	ID                                int       `json:"id"`
	FormulaNumber                     string    `json:"formula_number"`
	Author                            string    `json:"author"`
	FormulaName                       string    `json:"formula_name"`
	AutoOrderType                     string    `json:"auto_order_type"`
	KmType                            string    `json:"km_type"`
	AnalysisPeriod                    string    `json:"analysis_period"`
	ProductAvailabilityDaysLessThan15 bool      `json:"product_availability_days_less_than_15"`
	TransitGoodDisterDays             int       `json:"transit_good_dister_days"`
	TransitGoodStoreDays              int       `json:"transit_good_store_days"`
	AccountPermissionToOrder          bool      `json:"account_permission_to_order"`
	AccountAqnietStoreRemainder       bool      `json:"account_aqniet_store_remainder"`
	ExpirationDate                    string    `json:"expiration_date"`
	PackingNorm                       bool      `json:"packing_norm"`
	IsMtz                             bool      `json:"is_mtz"`
	SalesLessThanMtz                  bool      `json:"sales_less_than_mtz"`
	AccountIlliquid                   bool      `json:"account_illiquid"`
	AccountWriteOff                   bool      `json:"account_write_off"`
	NeedAutomatedCalculation          bool      `json:"need_automated_calculation"`
	FormulaStringName                 string    `json:"formula_string_name"`
	FormulaStringCode                 string    `json:"formula_string_code"`
	NoSalesNoRemainderToOrder         float64   `json:"no_sales_no_remainder_to_order"`
	Organization                      string    `json:"organization"`
	StoreHouse                        string    `json:"store_house"`
	Schedule                          Schedule  `json:"schedule" gorm:"-"`
	ScheduleStr                       string    `json:"-" gorm:"column:schedule"`
	Product                           []Product `json:"product"`
	MinOrderSum                       float32   `json:"min_order_sum"`
	AutoSentTo1c                      bool      `json:"auto_sent_to_1c" gorm:"column:auto_sent_to_1c"`
}

type Schedule struct {
	Time     string `json:"time"`
	EveryDay struct {
		IsSelected bool `json:"is_selected"`
		Every      int  `json:"every"`
		Options    []struct {
			Name       string `json:"name"`
			IsSelected bool   `json:"is_selected"`
		}
	} `json:"every_day"`
	EveryWeek struct {
		IsSelected bool `json:"is_selected"`
		Every      int  `json:"every"`
		Options    []struct {
			Name       string `json:"name"`
			IsSelected bool   `json:"is_selected"`
		}
	} `json:"every_week"`
}

type FormulaParameters struct {
	NameRus string `json:"name_rus"`
	Code    string `json:"code"`
}
