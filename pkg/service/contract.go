package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

func GetContractDetails(contractId int) (contract model.Contract, err error) {
	contractWithJsonB, err := repository.GetContractDetails(contractId)
	if err != nil {
		return model.Contract{}, err
	}

	contract, err = ConvertContractFromJsonB(contractWithJsonB)
	if err != nil {
		return model.Contract{}, err
	}

	switch contractWithJsonB.Status {
	case "черновик":
		contract.Status = "DRAFT"
	case "на согласовании":
		contract.Status = "ON_APPROVAL"
	case "в работе":
		contract.Status = "ACTIVE"
	case "заверщённый":
		contract.Status = "EXPIRED"
	case "отменен":
		contract.Status = "CANCELED"
	default:
		contract.Status = "UNKNOWN"
	}

	return contract, nil
}

func AddAdditionalAgreement(contract model.Contract) error {
	var contractWithJson model.ContractWithJsonB

	contractWithJson.PrevContractId = contract.PrevContractId
	contractWithJson.Comment = contract.Comment
	contractWithJson.Manager = contract.Manager
	contractWithJson.KAM = contract.KAM
	//contractWithJson.Status = contract.Status

	prevContractDetails, err := repository.GetContractDetails(contract.PrevContractId)
	if err != nil {
		return err
	}

	if prevContractDetails.Status != "в работе" {
		return errors.New(fmt.Sprintf("статус договора - [%s]. Вы не можете добавить к нему ДС", prevContractDetails.Status))
	}

	switch contract.Status {
	case "DRAFT":
		contractWithJson.Status = "черновик"
	case "ON_APPROVAL":
		contractWithJson.Status = "на согласовании"
	case "ACTIVE":
		contractWithJson.Status = "в работе"
	case "EXPIRED":
		contractWithJson.Status = "заверщённый"
	default:
		contractWithJson.Status = "неизвестный"
	}

	contractWithJson.Type = prevContractDetails.Type

	requisites, err := json.Marshal(contract.Requisites)
	if err != nil {
		return err
	}
	contractWithJson.Requisites = string(requisites)

	supplierCompanyManager, err := json.Marshal(contract.SupplierCompanyManager)
	if err != nil {
		return err
	}
	contractWithJson.SupplierCompanyManager = string(supplierCompanyManager)

	contractParameters, err := json.Marshal(contract.ContractParameters)
	if err != nil {
		return err
	}
	contractWithJson.ContractParameters = string(contractParameters)

	products, err := json.Marshal(contract.Products)
	if err != nil {
		return err
	}
	contractWithJson.Products = string(products)

	discounts, err := json.Marshal(contract.Discounts)
	if err != nil {
		return err
	}
	contractWithJson.Discounts = string(discounts)

	log.Printf(">>>>>%+v", contractWithJson)
	return repository.CreateContract(contractWithJson)
}

func CreateContract(contract model.Contract) (err error) {
	contractsMiniInfo, err := GetAllContracts("")
	if err != nil {
		return err
	}
	fmt.Printf("contractsMiniInfo: %+v\n", contractsMiniInfo)
	fmt.Printf("contract: %+v\n", contract)

	for _, contractMiniInfo := range contractsMiniInfo {
		if contractMiniInfo.ContractNumber == contract.ContractParameters.ContractNumber {
			return errors.New("договор с таким номером уже существует")
		}
	}

	var contractWithJson model.ContractWithJsonB

	contractWithJson.Type = contract.Type
	contractWithJson.Comment = contract.Comment
	contractWithJson.Manager = contract.Manager
	contractWithJson.KAM = contract.KAM
	contractWithJson.WithTemperatureConditions = contract.WithTemperatureConditions
	//contractWithJson.Status = contract.Status
	switch contract.Status {
	case "DRAFT":
		contractWithJson.Status = "черновик"
	case "ON_APPROVAL":
		contractWithJson.Status = "на согласовании"
	case "ACTIVE":
		contractWithJson.Status = "в работе"
	case "EXPIRED":
		contractWithJson.Status = "заверщённый"
	case "CANCELED":
		contractWithJson.Status = "отменен"
	default:
		contractWithJson.Status = "неизвестный"
	}

	requisites, err := json.Marshal(contract.Requisites)
	if err != nil {
		return err
	}
	contractWithJson.Requisites = string(requisites)

	supplierCompanyManager, err := json.Marshal(contract.SupplierCompanyManager)
	if err != nil {
		return err
	}
	contractWithJson.SupplierCompanyManager = string(supplierCompanyManager)

	contractParameters, err := json.Marshal(contract.ContractParameters)
	if err != nil {
		return err
	}
	contractWithJson.ContractParameters = string(contractParameters)

	products, err := json.Marshal(contract.Products)
	if err != nil {
		return err
	}
	contractWithJson.Products = string(products)

	discounts, err := json.Marshal(contract.Discounts)
	if err != nil {
		return err
	}
	contractWithJson.Discounts = string(discounts)

	log.Printf(">>>>>%+v", contractWithJson)
	return repository.CreateContract(contractWithJson)
}

