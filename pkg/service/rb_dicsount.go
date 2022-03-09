package service

import (
	"admin_panel/db"
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func RbDiscountForSalesGrowth(rb model.RBRequest) (float32, float32, float32) {
	pastTimeFrom, err := ConvertTime(rb.PeriodFrom)
	if err != nil {
	}
	pastTimeTo, err := ConvertTime(rb.PeriodTo)

	pastPeriod := model.RBRequest{
		BIN:            rb.BIN,
		Type:           rb.Type,
		ContractorName: rb.ContractorName,
		PeriodFrom:     pastTimeFrom,
		PeriodTo:       pastTimeTo,
	}
	fmt.Println("pastPeriod", pastPeriod)
	fmt.Println("rbM", rb)

	// берем growth and percent ->
	//repository.GetRbSalesGrowth(rb.BIN)

	presentPeriod, err := GetSales1C(rb, "sales")
	oldPeriod, err := GetSales1C(pastPeriod, "sales")
	var preCoutnt float32
	var pastCount float32

	fmt.Println("presentPeriod", presentPeriod)

	for _, present := range presentPeriod.SalesArr {
		preCoutnt += present.Total
	}
	for _, past := range oldPeriod.SalesArr {
		pastCount += past.Total

	}

	total := (pastCount * 100 / preCoutnt) - 100
	// var total float32
	// total =   1_500_000* 100 / 1_000_000  - 100

	// ты должен взять сумму прироста - то есть ты будешь с ним сравнивать
	// и также ты из бд должен взять сумму скидки и дать ему скидку
	if total > 10 {

	}

	// call 1C
	// call again 1C
	// считаем сумму с обеиъ
	// после чего находим

	return pastCount, preCoutnt, total
}

func ConvertTime(date string) (string, error) {
	timeSplit := strings.Split(date, ".")
	if len(timeSplit) != 3 {
		return "", errors.New("len of time must be 3")
	}
	fmt.Println(timeSplit)
	convertYear, err := strconv.Atoi(timeSplit[2])
	if err != nil {
		log.Println(err)
		return "", err
	}
	convertYear -= 1
	updateTime := fmt.Sprintf("%s.%s.%d", timeSplit[0], timeSplit[1], convertYear)
	//fmt.Println(sprintf)

	return updateTime, nil
}

func DiscountRBPeriodTime(req model.RBRequest) ([]model.RbDTO, error) {
	var rbDTOsl []model.RbDTO

	// parsing string by TIME
	layoutISO := "02.1.2006"
	var count int
	// parsing string to Time
	reqPeriodFrom, _ := time.Parse(layoutISO, req.PeriodFrom)
	reqPeriodTo, _ := time.Parse(layoutISO, req.PeriodTo)

	// get all contracts_code by BIN
	externalCodes := GetExternalCode(req.BIN)
	var contractsCode []string
	for _, value := range externalCodes {
		contractsCode = append(contractsCode, value.ExtContractCode)
	}
	contractsWithJson, err := repository.GetAllContractDetailByBIN(req.BIN, req.PeriodFrom, req.PeriodTo)
	if err != nil {
		return nil, err
	}
	fmt.Println("contractsWithJson", contractsWithJson)
	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	for _, contract := range contracts {
		fmt.Println("contract MESSAGE", contract.Discounts)
		for _, discount := range contract.Discounts {
			if discount.Code == "RB_DISCOUNT_FOR_PURCHASE_PERIOD" { // здесь сравниваешь тип скидки и берешь тот тип который тебе нужен
				for _, period := range discount.Periods {
					PeriodFrom, _ := time.Parse(layoutISO, period.PeriodFrom)
					PeriodTo, _ := time.Parse(layoutISO, period.PeriodTo)
					if PeriodFrom.After(reqPeriodFrom) || PeriodTo.Before(reqPeriodTo) {
						reqBrand := model.ReqBrand{
							ClientBin:      req.BIN,
							DateStart:      req.PeriodFrom,
							DateEnd:        req.PeriodFrom,
							TypeValue:      "",
							TypeParameters: nil,
							Contracts:      contractsCode, // необходимо получить коды контрактов
						}
						purchase, _ := GetPurchase(reqBrand)
						for _, amount := range purchase.PurchaseArr {
							count += amount.Total
						}
						if period.PurchaseAmount < float32(count) {
							total := float32(count) * period.DiscountPercent / 100

							RbDTO := model.RbDTO{
								ContractNumber:       contract.ContractParameters.ContractNumber,
								StartDate:            period.PeriodFrom,
								EndDate:              period.PeriodTo,
								TypePeriod:           period.Type,
								DiscountPercent:      period.DiscountPercent,
								DiscountAmount:       total,
								TotalWithoutDicsount: float32(count),
							}
							rbDTOsl = append(rbDTOsl, RbDTO)

						}

					}

				}

			}
		}
	}

	return rbDTOsl, nil
}

func GetExternalCode(bin string) []model.ContractCode {
	var ExtContractCode []model.ContractCode
	db.GetDBConn().Raw("SELECT ext_contract_code FROM contracts WHERE requisites ->> 'bin' =  $1", bin).Scan(&ExtContractCode)

	return ExtContractCode
}
