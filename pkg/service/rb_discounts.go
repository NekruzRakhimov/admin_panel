package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"admin_panel/utils"
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	sheet    = "Sheet1"
	RB1Name  = "Скидка за объем закупа"
	RB2Name  = "Скидка  на группы товаров"
	RB3Name  = "Скидка за выполнение плана закупа по препаратам"
	RB4Name  = "Скидка за представленность"
	RB5Name  = "Скидка на фиксированную сумма выплаты за МТЗ"
	RB6Name  = "Скидка на МТЗ"
	RB7Name  = "Скидка на РБ за выполнение плана закупа по брендам %"
	RB8Name  = "Скидка на РБ от закупа продукции %"
	RB9Name  = "Скидка на РБ по филиалам"
	RB10Name = "Скидка РБ по логистике"
	RB11Name = "Скидка  РБ за поддержание ассортимента"
	RB12Name = "Скидка РБ объем за закупку в промежутке времени"
	RB13Name = "Скидка  РБ за прирост продаж"
)

const (
	RB1Code  = "TOTAL_AMOUNT_OF_SELLING"
	RB2Code  = "DISCOUNT_BRAND"
	RB3Code  = "DISCOUNT_PLAN_LEASE"
	RB4Code  = "DISCOUNT_FOR_REPRESENTATION"
	RB5Code  = "DISCOUNT_FOR_FIX_SUM_MTZ"
	RB6Code  = "DISCOUNT_FOR_MTZ"
	RB7Code  = "DISCOUNT_FOR_LEASE_PERCENT"
	RB8Code  = "DISCOUNT_FOR_LEASE_GENERAL"
	RB9Code  = "DISCOUNT_FOR_FILIAL"
	RB10Code = "DISCOUNT_FOR_LOGISTIC"
	RB11Code = "DISCOUNT_FOR_ASSORTMENT"
	RB12Code = "RB_DISCOUNT_FOR_PURCHASE_PERIOD"
	RB13Code = "RB_DISCOUNT_FOR_SALES_GROWTH"
)

func GetRB1stType(request models.RBRequest, contracts []models.Contract) ([]models.RbDTO, error) {
	//TODO: посмотри потом
	//testBin := "060840003599"
	//req := models.ReqBrand{
	//	ClientBin:   request.BIN,
	//	Beneficiary: request.ContractorName,
	//	DateStart:   request.PeriodFrom,
	//	DateEnd:     request.PeriodTo,
	//	Type:        "sales",
	//}

	present := models.ReqBrand{
		ClientBin:      request.BIN,
		Beneficiary:    "",
		DateStart:      request.PeriodFrom,
		DateEnd:        request.PeriodTo,
		Type:           "",
		TypeValue:      "",
		TypeParameters: nil,
		Contracts:      nil,
	}

	sales, err := GetSales1C(present, "sales_brand_only")
	//sales, err := GetSales(req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("###%+v\n", contracts)
	totalAmount := GetTotalAmount(sales)

	contractRB := DefiningRBReport(contracts, totalAmount, request)

	return contractRB, nil
}

