package service

import (
	"admin_panel/db"
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func GetRB13thType(rb models.RBRequest) ([]models.RbDTO, error) {
	log.Println("ФУНКЦИЯ ПО ПРИРОСТУ ВЫЗВАЛАСЬ===================================================================================================")
	fmt.Println("ФУНКЦИЯ ПО ПРИРОСТУ ВЫЗВАЛАСЬ===================================================================================================")
	var rbDTOsl []models.RbDTO

	// чтобы преобразоват дату в ввиде День.Месяц.Год
	layoutISO := "02.1.2006"

	// parsing string to Time
	reqPeriodFrom, _ := time.Parse(layoutISO, rb.PeriodFrom)
	reqPeriodTo, _ := time.Parse(layoutISO, rb.PeriodTo)
	fmt.Println(reqPeriodFrom)
	fmt.Println(reqPeriodTo)
	contractsWithJson, err := repository.GetAllContractDetailByBIN(rb.BIN, rb.PeriodFrom, rb.PeriodTo)
	if err != nil {
		return nil, err
	}
	fmt.Println("contractsWithJson=============================================", contractsWithJson)
	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	for _, contract := range contracts {
		fmt.Println("contract MESSAGE", contract.Discounts)

		// от сюда берем скидки и периоды
		for _, discount := range contract.Discounts {
			// после всех проверок логика начнется
			if discount.Code == "RB_DISCOUNT_FOR_SALES_GROWTH" {
				for _, period := range discount.Periods {
					periodFrom, _ := time.Parse(layoutISO, period.PeriodFrom)
					periodTo, _ := time.Parse(layoutISO, period.PeriodTo)
					fmt.Println(periodFrom)
					fmt.Println(periodTo)

					//if periodFrom.After(reqPeriodFrom) || periodTo.Before(reqPeriodTo) {
					pastTimeFrom, err := ConvertTime(period.PeriodFrom)
					if err != nil {
						return nil, err
					}
					pastTimeTo, err := ConvertTime(period.PeriodTo)
					if err != nil {
						return nil, err
					}

					// это чтобы брали на 1 год меньше
					pastPeriod := models.ReqBrand{
						ClientBin:      rb.BIN,
						DateStart:      pastTimeFrom,
						DateEnd:        pastTimeTo,
						Type:           "",
						TypeValue:      "",
						TypeParameters: nil,
						Contracts:      nil,
					}

					// Это необходимо, чтобы получить продажи за тек период
					present := models.ReqBrand{
						ClientBin:      rb.BIN,
						Beneficiary:    "",
						DateStart:      rb.PeriodFrom,
						DateEnd:        rb.PeriodTo,
						Type:           "",
						TypeValue:      "",
						TypeParameters: nil,
						Contracts:      nil,
					}
					// берем продажи за тек год и за 1 год меньше
					presentPeriod, err := GetSales1C(present, "sales_brand_only")
					fmt.Println("PRESENT", presentPeriod)

					if err != nil {
						return nil, err
					}
					oldPeriod, err := GetSales1C(pastPeriod, "sales_brand_only")
					fmt.Println("PAST ==================", oldPeriod)
					if err != nil {
						return nil, err
					}
					var preCount float32
					var pastCount float32

					// считаем за тек период
					for _, present := range presentPeriod.SalesArr {
						preCount += present.Total
					}
					// считаем за прошлый год
					for _, past := range oldPeriod.SalesArr {
						pastCount += past.Total

					}
					fmt.Println("Сумма за настоящее", preCount)
					fmt.Println("Сумма за прошлый год", pastCount)

					// находим прирост в процентах
					//growthPercent := (pastCount * 100 / preCount) - 100

					// находим разницу за нынешний год
					diff := preCount - pastCount
					growthPercent := (100 * diff) / pastCount
					fmt.Println("growthPercent", growthPercent)
					// проверяем разницу с тек по прошлогодний год, если процент прироста выше, логика выполнится

					fmt.Println("growth_percent", period.GrowthPercent)
					fmt.Println("discount percent", period.DiscountPercent)
					fmt.Println()
					if growthPercent > period.GrowthPercent {
						discountAmount := preCount * period.DiscountPercent / 100

						fmt.Println("discountAmount", discountAmount)

						rbDTO := models.RbDTO{
							ContractNumber:       contract.ContractParameters.ContractNumber,
							StartDate:            period.PeriodFrom,
							EndDate:              period.PeriodTo,
							TypePeriod:           "",
							BrandName:            "",
							ProductCode:          "",
							DiscountPercent:      period.DiscountPercent,
							DiscountAmount:       discountAmount,
							TotalWithoutDiscount: preCount,
							DiscountType:         RB13Name,
						}
						rbDTOsl = append(rbDTOsl, rbDTO)

					} else {
						rbDTO := models.RbDTO{
							ContractNumber:       contract.ContractParameters.ContractNumber,
							StartDate:            period.PeriodFrom,
							EndDate:              period.PeriodTo,
							TypePeriod:           "",
							BrandName:            "",
							ProductCode:          "",
							DiscountPercent:      period.DiscountPercent,
							DiscountAmount:       0,
							TotalWithoutDiscount: preCount,
							DiscountType:         RB13Name,
						}
						rbDTOsl = append(rbDTOsl, rbDTO)
					}

				}

			}

		}
	}
	//}
	return rbDTOsl, nil
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

func GetRB12thType(req models.RBRequest) ([]models.RbDTO, error) {
	var rbDTOsl []models.RbDTO

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
						reqBrand := models.ReqBrand{
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

							RbDTO := models.RbDTO{
								ContractNumber:       contract.ContractParameters.ContractNumber,
								StartDate:            period.PeriodFrom,
								EndDate:              period.PeriodTo,
								TypePeriod:           period.Type,
								DiscountPercent:      period.DiscountPercent,
								DiscountAmount:       total,
								TotalWithoutDiscount: float32(count),
								DiscountType:         RB12Name,
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

func GetExternalCode(bin string) []models.ContractCode {
	var ExtContractCode []models.ContractCode
	db.GetDBConn().Raw("SELECT ext_contract_code FROM contracts WHERE requisites ->> 'bin' =  $1", bin).Scan(&ExtContractCode)

	return ExtContractCode
}
