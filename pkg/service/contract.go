package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	respContract, err := SaveContract1C(contract)
	if err != nil {
		return err
	}

	err = repository.SaveContractExternalCode(contractId, respContract.ContractCode)
	if err != nil {
		return err
	}


	return nil

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

func GetContractHistory(contractId int) (contracts []model.Contract, err error) {
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

	return contracts, nil
}

func FinishContract(contractId int) error {
	return repository.FinishContract(contractId)
}

func RevisionContract(contractId int, comment string) error {
	return repository.RevisionContract(contractId, comment)
}

func CounterpartyContract(binClient string) ([]model.Counterparty, error) {
	var binOrganizationAKNIET = "060540001442"
	client := &http.Client{}
	endpoint := fmt.Sprintf("http://188.225.10.191:5555/api/v2/counterparty/%s/%s", binClient, binOrganizationAKNIET)
	r, err := http.NewRequest("GET", endpoint, nil) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InQua3VzYWlub3ZAbWxhZGV4Lmt6IiwidXNlcklkIjoiNWQ2YzlhNGU0MDVjOWU3NmI3NDI4ZTk3IiwiaWF0IjoxNjMwMDM3MzczLCJleHAiOjE2NjE1NzMzNzN9.yXp9zxxOAJeH53vpa_4Ht4MBQDrThgxxYO1pxFK4t4M"
	//TODO: Надо токен в конфиге или переменой окружения хранить
	r.Header.Add("Authorization", bearer)

	res, err := client.Do(r)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	log.Println(res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
	var contractCounterparty []model.Counterparty
	// ----------> часть Unmarshall json ->
	err = json.Unmarshal(body, &contractCounterparty)
	if err != nil {

		return nil, err
	}

	return contractCounterparty, nil
}

func GetContractStatusChangesHistory(contractId int) (history []model.ContractStatusHistory, err error) {
	return repository.GetContractStatusChangesHistory(contractId)
}




func SaveContract1C(contract  model.Contract) (model.RespContract,  error) {
	var respContract1C model.RespContract

	saveContract := new(bytes.Buffer)
	err := json.NewEncoder(saveContract).Encode(contract)
	if err != nil {
		return respContract1C, err
	}
	client := &http.Client{}
	//endpoint := fmt.Sprintf("http://188.225.10.191:5555/api/v2/counterparty/%s/%s", binClient, binOrganizationAKNIET)
	r, err := http.NewRequest("POST", "http://192.168.0.33/AQG_ULAN/hs/integration/create_contract", saveContract) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")



	res, err := client.Do(r)
	if err != nil {
		//log.Fatal(err)
		return  respContract1C,err
	}
	log.Println(res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return respContract1C,  err

	}
	log.Println(string(body))

	// ----------> часть Unmarshall json ->
	err = json.Unmarshal(body, &respContract1C)
	if err != nil {
		return respContract1C, err
	}


	//TODO: необходим статус то что данные успешно сохранились в 1С и

	//TODO: также сделать проверку статус кода
	return respContract1C, nil
}
