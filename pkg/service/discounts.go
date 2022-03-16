package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"fmt"
	"log"
	"strings"
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

func SaveDoubtedDiscounts(request model.RBRequest) error {
	for _, discount := range request.DoubtedDiscounts {
		if err := repository.SaveDoubtedDiscounts(request.BIN, request.PeriodFrom, request.PeriodTo, discount.ContractNumber, discount.Discounts); err != nil {
			return err
		}
	}

	return nil
}

// GetAllRBByContractorBIN RB #1
func GetAllRBByContractorBIN(request model.RBRequest) ([]model.RbDTO, error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	//TODO: посмотри потом
	//testBin := "060840003599"
	req := model.ReqBrand{
		ClientBin:   request.BIN,
		Beneficiary: request.ContractorName,
		DateStart:   request.PeriodFrom,
		DateEnd:     request.PeriodTo,
		Type:        "sales",
	}

	sales, err := GetSales(req)

	fmt.Printf("###%+v\n", contracts)
	totalAmount := GetTotalAmount(sales)

	contractRB := DefiningRBReport(contracts, totalAmount)

	return contractRB, nil
}

//func GetAllRBSecondType(request model.RBRequest) ([]model.RbDTO, error) {
//	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
//	if err != nil {
//		return nil, err
//	}
//
//	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
//	if err != nil {
//		return nil, err
//	}
//
//	var RBs []model.RbDTO
//
//	for _, contract := range contracts {
//		for _, discount := range contract.Discounts {
//			if discount.Code == "DISCOUNT_BRAND" && discount.IsSelected {
//				//req := model.ReqBrand{
//				//	ClientBin:   request.BIN,
//				//	Beneficiary: request.ContractorName,
//				//	DateStart:   request.PeriodFrom,
//				//	DateEnd:     request.PeriodTo,
//				//	Type:        "sales",
//				//	TypeValue:   "brand",
//				//}
//				var sales model.Sales
//				if contract.Requisites.BIN == "190241035031" {
//					err = json.Unmarshal([]byte(rb2Mock), &sales)
//				}
//				for _, brand := range contract.DiscountBrand {
//					totalAmount := GetTotalFromSalesByBrand(sales, brand.BrandName)
//					RBs = append(RBs, model.RbDTO{
//						ID:              contract.ID,
//						ContractNumber:  contract.ContractParameters.ContractNumber,
//						StartDate:       contract.ContractParameters.StartDate,
//						EndDate:         contract.ContractParameters.EndDate,
//						DiscountPercent: float32(brand.DiscountPercent),
//						DiscountAmount:  float32(float64(totalAmount) * brand.DiscountPercent / 100),
//					})
//				}
//			}
//		}
//	}
//
//	return RBs, nil
//}

func GetAllRBSecondTypeMock(request model.RBRequest) ([]model.RbDTO, error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	var RBs []model.RbDTO

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == "DISCOUNT_BRAND" && discount.IsSelected {
				//req := model.ReqBrand{
				//	ClientBin:   request.BIN,
				//	Beneficiary: request.ContractorName,
				//	DateStart:   request.PeriodFrom,
				//	DateEnd:     request.PeriodTo,
				//	Type:        "sales",
				//	TypeValue:   "brand",
				//}
				if contract.Requisites.BIN == "190241035031" {
					for _, brand := range contract.DiscountBrand {
						var DiscountAmount float32
						switch brand.BrandName {
						case "7Stick":
							DiscountAmount = 20973100
						case "911":
							DiscountAmount = 41553600
						case "Always":
							DiscountAmount = 34978340
						case "Aura":
							DiscountAmount = 159693
						}

						RBs = append(RBs, model.RbDTO{
							ID:              contract.ID,
							ContractNumber:  contract.ContractParameters.ContractNumber,
							StartDate:       contract.ContractParameters.StartDate,
							EndDate:         contract.ContractParameters.EndDate,
							BrandName:       brand.BrandName,
							DiscountPercent: float32(brand.DiscountPercent),
							DiscountAmount:  DiscountAmount * float32(brand.DiscountPercent) / 100,
						})
					}
					fmt.Printf(">>BRANDS %+v\n", contract.DiscountBrand)
				}
			}
		}
	}

	return RBs, nil
}

