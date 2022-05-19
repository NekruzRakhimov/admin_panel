package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"admin_panel/utils"
	"encoding/json"
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
	DD1Code = "DEFERRED_DISCOUNT_FOR_TIMELY_PAYMENT"
	DD2Code = "DEFERRED_DISCOUNT_FOR_TIMELY_PAYMENT_IN_CREDIT_NOTE_FORM"
	DD3Code = "DEFERRED_DISCOUNT_FOR_TIMELY_PAYMENT_IN_PAY_TO_ACCOUNT_FORM"
	DD4Code = "DEFERRED_DISCOUNT_FOR_TIMELY_PAYMENT_IN_GOODS_FORM"
	DD5Code = "DEFERRED_DISCOUNT_FOR_IMPLEMENTATION_OF_QUARTERLY_PLAN_IN_CREDIT_NOTE_FORM"
	DD6Code = "DEFERRED_DISCOUNT_FOR_IMPLEMENTATION_OF_ANNUAL_PLAN_IN_CREDIT_NOTE_FORM"
)

func GetAllDeferredDiscounts(request models.RBRequest) (RbDTO []models.RbDTO, err error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.ClientCode, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	dd1st, err := GetDD1st(request, contracts)
	if err != nil {
		return nil, err
	}
	RbDTO = append(RbDTO, dd1st...)

	dd2nd, err := GetDD2nd(request, contracts)
	if err != nil {
		return nil, err
	}
	RbDTO = append(RbDTO, dd2nd...)

	dd3rd, err := GetDD3rd(request, contracts)
	if err != nil {
		return nil, err
	}
	RbDTO = append(RbDTO, dd3rd...)

	dd4th, err := GetDD4th(request, contracts)
	if err != nil {
		return nil, err
	}
	RbDTO = append(RbDTO, dd4th...)

	dd5th, err := GetDD5th(request, contracts)
	if err != nil {
		return nil, err
	}
	RbDTO = append(RbDTO, dd5th...)

	dd6th, err := GetDD6th(request, contracts)
	if err != nil {
		return nil, err
	}
	RbDTO = append(RbDTO, dd6th...)

	return RbDTO, nil
}

