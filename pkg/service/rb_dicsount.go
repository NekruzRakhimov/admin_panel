package service

import (
	"admin_panel/model"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func RbDiscountForSalesGrowth(rb model.RBRequest)  (float32, float32) {
	pastTimeFrom, err := ConvertTime(rb.PeriodTo)
	if err != nil {
	}
	pastTimeTo, err := ConvertTime(rb.PeriodFrom)

	pastPeriod := model.RBRequest{
		BIN:            rb.BIN,
		Type:           rb.Type,
		ContractorName: rb.ContractorName,
		PeriodFrom:    	pastTimeFrom,
		PeriodTo:       pastTimeTo,
	}




		// берем growth and percent ->
		//repository.GetRbSalesGrowth(rb.BIN)


	presentPeriod, err := GetSales1C(rb, "sales")
	oldPeriod, err := GetSales1C(pastPeriod, "sales")
	var preCoutnt float32
	var pastCount float32

	for _, present := range presentPeriod.SalesArr{
		preCoutnt += present.Total
	}
	for _, past := range oldPeriod.SalesArr{
		pastCount += past.Total

	}



	// call 1C
	// call again 1C
	// считаем сумму с обеиъ
	// после чего находим

	return pastCount, preCoutnt
}

func ConvertTime(date string)  (string, error) {
	timeSplit := strings.Split(date, ".")
	if len(timeSplit) != 3{
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