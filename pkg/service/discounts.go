package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"fmt"
	"log"
	"strings"
)

func GetAllRBByContractorBIN(request model.RBRequest) ([]model.RbDTO, error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	//TODO: посмотри потом
	testBin := "060840003599"
	sales, err := GetSales("01.01.2022"+TempDateCompleter, "01.01.2022"+TempDateCompleter, testBin)
	if err != nil {
		return nil, err
	}

	fmt.Printf("###%+v\n", contracts)
	totalAmount := GetTotalAmount(sales)

	contractRB := DefiningRBReport(contracts, totalAmount)

	return contractRB, nil
}

func FormExcelForRBReport(request model.RBRequest) error {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN)
	if err != nil {
		return err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return err
	}

	//TODO: посмотри потом
	testBin := "060840003599"
	sales, err := GetSales("01.01.2022"+TempDateCompleter, "01.01.2022"+TempDateCompleter, testBin)
	if err != nil {
		return err
	}

	totalAmount := GetTotalAmount(sales)

	fmt.Println(contracts)
	fmt.Println(totalAmount)
	//var conTotalAmount int
	//var rewardAmount int
	//if len(contracts) > 0 {
	//	if len(contracts[0].Discounts) > 0 {
	//		if len(contracts[0].Discounts[0].Periods) > 0 {
	//			conTotalAmount = contracts[0].Discounts[0].Periods[0].TotalAmount
	//			rewardAmount = contracts[0].Discounts[0].Periods[0].RewardAmount
	//		}
	//	}
	//}

	//f, err := excelize.OpenFile("files/reports/rb/rb_report_template.xlsx")
	//if err != nil {
	//	return err
	//}
	////var discount int
	////if conTotalAmount <= totalAmount {
	////	discount = rewardAmount
	//	//f.SetCellValue("Sheet1", "D102", rewardAmount)
	////}
	//
	////f.SetCellValue("Sheet1", "D102", discount)
	//
	//f.Save()
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