func EditContract(contract model.Contract) error {
	var contractWithJson model.ContractWithJsonB

	contractWithJson.ID = contract.ID
	contractWithJson.Type = contract.Type
	contractWithJson.Comment = contract.Comment
	contractWithJson.Manager = contract.Manager
	contractWithJson.KAM = contract.KAM
	contractWithJson.WithTemperatureConditions = contract.WithTemperatureConditions
	//contractWithJson.Status = contract.Status
	switch contract.Status {
	case "DRAFT":
		contractWithJson.Status = "черновик"
	case "ON_APPROVAL":
		contractWithJson.Status = "на согласовании"
	case "ACTIVE":
		contractWithJson.Status = "в работе"
	case "EXPIRED":
		contractWithJson.Status = "заверщённый"
	case "CANCELED":
		contractWithJson.Status = "отменен"
	default:
		contractWithJson.Status = "неизвестный"
	}

	requisites, err := json.Marshal(contract.Requisites)
	if err != nil {
		return err
	}
	contractWithJson.Requisites = string(requisites)

	supplierCompanyManager, err := json.Marshal(contract.SupplierCompanyManager)
	if err != nil {
		return err
	}
	contractWithJson.SupplierCompanyManager = string(supplierCompanyManager)

	contractParameters, err := json.Marshal(contract.ContractParameters)
	if err != nil {
		return err
	}
	contractWithJson.ContractParameters = string(contractParameters)

	products, err := json.Marshal(contract.Products)
	if err != nil {
		return err
	}
	contractWithJson.Products = string(products)

	discounts, err := json.Marshal(contract.Discounts)
	if err != nil {
		return err
	}
	contractWithJson.Discounts = string(discounts)

	log.Printf(">>>>>%+v", contractWithJson)
	return repository.EditContract(contractWithJson)
}

func GetAllContracts(contractType string) (contractsMiniInfo []model.ContractMiniInfo, err error) {
	contractsWithJson, err := repository.GetAllContracts(contractType)
	if err != nil {
		return nil, err
	}
	//fmt.Printf(">>>>>>>>>>>>>>>>>contractsWithJson%+v\n", contractsWithJson)

	contracts, err := ConvertContractsFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	//fmt.Printf(">>>>>>>>>>>>>>>>>contracts%+v\n", contracts)

	for _, contract := range contracts {
		contractMiniInfo := ConvertContractToContractMiniInfo(contract)
		contractsMiniInfo = append(contractsMiniInfo, contractMiniInfo)
	}
	//fmt.Printf(">>>>>>>>>>>>>>>>>contractsMiniInfo%+v\n", contractsMiniInfo)

	return contractsMiniInfo, nil
}

