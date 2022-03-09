package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
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

func DiscountRBPeriodTime(request model.RBRequest) ([]model.RbDTO, error) {
	//TODO: из базы мы должны взять массив скидок (типы периодов)
	// тут будем вызов массива по периодам (цикл)
	// внутри цикла возьмем данные по закупам из 1С
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == "" { // здесь сравниваешь тип скидки и берешь тот тип который тебе нужен

			}
		}
	}

	//var counts  []float32
	//var count float32
	////periods, _ := repository.GetDicsountPeriod(rb.BIN)
	//
	//
	//contractsWithJson, err := repository.GetAllContractDetailByBIN(rb.BIN, rb.PeriodFrom, rb.PeriodTo)
	//if err != nil {
	//	return nil, err
	//}
	//
	//contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	//if err != nil {
	//	return nil, err
	//}
	//var contractCode []string
	//
	//
	//for _, contract := range contracts{
	//	for _, discount := range contract.Discounts{
	//		for _, period := range discount.Periods{
	//			contractCode = append(contractCode, period.)
	//		}
	//	}
	//
	//	contractCode = append(contractCode, contract. )
	//
	//	model.ReqBrand{
	//		ClientBin:      "",
	//		Beneficiary:    "",
	//		DateStart:      "",
	//		DateEnd:        "",
	//		Type:           "",
	//		TypeValue:      "",
	//		TypeParameters: nil,
	//		Contracts: nil,
	//	}
	//
	//
	//	r := model.RBRequest{
	//		BIN:              rb.BIN,
	//		Type:             "",
	//		ContractorName:   "",
	//		PeriodFrom:       period.PeriodFrom,
	//		PeriodTo:         period.PeriodTo,
	//		DoubtedDiscounts: nil,
	//	}
	//
	//	purchase, _ := GetSales1C(r, "purchase")
	//
	//	for _, value := range  purchase.SalesArr{
	//		count += value.Total
	//	}
	//	counts = append(counts, count)
	//}

	return nil, nil
}