func GetRB2ndType(rbReq models.RBRequest) []models.RbDTO {
	brandTotal := map[string]float32{}
	var rbDtoSl []models.RbDTO
	fmt.Println("запрос от тебя", rbReq)
	rbBrand := models.ReqBrand{
		ClientBin: rbReq.BIN,
		DateStart: rbReq.PeriodFrom,
		DateEnd:   rbReq.PeriodTo,
	}
	fmt.Println("rbBrand", rbBrand)
	// берем бренды и их Total
	sales, _ := GetSales(rbBrand)

	// тут считаем общую сумму каждого бренда
	for _, sale := range sales.SalesArr {
		brandTotal[sale.BrandName] += sale.Total
	}

	fmt.Println("MAP: ", brandTotal)

	// берем скидки по брендам и название брендов
	dataBrands, contractNumb := repository.GetIDByBIN(rbReq.BIN)
	fmt.Println("dataBrand", dataBrands)
	fmt.Println("sales", sales.SalesArr)

	for _, brand := range dataBrands {
		for brandName, total := range brandTotal {
			if brand.BrandName == brandName {
				value, _ := strconv.ParseFloat(brand.DiscountPercent, 32)
				dicsount := float32(value)
				TotalPercent := (total * dicsount) / 100
				fmt.Println("Сумма со скдикой", TotalPercent)
				fmt.Println("Название бренда", brand)
				rbdro := models.RbDTO{
					ID:                   0,
					ContractNumber:       contractNumb,
					StartDate:            rbReq.PeriodFrom,
					EndDate:              rbReq.PeriodTo,
					TypePeriod:           "",
					BrandName:            brandName,
					ProductCode:          "",
					DiscountPercent:      dicsount,
					DiscountAmount:       TotalPercent,
					TotalWithoutDiscount: 0,
					LeasePlan:            0,
					RewardAmount:         0,
					DiscountType:         RB2Name,
				}
				rbDtoSl = append(rbDtoSl, rbdro)

			}
		}
	}

	return rbDtoSl

}

