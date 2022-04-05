package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"log"
)

func GetDD1st(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
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

	//totalAmountPurchase := GetTotalAmountPurchase(purchase)

	//	log.Printf("[CHECK PRES SAlES: %+v\n", purchase)
	//log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == DD1Code && discount.IsSelected == true {
				//sales, err := GetSales(req)
				purchase, err := GetPurchase(req)
				if err != nil {
					return nil, err
				}
				totalAmountPurchase := CountPurchaseByCode(purchase)
				var discountAmount float32
				for _, amount := range totalAmountPurchase {
					discountAmount = float32(amount) * discount.DiscountPercent / 100
					rbDTO = append(rbDTO, models.RbDTO{
						ContractNumber:       contract.ContractParameters.ContractNumber,
						StartDate:            request.PeriodTo,
						EndDate:              request.PeriodFrom,
						DiscountPercent:      discount.DiscountPercent,
						DiscountAmount:       discountAmount,
						TotalWithoutDiscount: float32(amount),
						DiscountType:         DD1Name,
					})

				}

			}
			log.Printf("[CHECK PRES DISCOUNT PERCENT]: %f\n", discount.DiscountPercent)
			//log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)

			log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))

			log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
		}
	}

	return rbDTO, nil

}

func GetDD2nd(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
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
	log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == DD2Code && discount.IsSelected == true {
				var discountAmount float32
				for _, amount := range totalAmountPurchase {
					discountAmount = float32(amount) * discount.DiscountPercent / 100
					rbDTO = append(rbDTO, models.RbDTO{
						ContractNumber:       contract.ContractParameters.ContractNumber,
						StartDate:            request.PeriodTo,
						EndDate:              request.PeriodFrom,
						DiscountPercent:      discount.DiscountPercent,
						DiscountAmount:       discountAmount,
						TotalWithoutDiscount: float32(amount),
						DiscountType:         DD2Name,
					})

				}

			}
			log.Printf("[CHECK PRES DISCOUNT PERCENT]: %f\n", discount.DiscountPercent)
			log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)
			log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))

			log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
		}
	}

	return rbDTO, nil

}

func GetDD3rd(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
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
	log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == DD3Code && discount.IsSelected == true {
				var discountAmount float32

				for _, amount := range totalAmountPurchase {
					discountAmount = float32(amount) * discount.DiscountPercent / 100
					rbDTO = append(rbDTO, models.RbDTO{
						ContractNumber:       contract.ContractParameters.ContractNumber,
						StartDate:            request.PeriodTo,
						EndDate:              request.PeriodFrom,
						DiscountPercent:      discount.DiscountPercent,
						DiscountAmount:       discountAmount,
						TotalWithoutDiscount: float32(amount),
						DiscountType:         DD3Name,
					})

				}

			}
			log.Printf("[CHECK PRES DISCOUNT PERCENT]: %f\n", discount.DiscountPercent)
			log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)
			log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))

			log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
		}
	}

	return rbDTO, nil

}

func GetDD4th(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
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
	log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == DD4Code && discount.IsSelected == true {
				var discountAmount float32
				for _, amount := range totalAmountPurchase {
					discountAmount = float32(amount) * discount.DiscountPercent / 100
					rbDTO = append(rbDTO, models.RbDTO{
						ContractNumber:       contract.ContractParameters.ContractNumber,
						StartDate:            request.PeriodTo,
						EndDate:              request.PeriodFrom,
						DiscountPercent:      discount.DiscountPercent,
						DiscountAmount:       discountAmount,
						TotalWithoutDiscount: float32(amount),
						DiscountType:         DD4Name,
					})

				}

			}
			log.Printf("[CHECK PRES DISCOUNT PERCENT]: %f\n", discount.DiscountPercent)
			log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)
			log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))

			log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
		}
	}

	return rbDTO, nil

}

func GetDD5th(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
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
	log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == DD5Code && discount.IsSelected == true {
				var discountAmount float32
				for _, amount := range totalAmountPurchase {
					discountAmount = float32(amount) * discount.DiscountPercent / 100
					rbDTO = append(rbDTO, models.RbDTO{
						ContractNumber:       contract.ContractParameters.ContractNumber,
						StartDate:            request.PeriodTo,
						EndDate:              request.PeriodFrom,
						DiscountPercent:      discount.DiscountPercent,
						DiscountAmount:       discountAmount,
						TotalWithoutDiscount: float32(amount),
						DiscountType:         DD5Name,
					})

				}

			}
			log.Printf("[CHECK PRES DISCOUNT PERCENT]: %f\n", discount.DiscountPercent)
			log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)
			log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))

			log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
		}
	}

	return rbDTO, nil

}

func GetDD6th(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
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
	log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == DD6Code && discount.IsSelected == true {
				var discountAmount float32
				for _, amount := range totalAmountPurchase {
					discountAmount = float32(amount) * discount.DiscountPercent / 100
					rbDTO = append(rbDTO, models.RbDTO{
						ContractNumber:       contract.ContractParameters.ContractNumber,
						StartDate:            request.PeriodTo,
						EndDate:              request.PeriodFrom,
						DiscountPercent:      discount.DiscountPercent,
						DiscountAmount:       discountAmount,
						TotalWithoutDiscount: float32(amount),
						DiscountType:         DD6Name,
					})

				}

			}
			log.Printf("[CHECK PRES DISCOUNT PERCENT]: %f\n", discount.DiscountPercent)
			log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)
			log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))

			log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
		}
	}

	return rbDTO, nil

}
