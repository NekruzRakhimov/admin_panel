package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"admin_panel/utils"
	"fmt"
	"log"
	"time"
)

const (
	sheet    = "Sheet1"
	RB1Name  = "Скидка за объем закупа"
	RB2Name  = "Скидка  на группы товаров"
	RB3Name  = "Скидка за выполнение плана закупа по препаратам"
	RB4Name  = "Скидка за представленность"
	RB5Name  = "Скидка за фиксированную сумму рб за выполнение плана закупа по брендам"
	RB6Name  = "Скидка на МТЗ"
	RB7Name  = "Скидка за выполнение плана продаж по бренду"
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
	RB7Code  = "DISCOUNT_FOR_FULFILLING_SALES"
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

	externalCodes := GetExternalCode(request.BIN)
	contractsCode := JoinContractCode(externalCodes)

	reqBrand := models.ReqBrand{
		ClientBin:      request.BIN,
		DateStart:      request.PeriodFrom,
		DateEnd:        request.PeriodTo,
		TypeValue:      "",
		TypeParameters: nil,
		Contracts:      contractsCode, // необходимо получить коды контрактов
	}
	purchase, _ := GetPurchase(reqBrand)
	totalAmount := GetPurchaseTotalAmount(purchase)

	//totalPurchaseCode := CountPurchaseByCode(purchase)
	//
	//present := models.ReqBrand{
	//	ClientBin:      request.BIN,
	//	Beneficiary:    "",
	//	DateStart:      request.PeriodFrom,
	//	DateEnd:        request.PeriodTo,
	//	Type:           "",
	//	TypeValue:      "",
	//	TypeParameters: nil,
	//	Contracts:      nil,
	//}
	//
	//sales, err := GetSales1C(present, "sales_brand_only")
	//sales, err := GetSales(req)
	//if err != nil {
	//	return nil, err
	//}

	fmt.Printf("###%+v\n", contracts)
	log.Printf("[PURCHASE] %f ", totalAmount)

	contractRB := DefiningRBReport(contracts, totalAmount, request)

	return contractRB, nil
}

func GetPurchaseTotalAmount(purchases models.Purchase) (totalAmount float32) {
	for _, purchase := range purchases.PurchaseArr {
		totalAmount += float32(purchase.Total)
	}

	return totalAmount
}

//func GetRB2ndType(rbReq models.RBRequest) []models.RbDTO {
//	brandTotal := map[string]float32{}
//	var rbDtoSl []models.RbDTO
//
//	rbBrand := models.ReqBrand{
//		ClientBin: rbReq.BIN,
//		DateStart: rbReq.PeriodFrom,
//		DateEnd:   rbReq.PeriodTo,
//	}
//
//	// берем бренды и их Total // общую сумму не зависимо от договора
//	sales, _ := GetSales(rbBrand)
//
//	// тут считаем общую сумму каждого бренда
//	for _, sale := range sales.SalesArr {
//		// считаем общую сумму по брендам, и чтобы они не дублировались
//		brandTotal[sale.BrandName] += sale.Total
//	}
//
//	// берем скидки по брендам и название брендов
//	dataBrands, err := repository.GetIDByBIN(rbReq.BIN)
//	if err != nil {
//		return nil
//	}
//	fmt.Println("dataBrand", dataBrands)
//
//	for brandName, total := range brandTotal {
//
//		for _, brand := range dataBrands {
//			// сравниваем бренды, то есть если бин - 160140011654- то у него всего 2 бренда
//			//[Sante:1579 Silver Care:19410]
//			if brand.Brand == brandName {
//				value, _ := strconv.ParseFloat(brand.DiscountPercent, 32)
//				dicsount := float32(value)
//				TotalPercent := (total * dicsount) / 100
//				rbdro := models.RbDTO{
//					ID:                   0,
//					ContractNumber:       brand.ContractNumber,
//					StartDate:            rbReq.PeriodFrom,
//					EndDate:              rbReq.PeriodTo,
//					TypePeriod:           "",
//					BrandName:            brandName,
//					ProductCode:          "",
//					DiscountPercent:      dicsount,
//					DiscountAmount:       TotalPercent,
//					TotalWithoutDiscount: 0,
//					LeasePlan:            0,
//					RewardAmount:         0,
//					DiscountType:         RB2Name,
//				}
//				rbDtoSl = append(rbDtoSl, rbdro)
//
//			}
//		}
//	}
//	//}
//
//	return rbDtoSl
//
//}