func FormExcelForDeferredDiscounts(request models.RBRequest) error {

	req := models.ReqBrand{
		ClientCode:     request.BIN,
		DateStart:      request.PeriodFrom,
		DateEnd:        request.PeriodTo,
		TypeValue:      "purchase_brand_only",
		TypeParameters: nil,
		//Contracts:      contractsCode, // необходимо получить коды контрактов
	}
	tempPurchases, _ := GetPurchaseBrandOnly(req)

	purchases := UnifyPurchaseBrandOnlyResponse(tempPurchases)

	totalAmount := GetPurchaseTotalAmount(purchases)

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

	const sheetAllPurchases = "итог"

	f.NewSheet(sheetAllPurchases)
	f.SetCellValue(sheetAllPurchases, "A1", "Бренд")
	f.SetCellValue(sheetAllPurchases, "B1", "Номер бренда")
	f.SetCellValue(sheetAllPurchases, "C1", "Период")
	f.SetCellValue(sheetAllPurchases, "D1", "Стоимость")
	f.SetCellValue(sheetAllPurchases, "E1", "Количество")
	f.SetCellValue(sheetAllPurchases, "F1", "Итог:")

	var lastRow int

	period := fmt.Sprintf("%s-%s", request.PeriodFrom, request.PeriodTo)
	fmt.Println("<request>: ", period)

	for i, s := range purchases.PurchaseArr {
		f.SetCellValue(sheetAllPurchases, fmt.Sprintf("%s%d", "A", i+2), s.BrandName)
		f.SetCellValue(sheetAllPurchases, fmt.Sprintf("%s%d", "B", i+2), s.BrandCode)
		f.SetCellValue(sheetAllPurchases, fmt.Sprintf("%s%d", "С", i+2), period)
		f.SetCellValue(sheetAllPurchases, fmt.Sprintf("%s%d", "D", i+2), utils.FloatToMoneyFormat(s.Total/s.QntTotal))
		f.SetCellValue(sheetAllPurchases, fmt.Sprintf("%s%d", "E", i+2), utils.FloatToMoneyFormat(s.QntTotal))
		f.SetCellValue(sheetAllPurchases, fmt.Sprintf("%s%d", "F", i+2), utils.FloatToMoneyFormat(s.Total))
		lastRow = i
	}

	lastRow += 3

	f.SetCellValue(sheetAllPurchases, fmt.Sprintf("%s%d", "E", lastRow), "Итог:")
	f.SetCellValue(sheetAllPurchases, fmt.Sprintf("%s%d", "F", lastRow), utils.FloatToMoneyFormat(totalAmount))
	err = f.SetCellStyle(sheetAllPurchases, fmt.Sprintf("%s%d", "A", lastRow), fmt.Sprintf("%s%d", "F", lastRow), style)
	err = f.SetCellStyle(sheetAllPurchases, fmt.Sprintf("%s%d", "A", 1), fmt.Sprintf("%s%d", "F", 1), style)
	err = f.SetCellStyle(sheetAllPurchases, "A1", "D1", style)

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

		var totalDiscountsSum float64
		var i int
		var lastRow int
		for _, contract := range contracts {
			if contract.DiscountType == DD1Name {
				f.SetCellValue(DD1Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
				f.SetCellValue(DD1Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
				f.SetCellValue(DD1Name, fmt.Sprintf("%s%d", "C", i+2), DD1Name)
				f.SetCellValue(DD1Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
				f.SetCellValue(DD1Name, fmt.Sprintf("%s%d", "E", i+2), utils.FloatToMoneyFormat(float64(contract.DiscountAmount)))
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

		var totalDiscountsSum float64
		var i int
		var lastRow int
		for _, contract := range contracts {
			if contract.DiscountType == DD2Name {
				f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
				f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
				f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "C", i+2), DD2Name)
				f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
				f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "E", i+2), utils.FloatToMoneyFormat(float64(contract.DiscountAmount)))
				totalDiscountsSum += contract.DiscountAmount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(DD2Name, fmt.Sprintf("%s%d", "E", lastRow), utils.FloatToMoneyFormat(float64(totalDiscountsSum)))
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

		var totalDiscountsSum float64
		var i int
		var lastRow int
		for _, contract := range contracts {
			if contract.DiscountType == DD3Name {
				f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
				f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
				f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "C", i+2), DD3Name)
				f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
				f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "E", i+2), utils.FloatToMoneyFormat(float64(contract.DiscountAmount)))
				totalDiscountsSum += contract.DiscountAmount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(DD3Name, fmt.Sprintf("%s%d", "E", lastRow), utils.FloatToMoneyFormat(float64(totalDiscountsSum)))
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

		var totalDiscountsSum float64
		var i int
		var lastRow int
		for _, contract := range contracts {
			if contract.DiscountType == DD4Name {
				f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
				f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
				f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "C", i+2), DD4Name)
				f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
				f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "E", i+2), utils.FloatToMoneyFormat(float64(contract.DiscountAmount)))
				totalDiscountsSum += contract.DiscountAmount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(DD4Name, fmt.Sprintf("%s%d", "E", lastRow), utils.FloatToMoneyFormat(float64(totalDiscountsSum)))
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

		var totalDiscountsSum float64
		var i int
		var lastRow int
		for _, contract := range contracts {
			if contract.DiscountType == DD5Name {
				f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
				f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
				f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "C", i+2), DD5Name)
				f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
				f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "E", i+2), utils.FloatToMoneyFormat(float64(contract.DiscountAmount)))
				totalDiscountsSum += contract.DiscountAmount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(DD5Name, fmt.Sprintf("%s%d", "E", lastRow), utils.FloatToMoneyFormat(float64(totalDiscountsSum)))
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

		var totalDiscountsSum float64
		var i int
		var lastRow int
		for _, contract := range contracts {
			if contract.DiscountType == DD6Name {
				f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
				f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
				f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "C", i+2), DD6Name)
				f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
				f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "E", i+2), utils.FloatToMoneyFormat(float64(contract.DiscountAmount)))
				totalDiscountsSum += contract.DiscountAmount
				lastRow = i + 2
				i++
			}
		}
		lastRow += 1
		f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(DD6Name, fmt.Sprintf("%s%d", "E", lastRow), utils.FloatToMoneyFormat(float64(totalDiscountsSum)))
		err = f.SetCellStyle(DD6Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(DD6Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	}

	f.DeleteSheet("Sheet1")
	f.SaveAs("files/reports/dd/reportDD.xlsx")
	return nil
}

func StoreDdReports(rbDTOs []models.RbDTO) error {
	var (
		checkedIDs []int
		localRbDTO []models.RbDTO
	)

	for i := 0; i < len(rbDTOs); i++ {
		if contains(checkedIDs, rbDTOs[i].ID) || rbDTOs[i].ID == 0 {
			continue
		}

		var totalDiscountAmount float64

		for j := i; j < len(rbDTOs); j++ {
			if rbDTOs[i].ID == rbDTOs[j].ID {
				checkedIDs = append(checkedIDs, rbDTOs[i].ID)
				totalDiscountAmount += rbDTOs[j].DiscountAmount
				localRbDTO = append(localRbDTO, rbDTOs[i])
			}
		}

		contract, err := GetContractDetails(rbDTOs[i].ID)
		if err != nil {
			return err
		}

		report := models.StoredReport{
			Bin:                        contract.Requisites.BIN,
			ContractID:                 contract.ID,
			StartDate:                  contract.ContractParameters.StartDate,
			EndDate:                    contract.ContractParameters.EndDate,
			ContractAmount:             contract.ContractParameters.ContractAmount,
			DiscountAmount:             totalDiscountAmount,
			ContractNumber:             contract.ContractParameters.ContractNumber,
			Beneficiary:                contract.Requisites.Beneficiary,
			ContractAmountWithDiscount: contract.ContractParameters.ContractAmount - totalDiscountAmount,
		}

		if contract.AdditionalAgreementNumber != 0 {
			var contractType string
			//ДС №1 к Договору маркетинговых услуг №1111 ИП  “Adal Trade“
			//marketing_services
			//supply
			switch contract.Type {
			case "marketing_services":
				contractType = "маркетинговых услуг"
			case "supply":
				contractType = "поставок"
			}

			report.ContractNumber = fmt.Sprintf("ДС №%d к Договору %s №%s %s",
				contract.AdditionalAgreementNumber, contractType,
				contract.ContractParameters.ContractNumber,
				contract.Requisites.Beneficiary)
		}

		contentJson, err := json.Marshal(localRbDTO)
		if err != nil {
			return err
		}

		report.Content = contentJson

		if err = repository.AddOrUpdateDdReport(report); err != nil {
			return err
		}
	}

	return nil
}

