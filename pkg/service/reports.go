package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"fmt"
)

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

func StoreReports(rbDTOs []models.RbDTO) error {
	var (
		checkedIDs []int
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
