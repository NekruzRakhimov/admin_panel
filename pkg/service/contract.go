package service

import (
	"admin_panel/db"
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

func GetContractDetails(contractId int) (contract models.Contract, err error) {
	contractWithJsonB, err := repository.GetContractDetails(contractId)
	if err != nil {
		return models.Contract{}, err
	}

	contract, err = ConvertContractFromJsonB(contractWithJsonB)
	if err != nil {
		return models.Contract{}, err
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

func AddAdditionalAgreement(contract models.Contract) error {
	var contractWithJson models.ContractWithJsonB

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

	contractWithJson.AdditionalAgreementNumber = prevContractDetails.AdditionalAgreementNumber + 1

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

func CreateContract(contract models.Contract) (err error) {

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

	var contractWithJson models.ContractWithJsonB

	contractWithJson.Type = contract.Type
	contractWithJson.Comment = contract.Comment
	contractWithJson.Manager = contract.Manager
	contractWithJson.KAM = contract.KAM
	contractWithJson.WithTemperatureConditions = contract.WithTemperatureConditions
	//contractWithJson.DiscountBrand = contract.BrandName
	contractWithJson.DiscountBrand = contract.DiscountBrand
	//	contractWithJson.DiscountBrand = contract.DiscountBrand
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

func EditContract(contract models.Contract) error {
	var contractWithJson models.ContractWithJsonB

	contractWithJson.ID = contract.ID
	contractWithJson.Type = contract.Type
	contractWithJson.Comment = contract.Comment
	contractWithJson.Manager = contract.Manager
	contractWithJson.KAM = contract.KAM
	contractWithJson.WithTemperatureConditions = contract.WithTemperatureConditions
	contractWithJson.DiscountBrand = contract.DiscountBrand
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

func GetAllContracts(contractType string) (contractsMiniInfo []models.ContractMiniInfo, err error) {
	// здесь ты получаешь все поля
	contractsWithJson, err := repository.GetAllContracts(contractType)
	if err != nil {
		return nil, err
	}
	//fmt.Printf(">>>>>>>>>>>>>>>>>contractsWithJson%+v\n", contractsWithJson)

	log.Println(contractsWithJson, "ПОСМОТРИ РЕЗУЛЬТАТ")
	// до этого момента я получаю нужный результат
	fmt.Printf("my_logs[ %+v\n]", contractsWithJson)

	// TODO: проблема либо тут
	contracts, err := ConvertContractsFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	//fmt.Printf(">>>>>>>>>>>>>>>>>contracts%+v\n", contracts)

	for _, contract := range contracts {
		//fmt.Printf(">>>>>>>>>>>>>>>>>loop contract>>>>>>>%+v\n", contract)
		//fmt.Printf("отдельные данные:  %t", contract.IsExtendContract)
		fmt.Printf("BOOOL %t\n", contract.IsExtendContract)
		fmt.Printf("%v\n", contract.IsExtendContract)
		// проблема тут
		contractMiniInfo := ConvertContractToContractMiniInfo(contract)
		contractsMiniInfo = append(contractsMiniInfo, contractMiniInfo)
	}
	//fmt.Printf(">>>>>>>>>>>>>>>>>contractsMiniInfo%+v\n", contractsMiniInfo)

	return contractsMiniInfo, nil
}

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

	return contractMiniInfo
}

func ConvertContractsFromJsonB(contractsWithJsonB []models.ContractWithJsonB) (contracts []models.Contract, err error) {
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

	log.Println("ConvertContractFromJsonB=======================", contractWithJson.ID, contractWithJson.IsExtendContract, contractWithJson.ExtendDate)

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

	contract.IsExtendContract = contract.ContractParameters.IsExtendContract

	contract.ExtendDate = contract.ContractParameters.ExtendDate
	log.Println("ДАННЫЕ ПО КОНТРАКТУ", contract)
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

	//code, err := CheckContractNumber(contractFor1C)
	//if err != nil {
	//	return err
	//}
	//if code != 200{
	//
	//respFrom1C, err := SaveContract1C(contractFor1C)
	//if err != nil {
	//	return err
	//}
	//
	//if respFrom1C.Status != "success" {
	//	return errors.New("не удалось сохранить договор в 1С. Повторите попытку позже")
	//}
	//
	//if err = repository.SaveContractExternalCode(contractId, respFrom1C.ContractCode); err != nil {
	//	return err
	//}
	//}

	return nil

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

func GetContractHistory(contractId int) (contractsMiniInfo []models.ContractMiniInfo, err error) {
	var contracts []models.Contract
	contractWithJsonB, err := repository.GetContractDetails(contractId)
	if err != nil {
		return nil, err
	}

	contract, err := ConvertContractFromJsonB(contractWithJsonB)
	if err != nil {
		return nil, err
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

		contract.ContractParameters.ContractNumber = fmt.Sprintf("ДС №%d к Договору %s №%s %s",
			contract.AdditionalAgreementNumber, contractType,
			contract.ContractParameters.ContractNumber,
			contract.Requisites.Beneficiary)
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

	var contractsMiniInfoWithoutDrafts []models.ContractMiniInfo

	for _, info := range contractsMiniInfo {
		if info.Status != "DRAFT" {
			contractsMiniInfoWithoutDrafts = append(contractsMiniInfoWithoutDrafts, info)
		}
	}

	for i := 0; i < len(contractsMiniInfoWithoutDrafts)-1; i++ {
		if contractsMiniInfoWithoutDrafts[i].ID > contractsMiniInfoWithoutDrafts[i+1].ID {
			temp := contractsMiniInfoWithoutDrafts[i]
			contractsMiniInfoWithoutDrafts[i] = contractsMiniInfoWithoutDrafts[i+1]
			contractsMiniInfoWithoutDrafts[i+1] = temp
		}
	}

	return contractsMiniInfoWithoutDrafts, nil
}

func FinishContract(contractId int) error {
	return repository.FinishContract(contractId)
}

func RevisionContract(contractId int, comment string) error {
	return repository.RevisionContract(contractId, comment)
}

func GetContractStatusChangesHistory(contractId int) (history []models.ContractStatusHistory, err error) {
	return repository.GetContractStatusChangesHistory(contractId)
}

func SearchContractByNumber(contractNumber, status string) ([]models.SearchContract, error) {
	return repository.SearchContractByNumber(contractNumber, status)

}

func SearchContractHistory(field, param string) ([]models.SearchContract, error) {
	return repository.SearchContractHistory(field, param)

}

func SearchHistoryExecution(field, param string) ([]models.SearchContract, error) {
	return repository.SearchHistoryExecution(field, param)

}

func ChangeDataContract(id int) error {
	return repository.ChangeDataContract(id)

}

func GetExternalCode(bin string) []models.ContractCode {
	var ExtContractCode []models.ContractCode
	db.GetDBConn().Raw("SELECT ext_contract_code FROM contracts WHERE requisites ->> 'bin' =  $1", bin).Scan(&ExtContractCode)

	return ExtContractCode
}