func ConvertContractToContractMiniInfo(contract model.Contract) (contractMiniInfo model.ContractMiniInfo) {
	if contract.Type == "marketing_services" {
		contractMiniInfo.ContractType = "Договор маркетинговых услуг"
	} else if contract.Type == "supply" {
		contractMiniInfo.ContractType = "Договор поставок"
	}

	switch contract.Status {
	case "черновик":
		contractMiniInfo.Status = "DRAFT"
	case "на согласовании":
		contractMiniInfo.Status = "ON_APPROVAL"
	case "в работе":
		contractMiniInfo.Status = "ACTIVE"
	case "заверщённый":
		contractMiniInfo.Status = "EXPIRED"
	case "отменен":
		contractMiniInfo.Status = "CANCELED"
	default:
		contractMiniInfo.Status = "UNKNOWN"
	}

	contractMiniInfo.ID = contract.ID
	contractMiniInfo.ContractorName = contract.Requisites.ContractorName
	contractMiniInfo.ContractNumber = contract.ContractParameters.ContractNumber
	contractMiniInfo.Amount = contract.ContractParameters.ContractAmount
	contractMiniInfo.Author = contract.Manager
	contractMiniInfo.CreatedAt = contract.CreatedAt
	contractMiniInfo.UpdatedAt = contract.UpdatedAt
	//contractMiniInfo.Status = contract.Status
	contractMiniInfo.Beneficiary = contract.Requisites.Beneficiary

	return contractMiniInfo
}

func ConvertContractsFromJsonB(contractsWithJsonB []model.ContractWithJsonB) (contracts []model.Contract, err error) {
	for _, contractWithJsonB := range contractsWithJsonB {
		contract, err := ConvertContractFromJsonB(contractWithJsonB)
		if err != nil {
			return nil, err
		}
		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func ConvertContractFromJsonB(contractWithJson model.ContractWithJsonB) (contract model.Contract, err error) {
	contract.ID = contractWithJson.ID
	contract.Type = contractWithJson.Type
	contract.Comment = contractWithJson.Comment
	contract.Manager = contractWithJson.Manager
	contract.KAM = contractWithJson.KAM
	contract.Status = contractWithJson.Status
	contract.CreatedAt = contractWithJson.CreatedAt
	contract.UpdatedAt = contractWithJson.UpdatedAt
	contract.WithTemperatureConditions = contractWithJson.WithTemperatureConditions
	contract.PrevContractId = contractWithJson.PrevContractId

	err = json.Unmarshal([]byte(contractWithJson.Requisites), &contract.Requisites)
	if err != nil {
		return model.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.SupplierCompanyManager), &contract.SupplierCompanyManager)
	if err != nil {
		return model.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.ContractParameters), &contract.ContractParameters)
	if err != nil {
		return model.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.Products), &contract.Products)
	if err != nil {
		return model.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.Discounts), &contract.Discounts)
	if err != nil {
		return model.Contract{}, err
	}

	return contract, nil
}

func ConformContract(contractId int, status string) error {
	if err := repository.ConformContract(contractId, status); err != nil {
		return err
	}

	//todo SAVE TO 1c
	contract, err := GetContractDetails(contractId)
	if err != nil {
		return err
	}

	contractFor1C := ConvertContractToContractDTOFor1CStruct(contract)
	parts := strings.Split(contractFor1C.ContractParameters.StartDate, " ")
	if len(parts) > 0 {
		contractFor1C.ContractParameters.StartDate = parts[0]
	}

	parts = strings.Split(contractFor1C.ContractParameters.EndDate, " ")
	if len(parts) > 0 {
		contractFor1C.ContractParameters.EndDate = parts[0]
	}

	parts = strings.Split(contractFor1C.CreatedAt, " ")
	if len(parts) > 0 {
		contractFor1C.CreatedAt = parts[0]
	}

	parts = strings.Split(contractFor1C.UpdatedAt, " ")
	if len(parts) > 0 {
		contractFor1C.UpdatedAt = parts[0]
	}

	_, _ = SaveContract1C(contractFor1C)
	//if err != nil {
	//	return err
	//}

	//err = repository.SaveContractExternalCode(contractId, respContract.ContractCode)
	//if err != nil {
	//	return err
	//}

	return nil

}