func GetAllDdStoredReports() (reports []models.StoredReport, err error) {
	reports, err = repository.GetAllDdStoredReports()
	if err != nil {
		return nil, err
	}

	for i := range reports {
		reports[i].ContractDate = fmt.Sprintf("%s-%s", reports[i].StartDate, reports[i].EndDate)
	}

	return reports, nil
}

func GetStoredDdReportDetails(storedReportID int) (rbDTOs []models.RbDTO, err error) {
	storedReport, err := repository.GetDdStoredReportDetails(storedReportID)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(storedReport.Content, &rbDTOs); err != nil {
		return nil, err
	}

	return
}

func GetExcelForDdStoredExcelReport(storedReportID int) error {
	filePath := "files/reports/dd/dd_stored_reports.xlsx"
	sheetName := "Итог"

	rbDTOs, err := GetStoredReportDetails(storedReportID)
	if err != nil {
		return err
	}

	f := excelize.NewFile()
	f.NewSheet(sheetName)
	//var discount int
	//if conTotalAmount <= totalAmount {
	//	discount = rewardAmount
	//}

	style, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#F5DEB3"}, Pattern: 1},
	})

	if err != nil {
		fmt.Println(err)
	}

	f.SetCellValue(sheetName, "A1", "Период")
	f.SetCellValue(sheetName, "B1", "Номер договора/ДС")
	f.SetCellValue(sheetName, "C1", "Тип скидки")
	f.SetCellValue(sheetName, "D1", "Бренд")
	f.SetCellValue(sheetName, "E1", "Код товара")
	f.SetCellValue(sheetName, "F1", "План закупа")
	f.SetCellValue(sheetName, "G1", "Сумма вознаграждения")
	f.SetCellValue(sheetName, "H1", "Скидка %")
	f.SetCellValue(sheetName, "I1", "Сумма скидки")
	err = f.SetCellStyle(sheetName, "A1", "I1", style)
	if err != nil {
		return err
	}

	var (
		i                 int
		lastRow           int
		totalDiscountsSum float64
	)

	for _, rbDTO := range rbDTOs {
		storedRbReport := convertRbDtoToRbDtoFroStoredRbReport(rbDTO)

		f.SetCellValue(sheetName, fmt.Sprintf("%s%d", "A", i+2), storedRbReport.Period)
		f.SetCellValue(sheetName, fmt.Sprintf("%s%d", "B", i+2), storedRbReport.ContractNumber)
		f.SetCellValue(sheetName, fmt.Sprintf("%s%d", "C", i+2), storedRbReport.DiscountType)
		f.SetCellValue(sheetName, fmt.Sprintf("%s%d", "D", i+2), storedRbReport.BrandName)
		f.SetCellValue(sheetName, fmt.Sprintf("%s%d", "E", i+2), storedRbReport.ProductCode)
		f.SetCellValue(sheetName, fmt.Sprintf("%s%d", "F", i+2), storedRbReport.LeasePlan)
		f.SetCellValue(sheetName, fmt.Sprintf("%s%d", "G", i+2), storedRbReport.RewardAmount)
		f.SetCellValue(sheetName, fmt.Sprintf("%s%d", "H", i+2), storedRbReport.DiscountPercent)
		f.SetCellValue(sheetName, fmt.Sprintf("%s%d", "I", i+2), storedRbReport.DiscountAmount)
		totalDiscountsSum += storedRbReport.DiscountAmount
		lastRow = i + 2
		i++
	}
	lastRow += 1
	f.SetCellValue(sheetName, fmt.Sprintf("%s%d", "H", lastRow), "Итог:")
	f.SetCellValue(sheetName, fmt.Sprintf("%s%d", "I", lastRow), totalDiscountsSum)
	err = f.SetCellStyle(sheetName, fmt.Sprintf("%s%d", "H", lastRow), fmt.Sprintf("%s%d", "I", lastRow), style)
	//err = f.SetCellStyle(sheetName, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)

	f.DeleteSheet("Sheet1")
	f.SaveAs(filePath)
	return nil
}
