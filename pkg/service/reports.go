package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
)

type rbDtoFroStoredRbReport struct {
	Period          string  `json:"period"`
	ContractNumber  string  `json:"contract_number"`
	DiscountType    string  `json:"discount_type"`
	BrandName       string  `json:"brand_name,omitempty"`
	ProductCode     string  `json:"product_code,omitempty"`
	LeasePlan       string  `json:"lease_plan"`
	RewardAmount    string  `json:"reward_amount"`
	DiscountPercent string  `json:"discount_percent"`
	DiscountAmount  float32 `json:"discount_amount"`
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetRbDtoTotalAmount(RbDTOs []models.RbDTO, contractID int) (totalAmount float32) {
	for _, RbDTO := range RbDTOs {
		if contractID == RbDTO.ID {
			totalAmount += RbDTO.DiscountAmount
		}
	}

	return totalAmount
}

func StoreRbReports(rbDTOs []models.RbDTO) error {
	var (
		checkedIDs []int
		localRbDTO []models.RbDTO
	)

	for i := 0; i < len(rbDTOs); i++ {
		if contains(checkedIDs, rbDTOs[i].ID) || rbDTOs[i].ID == 0 {
			continue
		}

		var totalDiscountAmount float32

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

		if err = repository.AddOrUpdateReport(report); err != nil {
			return err
		}
	}

	return nil
}

func GetAllStoredReports() (reports []models.StoredReport, err error) {
	reports, err = repository.GetAllStoredReports()
	if err != nil {
		return nil, err
	}

	for i := range reports {
		reports[i].ContractDate = fmt.Sprintf("%s-%s", reports[i].StartDate, reports[i].EndDate)
	}

	return reports, nil
}

func GetStoredReportDetails(storedReportID int) (rbDTOs []models.RbDTO, err error) {
	storedReport, err := repository.GetStoredReportDetails(storedReportID)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(storedReport.Content, &rbDTOs); err != nil {
		return nil, err
	}

	return
}

func GetExcelForStoredExcelReport(storedReportID int) error {
	filePath := "files/reports/rb/rb_stored_reports.xlsx"
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
		totalDiscountsSum float32
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

func convertRbDtoToRbDtoFroStoredRbReport(dto models.RbDTO) (storedRbReport rbDtoFroStoredRbReport) {
	emptyPlaceholder := "		-		"
	storedRbReport.Period = fmt.Sprintf("%s-%s", dto.StartDate, dto.EndDate)
	storedRbReport.ContractNumber = dto.ContractNumber
	storedRbReport.DiscountType = dto.DiscountType
	if dto.BrandName == "" {
		storedRbReport.BrandName = emptyPlaceholder
	} else {
		storedRbReport.BrandName = dto.BrandName
	}

	if dto.ProductCode == "" {
		storedRbReport.ProductCode = emptyPlaceholder
	} else {
		storedRbReport.ProductCode = dto.ProductCode
	}

	if dto.LeasePlan == 0 {
		storedRbReport.LeasePlan = emptyPlaceholder
	} else {
		storedRbReport.LeasePlan = fmt.Sprintf("%2f", dto.LeasePlan)
	}

	if dto.RewardAmount == 0 && dto.DiscountPercent > 0 {
		storedRbReport.RewardAmount = emptyPlaceholder
	} else {
		storedRbReport.RewardAmount = fmt.Sprintf("%2f", dto.RewardAmount)
	}

	if dto.DiscountPercent == 0 && dto.RewardAmount > 0 {
		storedRbReport.DiscountPercent = emptyPlaceholder
	} else {
		storedRbReport.DiscountPercent = fmt.Sprintf("%2f", dto.DiscountPercent)
	}

	storedRbReport.DiscountAmount = dto.DiscountAmount

	return storedRbReport
}