//func GetRB2ndType(rbReq models.RBRequest) []models.RbDTO {
//	brandTotal := map[string]float32{}
//	var rbDtoSl []models.RbDTO

func GetRB2ndType(rb models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO) {
	req := models.ReqBrand{
		ClientBin: rb.BIN,
		DateStart: rb.PeriodFrom,
		DateEnd:   rb.PeriodTo,
	}
	sales, _ := GetSales(req)
	mapBrands := CountSalesByBrand(sales)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == RB2Code && discount.IsSelected == true {

				for _, discountBrand := range discount.DiscountBrands {
					//   01.01.2022                01.01.2022       31.03.2022                 <= 31.03.2022
					//if discountBrand.PeriodFrom >= rb.PeriodFrom && discountBrand.PeriodTo <= rb.PeriodTo {
					for _, dataBrand := range discountBrand.Brands {
						for brand, total := range mapBrands {
							if brand == dataBrand.BrandName {
								var discountAmount float32
								//if total >= dataBrand.PurchaseAmount {
								discountAmount = total * dataBrand.DiscountPercent / 100

								rbDTO = append(rbDTO, models.RbDTO{
									ContractNumber:       contract.ContractParameters.ContractNumber,
									StartDate:            rb.PeriodFrom,
									EndDate:              rb.PeriodTo,
									BrandName:            dataBrand.BrandName,
									ProductCode:          dataBrand.BrandCode,
									DiscountPercent:      dataBrand.DiscountPercent,
									TotalWithoutDiscount: total,
									DiscountAmount:       discountAmount,
									DiscountType:         RB2Name,
								})
							}

						}

					}
				}
			}
		}
	}
	//}

	return rbDTO
}

func GetRB3rdType(request models.RBRequest, contracts []models.Contract) ([]models.RbDTO, error) {

	//req := models.ReqBrand{
	//	ClientBin:      request.BIN,
	//	Beneficiary:    request.ContractorName,
	//	DateStart:      request.PeriodFrom,
	//	DateEnd:        request.PeriodTo,
	//	Type:           "sales",
	//	TypeValue:      "sku",
	//	TypeParameters: GetAllProductsSku(contracts),
	//}
	//
	//sales, err := GetBrandSales(req)
	//if err != nil {
	//	return nil, err
	//}

	externalCodes := GetExternalCode(request.BIN)
	contractsCode := JoinContractCode(externalCodes)

	req := models.ReqBrand{
		ClientBin:      request.BIN,
		DateStart:      request.PeriodFrom,
		DateEnd:        request.PeriodTo,
		TypeValue:      "",
		TypeParameters: nil,
		Contracts:      contractsCode, // необходимо получить коды контрактов
	}
	purchase, _ := GetPurchase(req)
	totalAmount := GetPurchaseTotalAmount(purchase)

	//fmt.Printf("req \n\n%+v\n\n", req)
	//fmt.Printf("SALES \n\n%+v\n\n", sales)

	var RBs []models.RbDTO
	fmt.Println("*********************************************")
	for _, contract := range contracts {
		for _, product := range contract.Products {
			//total := GetTotalSalesForSku(sales, product.Sku)
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
			if totalAmount >= product.LeasePlan {
				rb.DiscountAmount = totalAmount * rb.DiscountPercent / 100
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
	externalCodes := GetExternalCode(request.BIN)
	contractsCode := JoinContractCode(externalCodes)
	//var contractsCode []string
	//for _, value := range externalCodes {
	//	contractsCode = append(contractsCode, value.ExtContractCode)
	//}

	req := models.ReqBrand{
		ClientBin:   request.BIN,
		Beneficiary: request.ContractorName,
		DateStart:   request.PeriodFrom,
		DateEnd:     request.PeriodTo,
		Contracts:   contractsCode,
	}

	//sales, err := GetSales(req)
	purchase, err := GetPurchase(req)
	totalAmountPurchase := CountPurchaseByCode(purchase)

	//totalAmountPurchase := GetTotalAmountPurchase(purchase)

	log.Printf("[CHECK PRES SAlES: %+v\n", purchase)
	log.Printf("[CHECK PRES TOTAL AMOUNT]: %v\n", totalAmountPurchase)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == RB4Code && discount.IsSelected == true {
				fmt.Println("Условия прошли")
				var discountAmount float32
				if repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code) {
					for _, amount := range totalAmountPurchase {
						if amount > 0 {
							discountAmount = float32(amount) * discount.DiscountPercent / 100
							rbDTO = append(rbDTO, models.RbDTO{
								ContractNumber:       contract.ContractParameters.ContractNumber,
								StartDate:            request.PeriodTo,
								EndDate:              request.PeriodFrom,
								DiscountPercent:      discount.DiscountPercent,
								DiscountAmount:       discountAmount,
								TotalWithoutDiscount: float32(amount),
								DiscountType:         RB4Name,
							})

						} else {
							rbDTO = append(rbDTO, models.RbDTO{
								ContractNumber:       contract.ContractParameters.ContractNumber,
								StartDate:            request.PeriodTo,
								EndDate:              request.PeriodFrom,
								DiscountPercent:      discount.DiscountPercent,
								DiscountAmount:       0,
								TotalWithoutDiscount: float32(amount),
								DiscountType:         RB4Name,
							})
						}
						log.Printf("[CHECK PRES DISCOUNT AMOUNT]: %v\n", discountAmount)

					}
					log.Printf("[CHECK PRES DISCOUNT PERCENT]: %f\n", discount.DiscountPercent)
					log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)

					log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))

					log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
				}
			}
		}
	}
	return rbDTO, nil

}

