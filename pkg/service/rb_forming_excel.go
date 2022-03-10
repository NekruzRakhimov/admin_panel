package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"fmt"
	"github.com/xuri/excelize/v2"
)

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

	var (
		isRB1 bool
		isRB2 bool
		isRB3 bool
		isRB4 bool
		isRB5 bool
		//isRB6  bool
		//isRB7  bool
		isRB8 bool
		//isRB9  bool
		isRB10 bool
		//isRB11 bool
		isRB12 bool
		isRB13 bool
	)

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
			if discount.Code == "DISCOUNT_FOR_LEASE_GENERAL" && discount.IsSelected {
				isRB8 = true
			}
			if discount.Code == RB5Code && discount.IsSelected {
				isRB5 = true
			}
			if discount.Code == RB10Code && discount.IsSelected {
				isRB10 = true
			}
			if discount.Code == RB12Code && discount.IsSelected {
				isRB12 = true
			}
			if discount.Code == RB13Code && discount.IsSelected {
				isRB12 = true
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
	//f.SetCellValue(sheet, fmt.Sprintf("%s%d", "F", lastRow), discount)
	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "E", lastRow), totalAmount)
	//_ = f.MergeCell(sheet, fmt.Sprintf("%s%d", "A", lastRow), fmt.Sprintf("%s%d", "B", lastRow))
	err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
	err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", 1), fmt.Sprintf("%s%d", "E", 1), style)
	err = f.SetCellStyle(sheet, "A1", "D1", style)
	//f.SetCellValue("Sheet1", "D102", discount)
	//RB1

	if isRB1 {
		f.NewSheet(RB1Name)
		f.SetCellValue(RB1Name, "A1", "Период")
		f.SetCellValue(RB1Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB1Name, "C1", "Тип скидки")
		f.SetCellValue(RB1Name, "D1", "Сумма вознаграждения")
		f.SetCellValue(RB1Name, "E1", "Сумма скидки")
		err = f.SetCellStyle(RB1Name, "A1", "E1", style)

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
				f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.ContractParameters.StartDate, contract.ContractParameters.EndDate))
				f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractParameters.ContractNumber)
				f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "C", i+2), "Скидка за объем закупа")
				f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "D", i+2), rewardASum)
				f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "E", i+2), discount)
				totalDiscountsSum += discount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB1Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(RB1Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	}

	if isRB2 {
		rbSecondType, err := GetAllRBSecondType(request)
		if err != nil {
			return err
		}
		f.NewSheet(RB2Name)
		f.SetCellValue(RB2Name, "A1", "Период")
		f.SetCellValue(RB2Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB2Name, "C1", "Тип скидки")
		f.SetCellValue(RB2Name, "D1", "Бренд")
		f.SetCellValue(RB2Name, "E1", "Скидка %")
		f.SetCellValue(RB2Name, "F1", "Сумма скидки")
		err = f.SetCellStyle(RB2Name, "A1", "F1", style)

		var totalDiscountsSum int
		for i, contract := range rbSecondType {
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "C", i+2), RB2Name)
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "D", i+2), contract.BrandName)
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountPercent)
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountAmount)
			totalDiscountsSum += int(contract.DiscountAmount)
			lastRow = i + 2
		}
		lastRow += 1
		f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "E", lastRow), "Итог:")
		f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "F", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	}

	if isRB3 {
		rbThirdType, err := GetRBThirdType(request)
		if err != nil {
			return err
		}

		f.NewSheet(RB3Name)
		f.SetCellValue(RB3Name, "A1", "Период")
		f.SetCellValue(RB3Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB3Name, "C1", "Тип скидки")
		f.SetCellValue(RB3Name, "D1", "Код товара")
		f.SetCellValue(RB3Name, "E1", "План закупа")
		f.SetCellValue(RB3Name, "F1", "Скидка %")
		f.SetCellValue(RB3Name, "G1", "Сумма скидки")
		err = f.SetCellStyle(RB3Name, "A1", "G1", style)

		var totalDiscountsSum float32
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rbThirdType {
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "C", i+2), RB3Name)
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountPercent)
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "G", i+2), contract.DiscountAmount)
			totalDiscountsSum += contract.DiscountAmount
			lastRow = i + 2
			i++
		}
		lastRow += 1
		f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "F", lastRow), "Итог:")
		f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "G", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB3Name, fmt.Sprintf("%s%d", "F", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
		//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
	}

	if isRB4 {
		rbFourthType, err := InfoPresentationDiscount(request)
		if err != nil {
			return err
		}

		f.NewSheet(RB4Name)
		f.SetCellValue(RB4Name, "A1", "Период")
		f.SetCellValue(RB4Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB4Name, "C1", "Тип скидки")
		f.SetCellValue(RB4Name, "D1", "Скидка %")
		f.SetCellValue(RB4Name, "E1", "Сумма скидки")
		//f.SetCellValue(RB4Name, "D1", "Код товара")
		//f.SetCellValue(RB4Name, "E1", "План закупа")
		err = f.SetCellStyle(RB4Name, "A1", "E1", style)

		var totalDiscountsSum float32
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rbFourthType {
			f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "C", i+2), RB4Name)
			f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			totalDiscountsSum += contract.DiscountAmount
			lastRow = i + 2
			i++
		}
		lastRow += 1
		f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB4Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
	}

	if isRB5 {
		rb5thType, err := GetRB5thType(request)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
		f.NewSheet(RB5Name)
		f.SetCellValue(RB5Name, "A1", "Период")
		f.SetCellValue(RB5Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB5Name, "C1", "Тип скидки")
		f.SetCellValue(RB5Name, "D1", "Бренд")
		f.SetCellValue(RB5Name, "E1", "Скидка %")
		f.SetCellValue(RB5Name, "F1", "Сумма скидки")
		err = f.SetCellStyle(RB5Name, "A1", "F1", style)

		var totalDiscountsSum int
		for i, contract := range rb5thType {
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "C", i+2), RB5Name)
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "D", i+2), contract.BrandName)
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountPercent)
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountAmount)
			totalDiscountsSum += int(contract.DiscountAmount)
			lastRow = i + 2
		}
		lastRow += 1
		f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "E", lastRow), "Итог:")
		f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "F", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB5Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)

	}

	if isRB8 {
		rbEighthType, err := GetRBEighthType(request)
		if err != nil {
			return err
		}

		f.NewSheet(RB8Name)
		f.SetCellValue(RB8Name, "A1", "Период")
		f.SetCellValue(RB8Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB8Name, "C1", "Тип скидки")
		f.SetCellValue(RB8Name, "D1", "Скидка %")
		f.SetCellValue(RB8Name, "E1", "Сумма скидки")
		//f.SetCellValue(RB4Name, "D1", "Код товара")
		//f.SetCellValue(RB4Name, "E1", "План закупа")
		err = f.SetCellStyle(RB8Name, "A1", "E1", style)

		var totalDiscountsSum float32
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rbEighthType {
			f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "C", i+2), RB8Name)
			f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			totalDiscountsSum += contract.DiscountAmount
			lastRow = i + 2
			i++
		}
		lastRow += 1
		f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB8Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
	}

	if isRB10 {
		rbFourthType, err := InfoPresentationDiscount(request)
		if err != nil {
			return err
		}

		f.NewSheet(RB10Name)
		f.SetCellValue(RB10Name, "A1", "Период")
		f.SetCellValue(RB10Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB10Name, "C1", "Тип скидки")
		f.SetCellValue(RB10Name, "D1", "Скидка %")
		f.SetCellValue(RB10Name, "E1", "Сумма скидки")
		//f.SetCellValue(RB4Name, "D1", "Код товара")
		//f.SetCellValue(RB4Name, "E1", "План закупа")
		err = f.SetCellStyle(RB10Name, "A1", "E1", style)

		var totalDiscountsSum float32
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rbFourthType {
			f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "C", i+2), RB10Name)
			f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			totalDiscountsSum += contract.DiscountAmount
			lastRow = i + 2
			i++
		}
		lastRow += 1
		f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB10Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
	}

	if isRB12 {
		rb12thType, err := RbDiscountForSalesGrowth(request)
		if err != nil {
			return err
		}

		f.NewSheet(RB12Name)
		f.SetCellValue(RB12Name, "A1", "Период")
		f.SetCellValue(RB12Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB12Name, "C1", "Тип скидки")
		f.SetCellValue(RB12Name, "D1", "Скидка %")
		f.SetCellValue(RB12Name, "E1", "Сумма скидки")
		//f.SetCellValue(RB4Name, "D1", "Код товара")
		//f.SetCellValue(RB4Name, "E1", "План закупа")
		err = f.SetCellStyle(RB12Name, "A1", "E1", style)

		var totalDiscountsSum float32
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rb12thType {
			f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "C", i+2), RB12Name)
			f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			totalDiscountsSum += contract.DiscountAmount
			lastRow = i + 2
			i++
		}
		lastRow += 1
		f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB12Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
	}

	if isRB13 {
		rb12thType, err := DiscountRBPeriodTime(request)
		if err != nil {
			return err
		}

		f.NewSheet(RB13Name)
		f.SetCellValue(RB13Name, "A1", "Период")
		f.SetCellValue(RB13Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB13Name, "C1", "Тип скидки")
		f.SetCellValue(RB13Name, "D1", "Скидка %")
		f.SetCellValue(RB13Name, "E1", "Сумма скидки")
		//f.SetCellValue(RB4Name, "D1", "Код товара")
		//f.SetCellValue(RB4Name, "E1", "План закупа")
		err = f.SetCellStyle(RB13Name, "A1", "E1", style)

		var totalDiscountsSum float32
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rb12thType {
			f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "C", i+2), RB13Name)
			f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			totalDiscountsSum += contract.DiscountAmount
			lastRow = i + 2
			i++
		}
		lastRow += 1
		f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB13Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
	}

	f.SaveAs("files/reports/rb/rb_report.xlsx")
	return nil
}