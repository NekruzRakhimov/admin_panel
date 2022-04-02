package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"fmt"
	"log"
	"strings"
)

func SaveDoubtedDiscounts(request models.RBRequest) error {
	for _, discount := range request.DoubtedDiscounts {
		if err := repository.SaveDoubtedDiscounts(request.BIN, request.PeriodFrom, request.PeriodTo, discount.ContractNumber, discount.Discounts); err != nil {
			return err
		}
	}

	return nil
}

func GetAllRBByContractorBIN(request models.RBRequest) (rbDTOs []models.RbDTO, err error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return
	}

	// #1
	RB1stType, err := GetRB1stType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, RB1stType...)

	// #2
	rb2ndType := GetRB2ndType(request)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb2ndType...)

	// #3
	rb3rdType, err := GetRB3rdType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb3rdType...)

	// #4
	rbFourthType, err := GetRB4thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rbFourthType...)

	// #5
	rb5thType, err := GetRB5thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb5thType...)

	// #6
	rb6thType, err := GetRB6thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb6thType...)

	// #7
	rb7thType, err := GetRB7thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb7thType...)

	// #8
	rb8thType, err := GetRB8thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb8thType...)

	// #10
	rbTenthType, err := GetRb10thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rbTenthType...)

	// #12
	rb12thType, err := GetRB12thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb12thType...)

	// #13
	rb13thType, err := GetRB13thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb13thType...)

	return
}

func GetTotalSalesForSku(sales models.Sales, sku string) (totalSum float32) {
	for _, s := range sales.SalesArr {
		if s.ProductCode == sku {
			totalSum += s.Total * s.QntTotal
		}
	}

	return totalSum
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

	totalAmount := GetTotalAmount(sales)

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
					DiscountAmount:  totalAmount * float32(discount.DiscountAmount) / 100,
					DiscountType:    RB8Name,
				}

				RBs = append(RBs, rb)
			}
		}

	}
	fmt.Println("*********************************************")
	return RBs, nil
}

func GetAllProductsSku(contracts []models.Contract) (SkuArr []string) {
	for _, contract := range contracts {
		for _, product := range contract.Products {
			SkuArr = append(SkuArr, product.Sku)
		}
	}

	return SkuArr
}

func GetTotalFromSalesByBrand(sales models.Sales, brand string) (totalAmount float32) {
	for _, s := range sales.SalesArr {
		if s.BrandName == brand {
			totalAmount += s.QntTotal * s.Total
		}
	}

	return totalAmount
}

func DefiningRBReport(contracts []models.Contract, totalAmount float32, request models.RBRequest) (contractsRB []models.RbDTO) {
	for _, contract := range contracts {
		var contractRB []models.RbDTO
		for _, discount := range contract.Discounts {
			if discount.Code == "TOTAL_AMOUNT_OF_SELLING" && discount.IsSelected {
				log.Printf("\n[CONTRACT_DISCOUNT][%s] %+v\n", contract.ContractParameters.ContractNumber, contract.Discounts)
				contractRB = DiscountToReportRB(discount, contract, totalAmount, request)
			}
		}
		contractsRB = append(contractsRB, contractRB...)
	}

	return contractsRB
}

