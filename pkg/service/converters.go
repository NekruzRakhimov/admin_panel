package service

import (
	"admin_panel/models"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func ConvertContractToContractMiniInfo(contract models.Contract) (contractMiniInfo models.ContractMiniInfo) {
	log.Println(contract.IsExtendContract, "contract.IsExtendContract")
	log.Println(contract.ID, "ID:")
	log.Println(contract.ExtendDate, "extend_date:")
	if contract.Type == "marketing_services" {
		contractMiniInfo.ContractType = "Договор маркетинговых услуг"
	} else if contract.Type == "supply" {
		contractMiniInfo.ContractType = "Договор поставок"
	} else if contract.PrevContractId != 0 {
		contractMiniInfo.ContractType = "ДС"
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
	// здесь не получаю true

	contractMiniInfo.ID = contract.ID
	contractMiniInfo.ContractorName = contract.Requisites.ContractorName
	contractMiniInfo.ContractNumber = contract.ContractParameters.ContractNumber
	contractMiniInfo.Amount = contract.ContractParameters.ContractAmount
	contractMiniInfo.Author = contract.Manager
	contractMiniInfo.CreatedAt = contract.CreatedAt
	contractMiniInfo.UpdatedAt = contract.UpdatedAt
	contractMiniInfo.AdditionalAgreementNumber = contract.AdditionalAgreementNumber
	//contractMiniInfo.Status = contract.Status
	contractMiniInfo.Beneficiary = contract.Requisites.Beneficiary
	contractMiniInfo.IsExtendContract = contract.IsExtendContract
	contractMiniInfo.ExtendDate = contract.ExtendDate
	contractMiniInfo.StartDate = contract.ContractParameters.StartDate
	contractMiniInfo.EndDate = contract.ContractParameters.EndDate
	contractMiniInfo.ContractName = contract.ContractParameters.ContractName
	contractMiniInfo.Bin = contract.Requisites.BIN
	contractMiniInfo.ContractTypeEng = contract.Type
	contractMiniInfo.View = contract.View
	return contractMiniInfo
}

func BulkConvertContractsFromJsonB(contractsWithJsonB []models.ContractWithJsonB) (contracts []models.Contract, err error) {
	for _, contractWithJsonB := range contractsWithJsonB {
		contract, err := ConvertContractFromJsonB(contractWithJsonB)

		fmt.Printf("цикл============================== %+v\n", contract)
		//TODO: done -> поле extendt - is_extent получаю
		if err != nil {
			return nil, err
		}
		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func ConvertContractFromJsonB(contractWithJson models.ContractWithJsonB) (contract models.Contract, err error) {

	//log.Println("ConvertContractFromJsonB=======================", contractWithJson.ID, contractWithJson.IsExtendContract, contractWithJson.ExtendDate)

	contract.ID = contractWithJson.ID
	contract.AdditionalAgreementNumber = contractWithJson.AdditionalAgreementNumber
	contract.Type = contractWithJson.Type
	contract.Comment = contractWithJson.Comment
	contract.Manager = contractWithJson.Manager
	contract.KAM = contractWithJson.KAM
	contract.Status = contractWithJson.Status
	contract.CreatedAt = contractWithJson.CreatedAt
	contract.UpdatedAt = contractWithJson.UpdatedAt
	contract.WithTemperatureConditions = contractWithJson.WithTemperatureConditions
	contract.PrevContractId = contractWithJson.PrevContractId
	contract.IsExtendContract = contractWithJson.IsExtendContract
	contract.ExtendDate = contractWithJson.ExtendDate
	contract.DiscountBrand = contractWithJson.DiscountBrand
	contract.ExtContractCode = contractWithJson.ExtContractCode
	contract.View = contractWithJson.View

	err = json.Unmarshal([]byte(contractWithJson.Requisites), &contract.Requisites)
	if err != nil {
		log.Println("[service][json.Unmarshal([]byte(contractWithJson.Requisites), &contract.Requisites)] error is: ", err.Error())
		return models.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.SupplierCompanyManager), &contract.SupplierCompanyManager)
	if err != nil {
		log.Println("[service][json.Unmarshal([]byte(contractWithJson.SupplierCompanyManager), &contract.SupplierCompanyManager)] error is: ", err.Error())
		return models.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.ContractParameters), &contract.ContractParameters)
	if err != nil {
		log.Println("[service][.Unmarshal([]byte(contractWithJson.ContractParameters), &contract.ContractParameters)] error is: ", err.Error())
		return models.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.Products), &contract.Products)
	if err != nil {
		log.Println("[service][json.Unmarshal([]byte(contractWithJson.Products), &contract.Products)] error is: ", err.Error())
		return models.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.Discounts), &contract.Discounts)
	if err != nil {
		log.Println("[service][json.Unmarshal([]byte(contractWithJson.Discounts), &contract.Discounts)] error is: ", err.Error())
		return models.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.Regions), &contract.Regions)
	if err != nil {
		log.Println("[service][json.Unmarshal([]byte(contractWithJson.Regions), &contract.Regions)] error is: ", err.Error())
		return models.Contract{}, err
	}

	contract.IsExtendContract = contract.ContractParameters.IsExtendContract

	contract.ExtendDate = contract.ContractParameters.ExtendDate
	log.Println("ДАННЫЕ ПО КОНТРАКТУ", contract)
	return contract, nil
}

func ConvertContractToContractDTOFor1CStruct(contract models.Contract) (contractFor1C models.ContractDTOFor1C) {

	contractFor1C = models.ContractDTOFor1C{
		ID:                     contract.ID,
		Type:                   contract.Type,
		PrevContractId:         contract.PrevContractId,
		Status:                 contract.Status,
		Requisites:             contract.Requisites,
		Manager:                contract.Manager,
		KAM:                    contract.KAM,
		SupplierCompanyManager: contract.SupplierCompanyManager,
		ContractParameters: models.ContractParametersDTOFor1C{
			ContractNumber:            contract.ContractParameters.ContractNumber,
			ContractAmount:            contract.ContractParameters.ContractAmount,
			Currency:                  contract.ContractParameters.CurrencyName,
			Prepayment:                contract.ContractParameters.Prepayment,
			DateOfDelivery:            contract.ContractParameters.DateOfDelivery,
			FrequencyDeferredDiscount: contract.ContractParameters.FrequencyDeferredDiscount,
			DeliveryAddress:           strings.Join(contract.ContractParameters.DeliveryAddress, "; "),
			DeliveryTimeInterval:      contract.ContractParameters.DeliveryTimeInterval,
			ReturnTimeDelivery:        contract.ContractParameters.ReturnTimeDelivery,
			// обновил поля
			CurrencyName:  contract.ContractParameters.CurrencyName,
			CurrencyCode:  contract.ContractParameters.CurrencyCode,
			PricetypeName: contract.ContractParameters.PricetypeName,
			PricetypeCode: contract.ContractParameters.PricetypeCode,
			// до сюда
			StartDate: contract.CreatedAt,
			EndDate:   contract.ContractParameters.ContractDate,
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
