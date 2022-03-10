package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"fmt"
	"github.com/xuri/excelize/v2"
)

const (
	DD1Name = "Отложенная скидка за своевременную оплату"
	DD2Name = "Отложенная скидка за своевременную оплату в виде Кредит Ноты"
	DD3Name = "Отложенная скидка за своевременную оплату в виде Оплаты на счет"
	DD4Name = "Отложенная скидка за своевременную оплату в виде Товара"
	DD5Name = "Отложенная скидка за выполнение квартального плана в виде кредит ноты"
	DD6Name = "Отложенная скидка за выполнение годового  плана в виде кредит ноты"
)
const (
	DD1Code = "deferred_discount_for_timely_payment"
	DD2Code = "deferred_discount_for_timely_payment_in_credit_note_form"
	DD3Code = "deferred_discount_for_timely_payment_in_pay_to_account_form"
	DD4Code = "deferred_discount_for_timely_payment_in_goods_form"
	DD5Code = "deferred_discount_for_implementation_of_quarterly_plan_in_credit_note_form"
	DD6Code = "deferred_discount_for_implementation_of_annual_plan_in_credit_note_form"
)

func GetAllDeferredDiscounts(request model.RBRequest) (RbDTO []model.RbDTO, err error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	totalAmount := float32(5000_000.0)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if (discount.Code == DD1Code || discount.Code == DD2Code || discount.Code == DD3Code ||
				discount.Code == DD4Code || discount.Code == DD5Code || discount.Code == DD6Code) && discount.IsSelected == true {
				RbDTO = append(RbDTO, model.RbDTO{
					ContractNumber:  contract.ContractParameters.ContractNumber,
					StartDate:       contract.ContractParameters.StartDate,
					EndDate:         contract.ContractParameters.EndDate,
					DiscountPercent: discount.DiscountPercent,
					DiscountAmount:  totalAmount * discount.DiscountPercent / 100,
					DiscountType:    discount.Name,
				})
			}
		}
	}

	return RbDTO, nil
}