func DiscountToReportRB(discount models.Discount, contract models.Contract, totalAmount float32, request models.RBRequest) (contractsRB []models.RbDTO) {
	var contractRB models.RbDTO

	if len(discount.Periods) > 0 {
		contractRB = models.RbDTO{
			ID:             contract.ID,
			ContractNumber: contract.ContractParameters.ContractNumber,
			DiscountType:   RB1Name,
			StartDate:      discount.Periods[0].PeriodFrom,
			EndDate:        discount.Periods[0].PeriodTo,
		}

		var totalDiscountAmount float32
		var totalDiscountRewardAmount int

		for _, period := range discount.Periods {
			if period.PeriodFrom >= request.PeriodFrom && period.PeriodTo <= request.PeriodTo {
				if period.TotalAmount <= totalAmount {
					if period.TotalAmount >= totalDiscountAmount {
						log.Printf("\n[CONTRACT_PERIODS][%s] %+v\n", contract.ContractParameters.ContractNumber, discount.Periods)
						totalDiscountAmount = period.TotalAmount
						totalDiscountRewardAmount = period.RewardAmount
					}
				} else {
					totalDiscountAmount = float32(period.RewardAmount)
				}
			}
		}

		contractRB.RewardAmount = totalDiscountAmount
		contractRB.DiscountAmount = float32(totalDiscountRewardAmount)

		//if len(discount.Periods) > 1 && totalAmount >= discount.Periods[1].TotalAmount && discount.Periods[1].RewardAmount > discount.Periods[0].RewardAmount {
		//	fmt.Printf("worked [totalAmount = %d AND discount.Periods[0].TotalAmount = %d]\n", totalAmount, discount.Periods[0].TotalAmount)
		//	contractRB.DiscountAmount = float32(discount.Periods[0].RewardAmount)
		//} else if totalAmount >= discount.Periods[0].TotalAmount {
		//	fmt.Printf("worked [totalAmount = %d AND discount.Periods[0].TotalAmount = %d]\n", totalAmount, discount.Periods[0].TotalAmount)
		//	contractRB.DiscountAmount = float32(discount.Periods[0].RewardAmount)
		//}
	}
	contractsRB = append(contractsRB, contractRB)

	return contractsRB
}

func GetTotalAmount(sales models.Sales) float32 {
	var amount float32
	for _, s := range sales.SalesArr {
		amount += s.Total
	}

	return amount
}

func GetTotalAmountPurchase(purchase models.Purchase) float32 {
	var amount float32
	for _, s := range purchase.PurchaseArr {
		amount += float32(s.Total)
	}

	return amount
}

func GetTotalAmountFrom1CDataSalesOrPurchases(data []models.GetData1CProducts) float32 {
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

func BulkConvertContractFromJsonB(contractsWithJson []models.ContractWithJsonB) (contracts []models.Contract, err error) {
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

func GetDoubtedDiscounts(request models.RBRequest) (doubtedDiscounts []models.DoubtedDiscount, err error) {
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
		var DoubtedDiscountDetails []models.DoubtedDiscountDetails
		for _, discount := range contract.Discounts {
			var DoubtedDiscountDetail models.DoubtedDiscountDetails
			if (discount.Code == RB4Code || discount.Code == RB11Code) && discount.IsSelected == true {
				DoubtedDiscountDetail.Name = discount.Name
				DoubtedDiscountDetail.Code = discount.Code
				DoubtedDiscountDetail.IsCompleted = repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code)
				DoubtedDiscountDetails = append(DoubtedDiscountDetails, DoubtedDiscountDetail)
			}
		}
		if len(DoubtedDiscountDetails) > 0 {
			doubtedDiscounts = append(doubtedDiscounts, models.DoubtedDiscount{
				ContractNumber: contract.ContractParameters.ContractNumber,
				Discounts:      DoubtedDiscountDetails,
			})
		}
	}

	return doubtedDiscounts, nil
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

func GetRB5thType(request models.RBRequest, contracts []models.Contract) (rbDTO []models.RbDTO, err error) {
	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == RB5Code && discount.IsSelected == true {
				rbDTO, err = RB5Details(request, contract, discount)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return rbDTO, nil
}

func RB5Details(request models.RBRequest, contract models.Contract, discount models.Discount) (rbDTO []models.RbDTO, err error) {
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

func GetTotalPurchasesForBrands(sales models.Sales, brand string) (totalAmount float32) {
	for _, s := range sales.SalesArr {
		if s.BrandCode == brand || s.BrandName == brand {
			totalAmount += s.Total * s.QntTotal
		}
	}

	return totalAmount
}

func GeAllBrands(brandsDTO []models.BrandDTO) (brands []string) {
	for _, brand := range brandsDTO {
		brands = append(brands, brand.BrandName)
	}

	return brands
}