// GetAllRBSecondType RB #2
func GetAllRBSecondType(request model.RBRequest) ([]model.RbDTO, error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	var RBs []model.RbDTO

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == "DISCOUNT_BRAND" && discount.IsSelected {
				//req := model.ReqBrand{
				//	ClientBin:   request.BIN,
				//	Beneficiary: request.ContractorName,
				//	DateStart:   request.PeriodFrom,
				//	DateEnd:     request.PeriodTo,
				//	Type:        "sales",
				//	TypeValue:   "brand",
				//}
				if contract.Requisites.BIN == "190241035031" {
					for _, brand := range contract.DiscountBrand {
						var DiscountAmount float32
						switch brand.BrandName {
						case "7Stick":
							DiscountAmount = 20973100
						case "911":
							DiscountAmount = 41553600
						case "Always":
							DiscountAmount = 34978340
						case "Aura":
							DiscountAmount = 159693
						}

						RBs = append(RBs, model.RbDTO{
							ID:              contract.ID,
							ContractNumber:  contract.ContractParameters.ContractNumber,
							StartDate:       contract.ContractParameters.StartDate,
							EndDate:         contract.ContractParameters.EndDate,
							BrandName:       brand.BrandName,
							DiscountPercent: float32(brand.DiscountPercent),
							DiscountAmount:  DiscountAmount * float32(brand.DiscountPercent) / 100,
						})
					}
					fmt.Printf(">>BRANDS %+v\n", contract.DiscountBrand)
				}
			}
		}
	}

	return RBs, nil
}