func GetRB5thType(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == RB5Code && discount.IsSelected == true {
				rbDTO, err = RB5thTypeDetails(request, contract, discount)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return rbDTO, nil
}

func RB5thTypeDetails(request models.RBRequest, contract models.Contract, discount models.Discount) (rbDTO []models.RbDTO, err error) {
	log.Printf("\n[DISCOUNT_DETAILS] %+v\n", discount)
	for _, discountBrand := range discount.DiscountBrands {
		if discountBrand.PeriodFrom >= request.PeriodFrom && discountBrand.PeriodTo <= request.PeriodTo {
			//req := models.ReqBrand{
			//	ClientBin:      request.BIN,
			//	Beneficiary:    request.ContractorName,
			//	DateStart:      request.PeriodFrom,
			//	DateEnd:        request.PeriodTo,
			//	Type:           "sales",
			//	TypeValue:      "brands",
			//	TypeParameters: GeAllBrands(discountBrand.Brands),
			//}
			//
			//sales, err := GetBrandSales(req)
			//if err != nil {
			//	return nil, err
			//}

			externalCodes := GetExternalCode(request.BIN)
			contractsCode := JoinContractCode(externalCodes)

			reqBrand := models.ReqBrand{
				ClientBin:      request.BIN,
				DateStart:      request.PeriodFrom,
				DateEnd:        request.PeriodTo,
				TypeValue:      "",
				TypeParameters: nil,
				Contracts:      contractsCode, // необходимо получить коды контрактов
			}
			purchase, _ := GetPurchase(reqBrand)
			totalAmount := GetPurchaseTotalAmount(purchase)

			for _, brand := range discountBrand.Brands {

				//totalAmount := GetTotalPurchasesForBrands(sales, brand.BrandName)
				var discountAmount float32
				if totalAmount >= brand.PurchaseAmount {
					discountAmount = totalAmount * brand.DiscountPercent / 100
				}

				rbDTO = append(rbDTO, models.RbDTO{
					ContractNumber:  contract.ContractParameters.ContractNumber,
					StartDate:       discountBrand.PeriodFrom,
					EndDate:         discountBrand.PeriodTo,
					BrandName:       brand.BrandName,
					ProductCode:     brand.BrandCode,
					DiscountPercent: brand.DiscountPercent,
					DiscountAmount:  discountAmount,
					DiscountType:    RB5Name,
				})
			}
		}
	}

	return rbDTO, nil
}

func GetRB6thType(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == RB6Code && discount.IsSelected == true {
				for _, discountBrand := range discount.DiscountBrands {
					if discountBrand.PeriodFrom >= request.PeriodFrom && discountBrand.PeriodTo <= request.PeriodTo {
						req := models.ReqBrand{
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

							rbDTO = append(rbDTO, models.RbDTO{
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

func GetRB7thType(rb models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
	req := models.ReqBrand{
		ClientBin: rb.BIN,
		DateStart: rb.PeriodFrom,
		DateEnd:   rb.PeriodTo,
	}
	sales, err := GetSales(req)
	mapBrands := CountSalesByBrand(sales)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == RB7Code && discount.IsSelected == true {

				for _, discountBrand := range discount.DiscountBrands {
					//   01.01.2022                01.01.2022       31.03.2022                 <= 31.03.2022
					if discountBrand.PeriodFrom >= rb.PeriodFrom && discountBrand.PeriodTo <= rb.PeriodTo {
						for _, dataBrand := range discountBrand.Brands {
							for brand, total := range mapBrands {
								if brand == dataBrand.BrandName {
									var discountAmount float32
									if total >= dataBrand.PurchaseAmount {
										discountAmount = total * dataBrand.DiscountPercent / 100
									}
									rbDTO = append(rbDTO, models.RbDTO{
										ContractNumber:       contract.ContractParameters.ContractNumber,
										StartDate:            discountBrand.PeriodFrom,
										EndDate:              discountBrand.PeriodTo,
										BrandName:            dataBrand.BrandName,
										ProductCode:          dataBrand.BrandCode,
										DiscountPercent:      dataBrand.DiscountPercent,
										TotalWithoutDiscount: total,
										DiscountAmount:       discountAmount,
										DiscountType:         RB7Name,
									})
								}

							}

						}
					}
				}
			}
		}
	}

	return rbDTO, nil
}

func GetRB8thType(request models.RBRequest, contracts []models.Contract) ([]models.RbDTO, error) {
	//req := models.ReqBrand{
	//	ClientBin:      request.BIN,
	//	Beneficiary:    request.ContractorName,
	//	DateStart:      request.PeriodFrom,
	//	DateEnd:        request.PeriodTo,
	//	Type:           "sales",
	//	TypeValue:      "",
	//	TypeParameters: nil,
	//}
	//
	//sales, err := GetBrandSales(req)
	//if err != nil {
	//	return nil, err
	//}

	//present := models.ReqBrand{
	//	ClientBin:      request.BIN,
	//	Beneficiary:    "",
	//	DateStart:      request.PeriodFrom,
	//	DateEnd:        request.PeriodTo,
	//	Type:           "",
	//	TypeValue:      "",
	//	TypeParameters: nil,
	//	Contracts:      nil,
	//}
	//
	//sales, err := GetSales1C(present, "sales_brand_only")
	////sales, err := GetSales(req)
	//if err != nil {
	//	return nil, err
	//}
	//
	//totalAmount := GetTotalAmount(sales)

	externalCodes := GetExternalCode(request.BIN)
	contractsCode := JoinContractCode(externalCodes)

	reqBrand := models.ReqBrand{
		ClientBin:      request.BIN,
		DateStart:      request.PeriodFrom,
		DateEnd:        request.PeriodTo,
		TypeValue:      "",
		TypeParameters: nil,
		Contracts:      contractsCode, // необходимо получить коды контрактов
	}
	purchase, _ := GetPurchase(reqBrand)
	totalAmount := GetPurchaseTotalAmount(purchase)

	var RBs []models.RbDTO
	fmt.Println("*********************************************")
	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == "DISCOUNT_FOR_LEASE_GENERAL" && discount.IsSelected == true {
				rb := models.RbDTO{
					ID:              contract.ID,
					ContractNumber:  contract.ContractParameters.ContractNumber,
					StartDate:       contract.ContractParameters.StartDate,
					EndDate:         contract.ContractParameters.EndDate,
					DiscountPercent: discount.DiscountPercent,
					DiscountAmount:  totalAmount * discount.DiscountPercent / 100,
					DiscountType:    RB8Name,
				}

				RBs = append(RBs, rb)
			}
		}
	}
	fmt.Println("*********************************************")
	return RBs, nil
}

func GetRB9thType(request models.RBRequest, contracts []models.Contract) ([]models.RbDTO, error) {
	//req := models.ReqBrand{
	//	ClientBin:      request.BIN,
	//	Beneficiary:    request.ContractorName,
	//	DateStart:      request.PeriodFrom,
	//	DateEnd:        request.PeriodTo,
	//	Type:           "sales",
	//	TypeValue:      "",
	//	TypeParameters: nil,
	//}
	//
	//sales, err := GetBrandSales(req)
	//if err != nil {
	//	return nil, err
	//}

	//present := models.ReqBrand{
	//	ClientBin:      request.BIN,
	//	Beneficiary:    "",
	//	DateStart:      request.PeriodFrom,
	//	DateEnd:        request.PeriodTo,
	//	Type:           "",
	//	TypeValue:      "",
	//	TypeParameters: nil,
	//	Contracts:      nil,
	//}
	//
	//sales, err := GetSales1C(present, "sales_brand_only")
	////sales, err := GetSales(req)
	//if err != nil {
	//	return nil, err
	//}
	//
	//totalAmount := GetTotalAmount(sales)

	externalCodes := GetExternalCode(request.BIN)
	contractsCode := JoinContractCode(externalCodes)

	reqBrand := models.ReqBrand{
		ClientBin:      request.BIN,
		DateStart:      request.PeriodFrom,
		DateEnd:        request.PeriodTo,
		TypeValue:      "",
		TypeParameters: nil,
		Contracts:      contractsCode, // необходимо получить коды контрактов
	}
	purchase, _ := GetPurchase(reqBrand)
	totalAmount := GetPurchaseTotalAmount(purchase)

	var RBs []models.RbDTO
	fmt.Println("*********************************************")
	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == RB9Code && discount.IsSelected == true {
				rb := models.RbDTO{
					ID:              contract.ID,
					ContractNumber:  contract.ContractParameters.ContractNumber,
					StartDate:       contract.ContractParameters.StartDate,
					EndDate:         contract.ContractParameters.EndDate,
					DiscountPercent: discount.DiscountPercent,
					DiscountAmount:  totalAmount * discount.DiscountPercent / 100,
					DiscountType:    RB9Name,
				}

				RBs = append(RBs, rb)
			}
		}
	}
	fmt.Println("*********************************************")
	return RBs, nil
}

func GetRb10thType(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
	req := models.ReqBrand{
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
				rbDTO = append(rbDTO, models.RbDTO{
					ContractNumber:  contract.ContractParameters.ContractNumber,
					StartDate:       request.PeriodTo,
					EndDate:         request.PeriodFrom,
					DiscountPercent: discount.DiscountPercent,
					DiscountAmount:  discountAmount,
					DiscountType:    RB10Name,
				})
			}
		}
	}
	return rbDTO, nil
}