func ConvertContractToContractDTOFor1CStruct(contract model.Contract) (contractFor1C model.ContractDTOFor1C) {
	contractFor1C = model.ContractDTOFor1C{
		ID:                     contract.ID,
		Type:                   contract.Type,
		PrevContractId:         contract.PrevContractId,
		Status:                 contract.Status,
		Requisites:             contract.Requisites,
		Manager:                contract.Manager,
		KAM:                    contract.KAM,
		SupplierCompanyManager: contract.SupplierCompanyManager,
		ContractParameters: model.ContractParametersDTOFor1C{
			ContractNumber:            contract.ContractParameters.ContractNumber,
			ContractAmount:            contract.ContractParameters.ContractAmount,
			Currency:                  contract.ContractParameters.Currency,
			Prepayment:                contract.ContractParameters.Prepayment,
			DateOfDelivery:            contract.ContractParameters.DateOfDelivery,
			FrequencyDeferredDiscount: contract.ContractParameters.FrequencyDeferredDiscount,
			DeliveryAddress:           strings.Join(contract.ContractParameters.DeliveryAddress, "; "),
			DeliveryTimeInterval:      contract.ContractParameters.DeliveryTimeInterval,
			ReturnTimeDelivery:        contract.ContractParameters.ReturnTimeDelivery,
			PriceType:                 "оптом",
			StartDate:                 contract.CreatedAt,
			EndDate:                   contract.ContractParameters.ContractDate,
		},
		WithTemperatureConditions: contract.WithTemperatureConditions,
		Products:                  contract.Products,
		Discounts:                 contract.Discounts,
		Comment:                   contract.Comment,
		CreatedAt:                 contract.CreatedAt,
		UpdatedAt:                 contract.UpdatedAt,
	}

	return contractFor1C
}

func CancelContract(contractId int) error {
	_, err := repository.GetContractDetails(contractId)
	if err != nil {
		return err
	}

	//switch contract.Status {
	//case "черновик":
	//	if err := repository.DisActiveContract(contractId); err != nil {
	//		return err
	//	}
	//case "на согласовании":
	//
	//}

	if err := repository.ChangeContractStatus(contractId, "отменен"); err != nil {
		return err
	}

	return nil
}

func GetContractHistory(contractId int) (contractsMiniInfo []model.ContractMiniInfo, err error) {
	var contracts []model.Contract
	contractWithJsonB, err := repository.GetContractDetails(contractId)
	if err != nil {
		return nil, err
	}

	contract, err := ConvertContractFromJsonB(contractWithJsonB)
	if err != nil {
		return nil, err
	}

	log.Printf("contract (outside the loop): %+v\n", contract)
	contracts = append(contracts, contract)

	if contract.PrevContractId != 0 {
		prevContractId := contract.PrevContractId
		for {
			if prevContractId == 0 {
				break
			}

			contractWithJsonBLoc, err := repository.GetContractDetails(prevContractId)
			if err != nil {
				return nil, err
			}

			contractLoc, err := ConvertContractFromJsonB(contractWithJsonBLoc)
			if err != nil {
				return nil, err
			}

			contracts = append(contracts, contractLoc)
			log.Printf("contractLoc (outside the loop): %+v\n", contractLoc)
			prevContractId = contractLoc.PrevContractId
		}
	}

	for _, contract := range contracts {
		contractMiniInfo := ConvertContractToContractMiniInfo(contract)
		contractsMiniInfo = append(contractsMiniInfo, contractMiniInfo)
	}

	return contractsMiniInfo, nil
}

func FinishContract(contractId int) error {
	return repository.FinishContract(contractId)
}

func RevisionContract(contractId int, comment string) error {
	return repository.RevisionContract(contractId, comment)
}

func GetContractStatusChangesHistory(contractId int) (history []model.ContractStatusHistory, err error) {
	return repository.GetContractStatusChangesHistory(contractId)
}

func SearchContractByNumber(contractNumber, status string) ([]model.SearchContract, error) {
	return repository.SearchContractByNumber(contractNumber, status)

}

func SearchContractHistory(field, param string) ([]model.SearchContract, error) {
	return repository.SearchContractHistory(field, param)

}

func ChangeDataContract(date string, id int, extendContract bool) error {
	return repository.ChangeDataContract(date, id, extendContract)

}