// GetRBThirdType RB #3
func GetRBThirdType(request model.RBRequest) ([]model.RbDTO, error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	req := model.ReqBrand{
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

	var RBs []model.RbDTO
	fmt.Println("*********************************************")
	for _, contract := range contracts {
		for _, product := range contract.Products {
			total := GetTotalSalesForSku(sales, product.Sku)
			rb := model.RbDTO{
				ID:              contract.ID,
				ContractNumber:  contract.ContractParameters.ContractNumber,
				StartDate:       contract.ContractParameters.StartDate,
				EndDate:         contract.ContractParameters.EndDate,
				ProductCode:     product.Sku,
				DiscountPercent: product.DiscountPercent,
				LeasePlan:       product.LeasePlan,
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

func GetTotalSalesForSku(sales model.Sales, sku string) (totalSum float32) {
	for _, s := range sales.SalesArr {
		if s.ProductCode == sku {
			totalSum += s.Total * s.QntTotal
		}
	}

	return totalSum
}

// GetRBEighthType RB #8
func GetRBEighthType(request model.RBRequest) ([]model.RbDTO, error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	req := model.ReqBrand{
		ClientBin:      request.BIN,
		Beneficiary:    request.ContractorName,
		DateStart:      request.PeriodFrom,
		DateEnd:        request.PeriodTo,
		Type:           "sales",
		TypeValue:      "",
		TypeParameters: nil,
	}

	sales, err := GetBrandSales(req)
	if err != nil {
		return nil, err
	}

	totalAmount := GetTotalAmount(sales)

	fmt.Printf("req \n\n%+v\n\n", req)
	fmt.Printf("SALES \n\n%+v\n\n", sales)

	var RBs []model.RbDTO
	fmt.Println("*********************************************")
	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == "DISCOUNT_FOR_LEASE_GENERAL" && discount.IsSelected == true {
				rb := model.RbDTO{
					ID:              contract.ID,
					ContractNumber:  contract.ContractParameters.ContractNumber,
					StartDate:       contract.ContractParameters.StartDate,
					EndDate:         contract.ContractParameters.EndDate,
					DiscountPercent: discount.DiscountPercent,
					DiscountAmount:  totalAmount * float32(discount.DiscountAmount) / 100,
				}

				RBs = append(RBs, rb)
			}
		}

	}
	fmt.Println("*********************************************")
	return RBs, nil
}

func GetAllProductsSku(contracts []model.Contract) (SkuArr []string) {
	for _, contract := range contracts {
		for _, product := range contract.Products {
			SkuArr = append(SkuArr, product.Sku)
		}
	}

	return SkuArr
}

func GetTotalFromSalesByBrand(sales model.Sales, brand string) (totalAmount float32) {
	for _, s := range sales.SalesArr {
		if s.BrandName == brand {
			totalAmount += s.QntTotal * s.Total
		}
	}

	return totalAmount
}

func DefiningRBReport(contracts []model.Contract, totalAmount float32) (contractsRB []model.RbDTO) {
	for _, contract := range contracts {
		var contractRB []model.RbDTO
		for _, discount := range contract.Discounts {
			if discount.Code == "TOTAL_AMOUNT_OF_SELLING" && discount.IsSelected {
				contractRB = DiscountToReportRB(contract.Discounts[0], contract, totalAmount)
			}
		}
		contractsRB = append(contractsRB, contractRB...)
	}

	return contractsRB
}

func DiscountToReportRB(discount model.Discount, contract model.Contract, totalAmount float32) (contractsRB []model.RbDTO) {
	var contractRB model.RbDTO

	if len(discount.Periods) > 0 {
		contractRB = model.RbDTO{
			ID:             contract.ID,
			ContractNumber: contract.ContractParameters.ContractNumber,
			StartDate:      discount.Periods[0].PeriodFrom,
			EndDate:        discount.Periods[0].PeriodTo,
		}
		if totalAmount >= discount.Periods[0].TotalAmount {
			fmt.Printf("worked [totalAmount = %d AND discount.Periods[0].TotalAmount = %d]\n", totalAmount, discount.Periods[0].TotalAmount)
			contractRB.DiscountAmount = float32(discount.Periods[0].RewardAmount)
		}
	}
	contractsRB = append(contractsRB, contractRB)

	return contractsRB
}

func GetTotalAmount(sales model.Sales) float32 {
	var amount float32
	for _, s := range sales.SalesArr {
		amount += s.Total * s.QntTotal
	}

	return amount
}

func GetTotalAmountFrom1CDataSalesOrPurchases(data []model.GetData1CProducts) float32 {
	var amount float32
	for _, s := range data {
		amount += s.Total * s.QntTotal
	}

	return amount
}

func TrimDate(fullDate string) string {
	arr := strings.Split(fullDate, " ")
	if len(arr) > 0 {
		return arr[0]
	}
	return ""
}

func BulkConvertContractFromJsonB(contractsWithJson []model.ContractWithJsonB) (contracts []model.Contract, err error) {
	for i := range contractsWithJson {
		contract, err := ConvertContractFromJsonB(contractsWithJson[i])
		if err != nil {
			log.Println("Error: service.BulkConvertContractFromJsonB. Error is: ", err.Error())
			continue
		}
		contracts = append(contracts, contract)
	}

	return
}

func GetDoubtedDiscounts(request model.RBRequest) (doubtedDiscounts []model.DoubtedDiscount, err error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	var (
	//hasPresentation bool
	//hasMTZ bool
	)

	for _, contract := range contracts {
		var DoubtedDiscountDetails []model.DoubtedDiscountDetails
		for _, discount := range contract.Discounts {
			var DoubtedDiscountDetail model.DoubtedDiscountDetails
			if (discount.Code == RB4Code || discount.Code == RB11Code) && discount.IsSelected == true {
				DoubtedDiscountDetail.Name = discount.Name
				DoubtedDiscountDetail.Code = discount.Code
				DoubtedDiscountDetail.IsCompleted = repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code)
				DoubtedDiscountDetails = append(DoubtedDiscountDetails, DoubtedDiscountDetail)
			}
		}
		if len(DoubtedDiscountDetails) > 0 {
			doubtedDiscounts = append(doubtedDiscounts, model.DoubtedDiscount{
				ContractNumber: contract.ContractParameters.ContractNumber,
				Discounts:      DoubtedDiscountDetails,
			})
		}
	}

	return doubtedDiscounts, nil
}

// InfoPresentationDiscount RB #4
func InfoPresentationDiscount(request model.RBRequest) (rbDTO []model.RbDTO, err error) {
	//..var rbBrands []model.RbDTO
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	req := model.ReqBrand{
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
				rbDTO = append(rbDTO, model.RbDTO{
					ContractNumber:  contract.ContractParameters.ContractNumber,
					StartDate:       request.PeriodTo,
					EndDate:         request.PeriodFrom,
					DiscountPercent: discount.DiscountPercent,
					DiscountAmount:  discountAmount,
				})
				log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
			}
		}
	}

	//rbDTO := []model.RbDTO{
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

// GetRbTenthType RB #10
func GetRbTenthType(request model.RBRequest) (rbDTO []model.RbDTO, err error) {
	//..var rbBrands []model.RbDTO
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	req := model.ReqBrand{
		ClientBin:   request.BIN,
		Beneficiary: request.ContractorName,
		DateStart:   request.PeriodFrom,
		DateEnd:     request.PeriodTo,
		Type:        "sales",
	}

	sales, err := GetSales(req)

	totalAmount := GetTotalAmount(sales)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == RB10Code && discount.IsSelected == true {
				var discountAmount float32
				if repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code) {
					discountAmount = totalAmount * discount.DiscountPercent / 100
				}
				rbDTO = append(rbDTO, model.RbDTO{
					ContractNumber:  contract.ContractParameters.ContractNumber,
					StartDate:       request.PeriodTo,
					EndDate:         request.PeriodFrom,
					DiscountPercent: discount.DiscountPercent,
					DiscountAmount:  discountAmount,
				})
			}
		}
	}
	return rbDTO, nil
}

