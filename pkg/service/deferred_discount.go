package service

import (
	"admin_panel/models"
	"log"
)

func Fill(request models.RBRequest, contract models.Contract) (float64, error) {
	req := models.ReqBrand{
		ClientCode:  request.ClientCode,
		Beneficiary: request.ContractorName,
		DateStart:   request.PeriodFrom,
		DateEnd:     request.PeriodTo,
		SchemeType:  contract.View,
	}
	purchase, err := GetPurchase(req)
	if err != nil {
		return 0, err
	}
	totalAmount := CountPurchase(purchase)

	return totalAmount, nil

}

func GetDD1st(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == DD1Code && discount.IsSelected == true {
				totalAmount, err := Fill(request, contract)
				if err != nil {
					return nil, err
				}

				var discountAmount float32
				//for _, amount := range totalAmountPurchase {
				discountAmount = float32(totalAmount) * discount.DiscountPercent / 100
				rbDTO = append(rbDTO, models.RbDTO{
					ContractNumber:       contract.ContractParameters.ContractNumber,
					StartDate:            request.PeriodFrom,
					EndDate:              request.PeriodTo,
					DiscountPercent:      discount.DiscountPercent,
					DiscountAmount:       discountAmount,
					TotalWithoutDiscount: float32(totalAmount),
					DiscountType:         DD1Name,
				})

			}

		}
		log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
	}

	return rbDTO, nil

}

func GetDD2nd(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == DD2Code && discount.IsSelected == true {
				totalAmount, err := Fill(request, contract)
				if err != nil {
					return nil, err
				}
				var discountAmount float32
				//for _, amount := range totalAmountPurchase {
				discountAmount = float32(totalAmount) * discount.DiscountPercent / 100
				rbDTO = append(rbDTO, models.RbDTO{
					ContractNumber:       contract.ContractParameters.ContractNumber,
					StartDate:            request.PeriodTo,
					EndDate:              request.PeriodFrom,
					DiscountPercent:      discount.DiscountPercent,
					DiscountAmount:       discountAmount,
					TotalWithoutDiscount: float32(totalAmount),
					DiscountType:         DD2Name,
				})

			}

		}
		//log.Printf("[CHECK PRES DISCOUNT PERCENT]: %f\n", discount.DiscountPercent)
		//log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)
		//log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))

		log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
	}
	//}

	return rbDTO, nil

}

func GetDD3rd(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == DD3Code && discount.IsSelected == true {
				totalAmount, err := Fill(request, contract)
				if err != nil {
					return nil, err
				}
				var discountAmount float32
				discountAmount = float32(totalAmount) * discount.DiscountPercent / 100
				rbDTO = append(rbDTO, models.RbDTO{
					ContractNumber:       contract.ContractParameters.ContractNumber,
					StartDate:            request.PeriodTo,
					EndDate:              request.PeriodFrom,
					DiscountPercent:      discount.DiscountPercent,
					DiscountAmount:       discountAmount,
					TotalWithoutDiscount: float32(totalAmount),
					DiscountType:         DD3Name,
				})

			}

		}
		log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
	}

	return rbDTO, nil

}

func GetDD4th(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == DD4Code && discount.IsSelected == true {
				totalAmount, err := Fill(request, contract)
				if err != nil {
					return nil, err
				}
				var discountAmount float32
				discountAmount = float32(totalAmount) * discount.DiscountPercent / 100
				rbDTO = append(rbDTO, models.RbDTO{
					ContractNumber:       contract.ContractParameters.ContractNumber,
					StartDate:            request.PeriodTo,
					EndDate:              request.PeriodFrom,
					DiscountPercent:      discount.DiscountPercent,
					DiscountAmount:       discountAmount,
					TotalWithoutDiscount: float32(totalAmount),
					DiscountType:         DD4Name,
				})

			}

		}

		log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
	}

	return rbDTO, nil

}

func GetDD5th(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == DD5Code && discount.IsSelected == true {
				totalAmount, err := Fill(request, contract)
				if err != nil {
					return nil, err
				}

				var discountAmount float32

				discountAmount = float32(totalAmount) * discount.DiscountPercent / 100
				rbDTO = append(rbDTO, models.RbDTO{
					ContractNumber:       contract.ContractParameters.ContractNumber,
					StartDate:            request.PeriodTo,
					EndDate:              request.PeriodFrom,
					DiscountPercent:      discount.DiscountPercent,
					DiscountAmount:       discountAmount,
					TotalWithoutDiscount: float32(totalAmount),
					DiscountType:         DD5Name,
				})

			}

		}

		//log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))

		log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
	}

	return rbDTO, nil

}

func GetDD6th(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == DD6Code && discount.IsSelected == true {
				totalAmount, err := Fill(request, contract)
				if err != nil {
					return nil, err
				}

				var discountAmount float32
				discountAmount = float32(totalAmount) * discount.DiscountPercent / 100
				rbDTO = append(rbDTO, models.RbDTO{
					ContractNumber:       contract.ContractParameters.ContractNumber,
					StartDate:            request.PeriodTo,
					EndDate:              request.PeriodFrom,
					DiscountPercent:      discount.DiscountPercent,
					DiscountAmount:       discountAmount,
					TotalWithoutDiscount: float32(totalAmount),
					DiscountType:         DD6Name,
				})

			}

		}

		//log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))

		log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
	}

	return rbDTO, nil

}
