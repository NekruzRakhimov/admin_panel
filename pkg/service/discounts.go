package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

const (
	sheet  = "Sheet1"
	sheet2 = "Скидка за объем закупа"
	sheet3 = "Скидка  на группы товаров"
	sheet4 = "Скидка за выполнение плана закупа по препаратам"
	sheet5 = "Скидка за представленность"
)

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

func FormExcelForRBReport(request model.RBRequest) error {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		fmt.Println(">> 1")
		return err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		fmt.Println(">> 2")
		return err
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

	brandInfo := []model.BrandInfo{}
	sales, err := GetSalesBrand(req, brandInfo)
	if err != nil {
		fmt.Println(">> 3")
		fmt.Println(err.Error())
		return err
	}

	var isRB1 bool
	var isRB2 bool
	var isRB3 bool
	var isRB4 bool

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == "TOTAL_AMOUNT_OF_SELLING" && discount.IsSelected {
				isRB1 = true
			}
			if discount.Code == "DISCOUNT_BRAND" && discount.IsSelected {
				isRB2 = true
			}
			if discount.Code == "DISCOUNT_PLAN_LEASE" && discount.IsSelected {
				isRB3 = true
			}
			if discount.Code == "DISCOUNT_FOR_REPRESENTATION" && discount.IsSelected {
				isRB4 = true
			}
		}
	}

	totalAmount := GetTotalAmount(sales)

	fmt.Println(contracts)
	fmt.Println(totalAmount)
	var conTotalAmount float32
	var rewardAmount int
	if len(contracts) > 0 {
		for _, discount := range contracts[0].Discounts {
			if discount.Code == "TOTAL_AMOUNT_OF_SELLING" && discount.IsSelected == true {
				if len(contracts[0].Discounts[0].Periods) > 0 {
					conTotalAmount = contracts[0].Discounts[0].Periods[0].TotalAmount
					rewardAmount = contracts[0].Discounts[0].Periods[0].RewardAmount
				}
			}
		}
		if len(contracts[0].Discounts) > 0 {

		}
	}

	f := excelize.NewFile()
	//if err != nil {
	//	return err
	//}
	var discount int
	if conTotalAmount <= totalAmount {
		discount = rewardAmount
	}
	style, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#F5DEB3"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	f.SetCellValue(sheet, "A1", "Номенклатура")
	f.SetCellValue(sheet, "B1", "Номер продукта")
	f.SetCellValue(sheet, "C1", "Стоимость")
	f.SetCellValue(sheet, "D1", "Количество")
	f.SetCellValue(sheet, "E1", "Итог:")

	fmt.Printf(">>arr>>%+v", sales.SalesArr)

	var lastRow int
	for i, s := range sales.SalesArr {
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", i+2), s.ProductName)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", i+2), s.ProductCode)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "C", i+2), s.Total)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "D", i+2), s.QntTotal)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "E", i+2), s.QntTotal*s.Total)
		lastRow = i
	}

	lastRow += 3

	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
	//f.SetCellValue(sheet, fmt.Sprintf("%s%d", "F", lastRow-1), "Сумма / Процент РБ")
	//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "F", lastRow-1), fmt.Sprintf("%s%d", "F", lastRow-1), style)
	//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "F", lastRow), fmt.Sprintf("%s%d", "F", lastRow), style)
	//f.SetCellValue(sheet, fmt.Sprintf("%s%d", "F", lastRow), discount)
	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "E", lastRow), totalAmount)
	//_ = f.MergeCell(sheet, fmt.Sprintf("%s%d", "A", lastRow), fmt.Sprintf("%s%d", "B", lastRow))
	err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
	err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", 1), fmt.Sprintf("%s%d", "E", 1), style)
	err = f.SetCellStyle(sheet, "A1", "D1", style)
	//f.SetCellValue("Sheet1", "D102", discount)
	//RB1
	if isRB1 {
		f.NewSheet(sheet2)
		f.SetCellValue(sheet2, "A1", "Период")
		f.SetCellValue(sheet2, "B1", "Номер договора/ДС")
		f.SetCellValue(sheet2, "C1", "Тип скидки")
		f.SetCellValue(sheet2, "D1", "Сумма вознаграждения")
		f.SetCellValue(sheet2, "E1", "Сумма скидки")
		err = f.SetCellStyle(sheet2, "A1", "E1", style)

		var totalDiscountsSum int
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range contracts {
			for _, discountStruct := range contract.Discounts {
				if discountStruct.Code != "TOTAL_AMOUNT_OF_SELLING" {
					continue
				}

				var rewardASum int
				var totalSum float32
				if len(contract.Discounts) > 0 {
					if len(contract.Discounts[0].Periods) > 0 {
						rewardASum = contract.Discounts[0].Periods[0].RewardAmount
						totalSum = contract.Discounts[0].Periods[0].TotalAmount
					}
				}
				if totalSum <= totalAmount {
					discount = rewardASum
				}
				f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.ContractParameters.StartDate, contract.ContractParameters.EndDate))
				f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "B", i+2), contract.ContractParameters.ContractNumber)
				f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "C", i+2), "Скидка за объем закупа")
				f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "D", i+2), rewardASum)
				f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "E", i+2), discount)
				totalDiscountsSum += discount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(sheet2, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(sheet2, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	}

	if isRB2 {
		rbSecondType, err := GetAllRBSecondTypeMock(request)
		if err != nil {
			return err
		}
		f.NewSheet(sheet3)
		f.SetCellValue(sheet3, "A1", "Период")
		f.SetCellValue(sheet3, "B1", "Номер договора/ДС")
		f.SetCellValue(sheet3, "C1", "Тип скидки")
		f.SetCellValue(sheet3, "D1", "Бренд")
		f.SetCellValue(sheet3, "E1", "Скидка %")
		f.SetCellValue(sheet3, "F1", "Сумма скидки")
		err = f.SetCellStyle(sheet3, "A1", "F1", style)

		var totalDiscountsSum int
		for i, contract := range rbSecondType {
			f.SetCellValue(sheet3, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(sheet3, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(sheet3, fmt.Sprintf("%s%d", "C", i+2), sheet3)
			f.SetCellValue(sheet3, fmt.Sprintf("%s%d", "D", i+2), contract.BrandName)
			f.SetCellValue(sheet3, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountPercent)
			f.SetCellValue(sheet3, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountAmount)
			totalDiscountsSum += int(contract.DiscountAmount)
			lastRow = i + 2
		}
		lastRow += 1
		f.SetCellValue(sheet3, fmt.Sprintf("%s%d", "E", lastRow), "Итог:")
		f.SetCellValue(sheet3, fmt.Sprintf("%s%d", "F", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(sheet3, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(sheet3, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	}

	if isRB3 {
		rbThirdType, err := GetRBThirdType(request)
		if err != nil {
			return err
		}

		f.NewSheet(sheet4)
		f.SetCellValue(sheet4, "A1", "Период")
		f.SetCellValue(sheet4, "B1", "Номер договора/ДС")
		f.SetCellValue(sheet4, "C1", "Тип скидки")
		f.SetCellValue(sheet4, "D1", "Код товара")
		f.SetCellValue(sheet4, "E1", "План закупа")
		f.SetCellValue(sheet4, "F1", "Скидка %")
		f.SetCellValue(sheet4, "G1", "Сумма скидки")
		err = f.SetCellStyle(sheet4, "A1", "G1", style)

		var totalDiscountsSum float32
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rbThirdType {
			f.SetCellValue(sheet4, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(sheet4, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(sheet4, fmt.Sprintf("%s%d", "C", i+2), sheet4)
			f.SetCellValue(sheet4, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			f.SetCellValue(sheet4, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			f.SetCellValue(sheet4, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountPercent)
			f.SetCellValue(sheet4, fmt.Sprintf("%s%d", "G", i+2), contract.DiscountAmount)
			totalDiscountsSum += contract.DiscountAmount
			lastRow = i + 2
			i++
		}
		lastRow += 1
		f.SetCellValue(sheet4, fmt.Sprintf("%s%d", "F", lastRow), "Итог:")
		f.SetCellValue(sheet4, fmt.Sprintf("%s%d", "G", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(sheet4, fmt.Sprintf("%s%d", "F", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
		//err = f.SetCellStyle(sheet3, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
	}

	if isRB4 {
		rbFourthType := InfoPresentationDiscount(request)
		if err != nil {
			return err
		}

		f.NewSheet(sheet5)
		f.SetCellValue(sheet5, "A1", "Период")
		f.SetCellValue(sheet5, "B1", "Номер договора/ДС")
		f.SetCellValue(sheet5, "C1", "Тип скидки")
		f.SetCellValue(sheet5, "D1", "Скидка %")
		f.SetCellValue(sheet5, "E1", "Сумма скидки")
		//f.SetCellValue(sheet5, "D1", "Код товара")
		//f.SetCellValue(sheet5, "E1", "План закупа")
		err = f.SetCellStyle(sheet5, "A1", "E1", style)

		var totalDiscountsSum float32
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rbFourthType {
			f.SetCellValue(sheet5, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(sheet5, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(sheet5, fmt.Sprintf("%s%d", "C", i+2), sheet5)
			f.SetCellValue(sheet5, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(sheet5, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			//f.SetCellValue(sheet5, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			//f.SetCellValue(sheet5, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			totalDiscountsSum += contract.DiscountAmount
			lastRow = i + 2
			i++
		}
		lastRow += 1
		f.SetCellValue(sheet5, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(sheet5, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(sheet5, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//err = f.SetCellStyle(sheet3, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
	}

	f.SaveAs("files/reports/rb/rb_report.xlsx")
	return nil
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
		hasPresentation bool
		//hasMTZ bool
	)

	for _, contract := range contracts {
		var DoubtedDiscountDetails []model.DoubtedDiscountDetails
		for _, discount := range contract.Discounts {
			var DoubtedDiscountDetail model.DoubtedDiscountDetails
			if discount.Code == "DISCOUNT_FOR_REPRESENTATION" && discount.IsSelected == true && hasPresentation == false {
				DoubtedDiscountDetail.Name = discount.Name
				DoubtedDiscountDetail.Code = discount.Code
				DoubtedDiscountDetail.IsCompleted = true

				DoubtedDiscountDetails = append(DoubtedDiscountDetails, DoubtedDiscountDetail)
				hasPresentation = true
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

func SaveDoubtedDiscountsResults(request model.DoubtedDiscountResponse) error {

	return nil
}
