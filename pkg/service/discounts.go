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
	sheet2 = "РБ №1"
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

	brandInfo := []model.BrandInfo{}
	sales, err := GetSalesBrand(req, brandInfo)

	fmt.Printf("###%+v\n", contracts)
	totalAmount := GetTotalAmount(sales)

	contractRB := DefiningRBReport(contracts, totalAmount)

	return contractRB, nil
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

	totalAmount := GetTotalAmount(sales)

	fmt.Println(contracts)
	fmt.Println(totalAmount)
	var conTotalAmount int
	var rewardAmount int
	if len(contracts) > 0 {
		if len(contracts[0].Discounts) > 0 {
			if len(contracts[0].Discounts[0].Periods) > 0 {
				conTotalAmount = contracts[0].Discounts[0].Periods[0].TotalAmount
				rewardAmount = contracts[0].Discounts[0].Periods[0].RewardAmount
			}
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

	fmt.Printf(">>arr>>%+v", sales.SalesArr)

	var lastRow int
	for i, s := range sales.SalesArr {
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", i+2), s.ProductName)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", i+2), s.ProductCode)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "C", i+2), s.Total)
		lastRow = i
	}

	lastRow += 3

	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", lastRow), "Итог:")
	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "D", lastRow-1), "Сумма / Процент РБ")
	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "D", lastRow), discount)
	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "C", lastRow), totalAmount)
	_ = f.MergeCell(sheet, fmt.Sprintf("%s%d", "A", lastRow), fmt.Sprintf("%s%d", "B", lastRow))
	err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	err = f.SetCellStyle(sheet, "A1", "D1", style)
	//f.SetCellValue("Sheet1", "D102", discount)

	f.NewSheet(sheet2)
	f.SetCellValue(sheet2, "A1", "Период")
	f.SetCellValue(sheet2, "B1", "Номер договора/ДС")
	f.SetCellValue(sheet2, "C1", "Тип скидки")
	f.SetCellValue(sheet2, "D1", "Сумма вознаграждения")
	f.SetCellValue(sheet2, "E1", "Сумма скидки")
	err = f.SetCellStyle(sheet2, "A1", "E1", style)

	var totalDiscountsSum int
	for i, contract := range contracts {
		f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.ContractParameters.StartDate, contract.ContractParameters.EndDate))
		f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "B", i+2), contract.ContractParameters.ContractNumber)
		f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "C", i+2), "Скидка за объем закупа")
		var rewardASum int
		var totalSum int
		if len(contract.Discounts) > 0 {
			if len(contract.Discounts[0].Periods) > 0 {
				rewardASum = contract.Discounts[0].Periods[0].RewardAmount
				totalSum = contract.Discounts[0].Periods[0].TotalAmount
			}
		}
		if totalSum <= totalAmount {
			discount = rewardASum
		}

		f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "D", i+2), rewardASum)
		f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "E", i+2), discount)
		totalDiscountsSum += discount
		lastRow = i + 2
	}
	lastRow += 1
	f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
	f.SetCellValue(sheet2, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)

	f.SaveAs("files/reports/rb/rb_report.xlsx")
	return nil
}

func DefiningRBReport(contracts []model.Contract, totalAmount int) (contractsRB []model.RbDTO) {
	for _, contract := range contracts {
		var contractRB []model.RbDTO
		if len(contract.Discounts) > 0 {
			contractRB = DiscountToReportRB(contract.Discounts[0], contract, totalAmount)
		}

		contractsRB = append(contractsRB, contractRB...)
	}

	return contractsRB
}

func DiscountToReportRB(discount model.Discount, contract model.Contract, totalAmount int) (contractsRB []model.RbDTO) {
	for _, period := range discount.Periods {
		if period.TotalAmount >= totalAmount {
			contractRB := model.RbDTO{
				ID:             contract.ID,
				ContractNumber: contract.ContractParameters.ContractNumber,
				StartDate:      period.PeriodFrom,
				EndDate:        period.PeriodTo,
				DiscountAmount: period.RewardAmount,
			}

			contractsRB = append(contractsRB, contractRB)
		}
	}

	return contractsRB
}

func GetTotalAmount(sales model.Sales) int {
	var amount int
	for _, s := range sales.SalesArr {
		amount += s.Total
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
