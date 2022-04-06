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
			ContractAmountWithDiscount: contract.ContractParameters.ContractAmount - totalDiscountAmount,
		}

		if contract.AdditionalAgreementNumber != 0 {
			report.ContractNumber += fmt.Sprintf(" - ДС №%d", contract.AdditionalAgreementNumber)
		}

		if err = repository.AddOrUpdateReport(report); err != nil {
			return err
		}
	}

	return nil
}

func GetAllStoredReports() (reports []models.StoredReport, err error) {
	return repository.GetAllStoredReports()
}