func GetRB12thType(req models.RBRequest, contracts []models.Contract) ([]models.RbDTO, error) {
	fmt.Println("====================вызов функции =========================================")
	//totalbyCode := map[string]int{}

	var rbDTOsl []models.RbDTO

	// parsing string by TIME
	//layoutISO := "02.1.2006"

	// parsing string to Time
	//reqPeriodFrom, _ := time.Parse(layoutISO, req.PeriodFrom)
	//reqPeriodTo, _ := time.Parse(layoutISO, req.PeriodTo)

	// get all contracts_code by BIN
	externalCodes := GetExternalCode(req.BIN)
	contractsCode := JoinContractCode(externalCodes)

	for _, contract := range contracts {
		fmt.Println("contract MESSAGE", contract.Discounts)
		for _, discount := range contract.Discounts {
			if discount.Code == "RB_DISCOUNT_FOR_PURCHASE_PERIOD" { // здесь сравниваешь тип скидки и берешь тот тип который тебе нужен
				for _, period := range discount.Periods {
					if period.PeriodFrom >= req.PeriodFrom && period.PeriodTo <= req.PeriodTo {
						//PeriodFrom, _ := time.Parse(layoutISO, period.PeriodFrom)
						//PeriodTo, _ := time.Parse(layoutISO, period.PeriodTo)
						//if PeriodFrom.After(reqPeriodFrom) || PeriodTo.Before(reqPeriodTo) {
						reqBrand := models.ReqBrand{
							ClientBin:      req.BIN,
							DateStart:      req.PeriodFrom,
							DateEnd:        req.PeriodTo,
							TypeValue:      "",
							TypeParameters: nil,
							Contracts:      contractsCode, // необходимо получить коды контрактов
						}
						purchase, _ := GetPurchase(reqBrand)

						totalPurchaseCode := CountPurchaseByCode(purchase)

						for _, amount := range totalPurchaseCode {
							fmt.Println("period.PurchaseAmount", period.PurchaseAmount)
							if float32(amount) > period.PurchaseAmount {
								total := float32(amount) * period.DiscountPercent / 100
								RbDTO := models.RbDTO{
									ContractNumber:       contract.ContractParameters.ContractNumber,
									StartDate:            period.PeriodFrom,
									EndDate:              period.PeriodTo,
									TypePeriod:           period.Type,
									DiscountPercent:      period.DiscountPercent,
									DiscountAmount:       total,
									TotalWithoutDiscount: float32(amount),
									LeasePlan:            period.PurchaseAmount,
									DiscountType:         RB12Name,
								}
								rbDTOsl = append(rbDTOsl, RbDTO)

							} else {
								RbDTO := models.RbDTO{
									ContractNumber:       contract.ContractParameters.ContractNumber,
									StartDate:            period.PeriodFrom,
									EndDate:              period.PeriodTo,
									TypePeriod:           period.Type,
									DiscountPercent:      period.DiscountPercent,
									DiscountAmount:       0.0,
									TotalWithoutDiscount: float32(amount), // эта сумма, котору. мы получаем от 1С
									LeasePlan:            period.PurchaseAmount,
									DiscountType:         RB12Name,
								}
								rbDTOsl = append(rbDTOsl, RbDTO)
							}

						}

					}

				}

			}

		}
	}
	//}

	return rbDTOsl, nil
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
			if discount.Code == RB13Code && discount.IsSelected == true {

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

					if presentPeriod.SalesArr == nil || oldPeriod.SalesArr == nil {
						rbDTOsl, _ = CheckPeriodNullGrowth(contract, period, RB13Name)
						//rbDTO := models.RbDTO{
						//	ContractNumber:       contract.ContractParameters.ContractNumber,
						//	StartDate:            period.PeriodFrom,
						//	EndDate:              period.PeriodTo,
						//	TypePeriod:           "",
						//	BrandName:            "",
						//	ProductCode:          "",
						//	DiscountPercent:      period.DiscountPercent,
						//	DiscountAmount:       0,
						//	TotalWithoutDiscount: 0,
						//	DiscountType:         RB13Name,
						//}
						//rbDTOsl = append(rbDTOsl, rbDTO)
						return rbDTOsl, nil

					}

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
						rbDTOsl, _ = CheckPeriodNullGrowth(contract, period, RB13Name)

						//rbDTO := models.RbDTO{
						//	ContractNumber:       contract.ContractParameters.ContractNumber,
						//	StartDate:            period.PeriodFrom,
						//	EndDate:              period.PeriodTo,
						//	TypePeriod:           "",
						//	BrandName:            "",
						//	ProductCode:          "",
						//	DiscountPercent:      period.DiscountPercent,
						//	DiscountAmount:       0,
						//	TotalWithoutDiscount: preCount,
						//	DiscountType:         RB13Name,
						//}
						//rbDTOsl = append(rbDTOsl, rbDTO)
					}

				}

			}

		}
	}
	//}
	return rbDTOsl, nil
}

