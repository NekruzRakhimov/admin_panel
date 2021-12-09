package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"encoding/json"
	"log"
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

	return contract, nil
}

func CreateContract(contract model.Contract) (err error) {
	var contractWithJson model.ContractWithJsonB

	contractWithJson.Type = contract.Type
	contractWithJson.Comment = contract.Comment
	contractWithJson.Manager = contract.Manager
	contractWithJson.KAM = contract.KAM
	contractWithJson.Status = contract.Status

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
	contractWithJson.Status = contract.Status

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

	contractMiniInfo.ID = contract.ID
	contractMiniInfo.ContractorName = contract.Requisites.ContractorName
	contractMiniInfo.ContractNumber = contract.ContractParameters.ContractNumber
	contractMiniInfo.Amount = contract.ContractParameters.ContractAmount
	contractMiniInfo.Author = contract.Manager
	contractMiniInfo.CreatedAt = contract.CreatedAt
	contractMiniInfo.UpdatedAt = contract.UpdatedAt
	contractMiniInfo.Status = contract.Status

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

func GetAllCurrency() (rights []model.Currency, err error) {
	return repository.GetAllCurrency()
}

func CreateMarketingContract(contract model.MarketingServicesContract) error {
	return repository.CreateMarketingContract(contract)
}

//func AddNewRight(right model.Right) error {
//	return repository.AddNewRight(right)
//}