func GetRB3rdType(request models.RBRequest, contracts []models.Contract) ([]models.RbDTO, error) {

	req := models.ReqBrand{
		ClientBin:      request.BIN,
		Beneficiary:    request.ContractorName,
		DateStart:      request.PeriodFrom,
		DateEnd:        request.PeriodTo,
		Type:           "sales",
		TypeValue:      "sku",
		TypeParameters: GetAllProductsSku(contracts),
	}

	sales, err := GetBrandSales(req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("req \n\n%+v\n\n", req)
	fmt.Printf("SALES \n\n%+v\n\n", sales)

	var RBs []models.RbDTO
	fmt.Println("*********************************************")
	for _, contract := range contracts {
		for _, product := range contract.Products {
			total := GetTotalSalesForSku(sales, product.Sku)
			rb := models.RbDTO{
				ID:              contract.ID,
				ContractNumber:  contract.ContractParameters.ContractNumber,
				StartDate:       contract.ContractParameters.StartDate,
				EndDate:         contract.ContractParameters.EndDate,
				ProductCode:     product.Sku,
				DiscountPercent: product.DiscountPercent,
				LeasePlan:       product.LeasePlan,
				DiscountType:    RB3Name,
			}
			if total >= product.LeasePlan {
				rb.DiscountAmount = total * rb.DiscountPercent / 100
			} else {
				rb.DiscountAmount = 0
			}

			RBs = append(RBs, rb)

		}
	}
	fmt.Println("*********************************************")
	return RBs, nil
}

func GetRB4thType(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
	req := models.ReqBrand{
		ClientBin:   request.BIN,
		Beneficiary: request.ContractorName,
		DateStart:   request.PeriodFrom,
		DateEnd:     request.PeriodTo,
		Type:        "sales_brand_only",
	}

	sales, err := GetSales(req)

	totalAmount := GetTotalAmount(sales)

	log.Printf("[CHECK PRES SAlES: %+v\n", sales)
	log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmount)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == RB4Code && discount.IsSelected == true {
				var discountAmount float32
				if repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code) {
					discountAmount = totalAmount * discount.DiscountPercent / 100
				}
				log.Printf("[CHECK PRES DISCOUNT PERCENT]: %f\n", discount.DiscountPercent)
				log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmount)
				log.Printf("[CHECK PRES DISCOUNT AMOUNT]: %f\n", discountAmount)
				log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))
				rbDTO = append(rbDTO, models.RbDTO{
					ContractNumber:  contract.ContractParameters.ContractNumber,
					StartDate:       request.PeriodTo,
					EndDate:         request.PeriodFrom,
					DiscountPercent: discount.DiscountPercent,
					DiscountAmount:  discountAmount,
					DiscountType:    RB4Name,
				})
				log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
			}
		}
	}

	//rbDTO := []models.RbDTO{
	//	{
	//		ContractNumber:       "9898989211",
	//		StartDate:            "01.01.2022",
	//		EndDate:              "01.02.2022",
	//		BrandName:            "Colgate",
	//		ProductCode:          "00002313",
	//		DiscountPercent:      5,
	//		DiscountAmount:       500,
	//		TotalWithoutDiscount: 100000,
	//	},
	//	{
	//		ContractNumber:       "9898989211",
	//		StartDate:            "01.01.2022",
	//		EndDate:              "01.02.2022",
	//		BrandName:            "Bella",
	//		ProductCode:          "5545454",
	//		DiscountPercent:      10,
	//		DiscountAmount:       500_000,
	//		TotalWithoutDiscount: 5_000_000,
	//	},
	//	{
	//		ContractNumber:       "11255656565",
	//		StartDate:            "01.02.2022",
	//		EndDate:              "01.04.2022",
	//		BrandName:            "Seni",
	//		ProductCode:          "065655",
	//		DiscountPercent:      7,
	//		DiscountAmount:       70_000,
	//		TotalWithoutDiscount: 1_000_000,
	//	},
	//}
	return rbDTO, nil

	//TODO: Доработать от сюда

	////ID, contract_number, discount, bin
	//infoPresentationDiscounts := repository.GetPurchase(rbReq.BIN)
	//
	//layoutISO := "02.1.2006"
	//
	//for _, value := range infoPresentationDiscounts {
	//	if len(infoPresentationDiscounts) > 2 {
	//		//  value.StartDate - эта дата, которую мы взяли из бд
	//		// 01
	//		// 21
	//		// 10
	//		timeDB, err := time.Parse(layoutISO, value.StartDate)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		timeReq, err := time.Parse(layoutISO, value.StartDate)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		if timeDB.Before(timeReq) {
	//			//TODO: например были созданы договора за 21 и за за 25 - а ты берешь дату от 26 числа
	//			//  то да, она будет считать за 21, после этого он будет считать за 25
	//			// но есть нюансы, что сперва выпадает тебе 25 число, а не 21
	//			// ты должен учесть этот момент
	//		}
	//
	//	}
	//}
	//
	////внутри него массив
	//presentationDiscount, err := PresentationDiscount(rbReq)
	//if err != nil {
	//	log.Println(err)
	//	return rbBrand
	//}
	//totalAmount := 0
	//totalWithDiscount := 0
	//
	//for _, value := range presentationDiscount.PurchaseArr {
	//	//rbBrand.ContractNumber = infoPresentationDiscounts.ContractNumber
	//	//rbBrand.StartDate = rbReq.PeriodFrom
	//	//rbBrand.EndDate = rbReq.PeriodTo
	//	//rbBrand.BrandName = value.BrandName
	//	//rbBrand.ProductCode = value.ProductCode
	//	//rbBrand.DiscountPercent = 10
	//
	//	//TODO: подсчет общей суммы
	//	totalAmount += value.Total
	//
	//}
	//totalWithDiscount = (totalAmount * 10) / 100
	//
	////rbBrand.ContractNumber = infoPresentationDiscounts.ContractNumber
	//rbBrand.StartDate = rbReq.PeriodFrom
	//rbBrand.EndDate = rbReq.PeriodTo
	////rbBrand.BrandName = value.BrandName
	////rbBrand.ProductCode = value.ProductCode
	//rbBrand.DiscountPercent = 10
	//rbBrand.TotalWithoutDiscount = float32(totalAmount)
	//rbBrand.DiscountAmount = float32(totalWithDiscount)

	//return rbBrand

}

func GetRB13thType(rb models.RBRequest, contracts []models.Contract) ([]models.RbDTO, error) {
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
					pastTimeFrom, err := utils.ConvertTime(period.PeriodFrom)
					if err != nil {
						return nil, err
					}
					pastTimeTo, err := utils.ConvertTime(period.PeriodTo)
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

func GetRB12thType(req models.RBRequest, contracts []models.Contract) ([]models.RbDTO, error) {
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