func FormExcelForDeferredDiscounts(request model.RBRequest) error {
	contracts, err := GetAllDeferredDiscounts(request)
	if err != nil {
		return err
	}

	var (
		isDD1 bool
		isDD2 bool
		isDD3 bool
		isDD4 bool
		isDD5 bool
		isDD6 bool
	)

	for _, contract := range contracts {
		if contract.DiscountType == DD1Name {
			isDD1 = true
		}
		if contract.DiscountType == DD2Name {
			isDD2 = true
		}
		if contract.DiscountType == DD3Name {
			isDD3 = true
		}
		if contract.DiscountType == DD4Name {
			isDD4 = true
		}
		if contract.DiscountType == DD5Name {
			isDD5 = true
		}
		if contract.DiscountType == DD6Name {
			isDD6 = true
		}
	}

	//totalAmount := float32(5000_000.0)

	f := excelize.NewFile()
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

	if isDD1 {
		f.NewSheet(DD1Name)
		f.SetCellValue(DD1Name, "A1", "Период")
		f.SetCellValue(DD1Name, "B1", "Номер договора/ДС")
		f.SetCellValue(DD1Name, "C1", "Тип скидки")
		f.SetCellValue(DD1Name, "D1", "Скидка %")
		f.SetCellValue(DD1Name, "E1", "Сумма скидки")
		err = f.SetCellStyle(DD1Name, "A1", "E1", style)

		var totalDiscountsSum float32
		var i int
		var lastRow int
		for _, contract := range contracts {
			if contract.DiscountType == DD1Name {
				f.SetCellValue(DD1Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
				f.SetCellValue(DD1Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
				f.SetCellValue(DD1Name, fmt.Sprintf("%s%d", "C", i+2), DD1Name)
				f.SetCellValue(DD1Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
				f.SetCellValue(DD1Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
				totalDiscountsSum += contract.DiscountAmount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(DD1Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(DD1Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(DD1Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(DD1Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	}

	if isDD2 {
		f.NewSheet(DD2Name)
		f.SetCellValue(DD2Name, "A1", "Период")
		f.SetCellValue(DD2Name, "B1", "Номер договора/ДС")
		f.SetCellValue(DD2Name, "C1", "Тип скидки")
		f.SetCellValue(DD2Name, "D1", "Скидка %")
		f.SetCellValue(DD2Name, "E1", "Сумма скидки")
		err = f.SetCellStyle(DD2Name, "A1", "E1", style)

		var totalDiscountsSum float32
		var i int
		var lastRow int
		for _, contract := range contracts {
			if contract.DiscountType == DD2Name {
				f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
				f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
				f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "C", i+2), DD2Name)
				f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
				f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
				totalDiscountsSum += contract.DiscountAmount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(DD2Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(DD2Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	}

	if isDD3 {
		f.NewSheet(DD3Name)
		f.SetCellValue(DD3Name, "A1", "Период")
		f.SetCellValue(DD3Name, "B1", "Номер договора/ДС")
		f.SetCellValue(DD3Name, "C1", "Тип скидки")
		f.SetCellValue(DD3Name, "D1", "Скидка %")
		f.SetCellValue(DD3Name, "E1", "Сумма скидки")
		err = f.SetCellStyle(DD3Name, "A1", "E1", style)

		var totalDiscountsSum float32
		var i int
		var lastRow int
		for _, contract := range contracts {
			if contract.DiscountType == DD3Name {
				f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
				f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
				f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "C", i+2), DD3Name)
				f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
				f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
				totalDiscountsSum += contract.DiscountAmount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(DD3Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(DD3Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	}

	if isDD4 {
		f.NewSheet(DD4Name)
		f.SetCellValue(DD4Name, "A1", "Период")
		f.SetCellValue(DD4Name, "B1", "Номер договора/ДС")
		f.SetCellValue(DD4Name, "C1", "Тип скидки")
		f.SetCellValue(DD4Name, "D1", "Скидка %")
		f.SetCellValue(DD4Name, "E1", "Сумма скидки")
		err = f.SetCellStyle(DD4Name, "A1", "E1", style)

		var totalDiscountsSum float32
		var i int
		var lastRow int
		for _, contract := range contracts {
			if contract.DiscountType == DD4Name {
				f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
				f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
				f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "C", i+2), DD4Name)
				f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
				f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
				totalDiscountsSum += contract.DiscountAmount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(DD4Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(DD4Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	}

	if isDD5 {
		f.NewSheet(DD5Name)
		f.SetCellValue(DD5Name, "A1", "Период")
		f.SetCellValue(DD5Name, "B1", "Номер договора/ДС")
		f.SetCellValue(DD5Name, "C1", "Тип скидки")
		f.SetCellValue(DD5Name, "D1", "Скидка %")
		f.SetCellValue(DD5Name, "E1", "Сумма скидки")
		err = f.SetCellStyle(DD5Name, "A1", "E1", style)

		var totalDiscountsSum float32
		var i int
		var lastRow int
		for _, contract := range contracts {
			if contract.DiscountType == DD5Name {
				f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
				f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
				f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "C", i+2), DD5Name)
				f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
				f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
				totalDiscountsSum += contract.DiscountAmount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(DD5Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(DD5Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	}

	if isDD6 {
		f.NewSheet(DD6Name)
		f.SetCellValue(DD6Name, "A1", "Период")
		f.SetCellValue(DD6Name, "B1", "Номер договора/ДС")
		f.SetCellValue(DD6Name, "C1", "Тип скидки")
		f.SetCellValue(DD6Name, "D1", "Скидка %")
		f.SetCellValue(DD6Name, "E1", "Сумма скидки")
		err = f.SetCellStyle(DD6Name, "A1", "E1", style)

		var totalDiscountsSum float32
		var i int
		var lastRow int
		for _, contract := range contracts {
			if contract.DiscountType == DD6Name {
				f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
				f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
				f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "C", i+2), DD6Name)
				f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
				f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
				totalDiscountsSum += contract.DiscountAmount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(DD6Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(DD6Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	}

	f.SaveAs("files/reports/dd/reportDD.xlsx")
	return nil
}