func GetRB5thType(request model.RBRequest) (rbDTO []model.RbDTO, err error) {
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
			if discount.Code == RB5Code && discount.IsSelected == true {
				for _, discountBrand := range discount.DiscountBrands {
					if discountBrand.PeriodFrom >= request.PeriodFrom && discountBrand.PeriodTo <= request.PeriodTo {
						req := model.ReqBrand{
							ClientBin:      request.BIN,
							Beneficiary:    request.ContractorName,
							DateStart:      request.PeriodFrom,
							DateEnd:        request.PeriodTo,
							Type:           "sales",
							TypeValue:      "brands",
							TypeParameters: GeAllBrands(discountBrand.Brands),
						}

						sales, err := GetBrandSales(req)
						if err != nil {
							return nil, err
						}

						for _, brand := range discountBrand.Brands {
							totalAmount := GetTotalPurchasesForBrands(sales, brand.BrandName)
							var discountAmount float32
							if totalAmount >= brand.PurchaseAmount {
								discountAmount = totalAmount * brand.DiscountPercent / 100
							}

							rbDTO = append(rbDTO, model.RbDTO{
								ContractNumber:  contract.ContractParameters.ContractNumber,
								StartDate:       discount.PeriodFrom,
								EndDate:         discount.PeriodTo,
								BrandName:       brand.BrandName,
								ProductCode:     brand.BrandCode,
								DiscountPercent: brand.DiscountPercent,
								DiscountAmount:  discountAmount,
								DiscountType:    RB5Name,
							})
						}

					}
				}
			}
		}
	}

	return rbDTO, nil
}

func GetRB6thType(request model.RBRequest) (rbDTO []model.RbDTO, err error) {
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
			if discount.Code == RB6Code && discount.IsSelected == true {
				for _, discountBrand := range discount.DiscountBrands {
					if discountBrand.PeriodFrom >= request.PeriodFrom && discountBrand.PeriodTo <= request.PeriodTo {
						req := model.ReqBrand{
							ClientBin:      request.BIN,
							Beneficiary:    request.ContractorName,
							DateStart:      request.PeriodFrom,
							DateEnd:        request.PeriodTo,
							Type:           "sales",
							TypeValue:      "brands",
							TypeParameters: GeAllBrands(discountBrand.Brands),
						}

						sales, err := GetBrandSales(req)
						if err != nil {
							return nil, err
						}

						for _, brand := range discountBrand.Brands {
							totalAmount := GetTotalPurchasesForBrands(sales, brand.BrandName)
							var discountAmount float32
							if totalAmount >= brand.PurchaseAmount {
								discountAmount = totalAmount * brand.DiscountPercent / 100
							}

							rbDTO = append(rbDTO, model.RbDTO{
								ContractNumber:  contract.ContractParameters.ContractNumber,
								StartDate:       discount.PeriodFrom,
								EndDate:         discount.PeriodTo,
								BrandName:       brand.BrandName,
								ProductCode:     brand.BrandCode,
								DiscountPercent: brand.DiscountPercent,
								DiscountAmount:  discountAmount,
								DiscountType:    RB6Name,
							})
						}

					}
				}
			}
		}
	}

	return rbDTO, nil
}

func GetRB7thType(request model.RBRequest) (rbDTO []model.RbDTO, err error) {
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
			if discount.Code == RB7Code && discount.IsSelected == true {
				for _, discountBrand := range discount.DiscountBrands {
					if discountBrand.PeriodFrom >= request.PeriodFrom && discountBrand.PeriodTo <= request.PeriodTo {
						req := model.ReqBrand{
							ClientBin:      request.BIN,
							Beneficiary:    request.ContractorName,
							DateStart:      request.PeriodFrom,
							DateEnd:        request.PeriodTo,
							Type:           "sales",
							TypeValue:      "brands",
							TypeParameters: GeAllBrands(discountBrand.Brands),
						}

						sales, err := GetBrandSales(req)
						if err != nil {
							return nil, err
						}

						for _, brand := range discountBrand.Brands {
							totalAmount := GetTotalPurchasesForBrands(sales, brand.BrandName)
							var discountAmount float32
							if totalAmount >= brand.PurchaseAmount {
								discountAmount = totalAmount * brand.DiscountPercent / 100
							}

							rbDTO = append(rbDTO, model.RbDTO{
								ContractNumber:  contract.ContractParameters.ContractNumber,
								StartDate:       discount.PeriodFrom,
								EndDate:         discount.PeriodTo,
								BrandName:       brand.BrandName,
								ProductCode:     brand.BrandCode,
								DiscountPercent: brand.DiscountPercent,
								DiscountAmount:  discountAmount,
								DiscountType:    RB7Name,
							})
						}

					}
				}
			}
		}
	}

	return rbDTO, nil
}

func GetTotalPurchasesForBrands(sales model.Sales, brand string) (totalAmount float32) {
	for _, s := range sales.SalesArr {
		if s.BrandCode == brand || s.BrandName == brand {
			totalAmount += s.Total * s.QntTotal
		}
	}

	return totalAmount
}

func GeAllBrands(brandsDTO []model.BrandDTO) (brands []string) {
	for _, brand := range brandsDTO {
		brands = append(brands, brand.BrandCode)
	}

	return brands
}