func CheckPeriodNullGrowth(contract models.Contract, period models.DiscountPeriod, discountType string) ([]models.RbDTO, error) {
	var rbDTOsl []models.RbDTO

	rbDTO := models.RbDTO{
		ContractNumber:       contract.ContractParameters.ContractNumber,
		StartDate:            period.PeriodFrom,
		EndDate:              period.PeriodTo,
		DiscountPercent:      period.DiscountPercent,
		DiscountAmount:       0,
		TotalWithoutDiscount: 0,
		DiscountType:         discountType,
	}
	rbDTOsl = append(rbDTOsl, rbDTO)
	return rbDTOsl, nil

}

// CountPurchaseByCode считываем итог по каждому контракт коду
func CountPurchaseByCode(purchase models.Purchase) map[string]float64 {
	totallyCode := map[string]float64{}
	for _, value := range purchase.PurchaseArr {
		if _, ok := totallyCode[value.ContractCode]; !ok {
			totallyCode[value.ContractCode] += value.Total
			//do something here
		}

	}
	return totallyCode
}

// CountSalesByBrand считываем итог по каждому контракт коду
func CountSalesByBrand(sales models.Sales) map[string]float32 {
	totallyCode := map[string]float32{}
	for _, value := range sales.SalesArr {
		//if _, ok := totallyCode[value.ContractCode]; !ok {
		totallyCode[value.BrandName] += value.Total
		//do something here
		//}

	}
	return totallyCode
}

// JoinContractCode собираем все контракт коды в слайс
func JoinContractCode(externalCodes []models.ContractCode) []string {

	var contractsCode []string
	for _, value := range externalCodes {
		fmt.Println("value.ExtContractCode===========================================================================", value.ExtContractCode)

		if value.ExtContractCode == "" {
			continue
		}
		contractsCode = append(contractsCode, value.ExtContractCode)
	}
	return contractsCode
}
