package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	baseUrl = "http://89.218.153.38:8081/AQG_ULAN/hs/integration/authorization"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth("username1", "password123"))

	return nil
}

func AddOperationExternalService(login, password string) (response []byte, statusCode int, err error) {
	client := &http.Client{
		Timeout:       60 * time.Second,
		CheckRedirect: redirectPolicyFunc,
	}

	body, err := json.Marshal(struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{
		Login:    login,
		Password: password,
	})
	if err != nil {
		log.Println("[repository.AddOperationExternalService]|[json.Marshal(&paymentRequest)] error is ", err.Error())
		return nil, http.StatusInternalServerError, err
	}

	req, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Println("[repository.AddOperationExternalService]|[http.NewRequest] error is ", err.Error())
		return nil, http.StatusInternalServerError, err
	}
	req.Header.Add("Authorization", "Basic "+basicAuth("http_client", "123456"))
	//req.SetBasicAuth("http_client", "123456" )
	//req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("[repository.AddOperationExternalService]|[client.Do] error is ", err.Error())
		return nil, http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	//handle response
	responseSTR, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[repository.AddOperationExternalService]|[ioutil.ReadAll(resp.Body)] error is ", err.Error())
		return nil, resp.StatusCode, err
	}

	log.Println("[repository.AddOperationExternalService]|[ioutil.ReadAll(resp.Body)] = ", string(responseSTR))

	if resp.StatusCode != 200 {
		log.Printf("[repository.AddOperationExternalService]|[resp.StatusCode = %d] error is %s", resp.StatusCode, string(responseSTR))
		return nil, resp.StatusCode, errors.New(string(responseSTR))
	}

	return responseSTR, http.StatusOK, nil
}

func GetLogin(payload io.Reader) (authResponse models.AuthResponse, err error) {
	fmt.Println("body", payload)
	client := &http.Client{
		Timeout:       60 * time.Second,
		CheckRedirect: redirectPolicyFunc,
	}

	if err != nil {
		log.Println("[repository.AddOperationExternalService]|[json.Marshal(&paymentRequest)] error is ", err.Error())
		return authResponse, err
	}

	req, err := http.NewRequest("POST", baseUrl, payload)
	if err != nil {
		log.Println("[repository.AddOperationExternalService]|[http.NewRequest] error is ", err.Error())
		return authResponse, err
	}
	req.Header.Set("Content-Type", "application/json") // This makes it work
	//req.Header.Add("Authorization", "Basic "+basicAuth("http_client", "123456"))
	req.SetBasicAuth("http_client", "123456")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return authResponse, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return authResponse, err
	}
	//	log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", body)

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return authResponse, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		log.Println(err)
		return authResponse, err
	}

	return authResponse, nil
}

func CounterpartyContract(binClient string) ([]models.Counterparty, error) {
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
	var contractCounterparty []models.Counterparty
	// ----------> часть Unmarshall json ->
	err = json.Unmarshal(body, &contractCounterparty)
	if err != nil {

		return nil, err
	}

	return contractCounterparty, nil
}

func SaveContract1C(contract models.ContractDTOFor1C) (models.RespContract, error) {
	fmt.Println("calling service 1C")
	log.Println("calling service 1C")

	var respContract1C models.RespContract
	saveContract := new(bytes.Buffer)
	err := json.NewEncoder(saveContract).Encode(contract)
	if err != nil {
		return respContract1C, err
	}
	client := &http.Client{}
	//endpoint := fmt.Sprintf("http://188.225.10.191:5555/api/v2/counterparty/%s/%s", binClient, binOrganizationAKNIET)
	r, err := http.NewRequest("POST", "http://89.218.153.38:8081/AQG_ULAN/hs/integration/create_contract", saveContract) // URL-encoded payload
	if err != nil {
		//log.Fatal(err)
		log.Println(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.SetBasicAuth("http_client", "123456")

	res, err := client.Do(r)
	if err != nil {
		//log.Fatal(err)
		return respContract1C, err
	}
	log.Println(res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return respContract1C, err

	}
	log.Println("ответ от 1С", string(body))

	// ----------> часть Unmarshall json ->
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	err = json.Unmarshal(body, &respContract1C)
	if err != nil {
		return respContract1C, err
	}

	//TODO: необходим статус то что данные успешно сохранились в 1С и

	//TODO: также сделать проверку статус кода
	return respContract1C, nil
}

func SearchByBinClient(bin models.ClientBin) (models.Client, error) {
	//var binOrganizationAKNIET = "060540001442"
	var binClient models.Client

	bodyBin := new(bytes.Buffer)
	err := json.NewEncoder(bodyBin).Encode(bin)
	if err != nil {
		return binClient, err
	}
	fmt.Println("BODY", bodyBin)
	client := &http.Client{}
	endpoint := fmt.Sprintf("http://89.218.153.38:8081/AQG_ULAN/hs/integration/client_search")
	r, err := http.NewRequest("POST", endpoint, bodyBin) // URL-encoded payload
	if err != nil {
		return binClient, errors.New("пишешь любой текст ошибки")

	}
	r.Header.Add("Content-Type", "application/json")
	// надо логин и пароль добавить в конфиг
	r.SetBasicAuth("http_client", "123456")

	res, err := client.Do(r)
	if err != nil {

		return binClient, err
	}
	log.Println(res.Status, "мы дошли до сюда")
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body), "RESPONSE")

	// ----------> часть Unmarshall json ->
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	err = json.Unmarshal(body, &binClient)
	if err != nil {
		return binClient, err
	}

	return binClient, nil

}

func GetCurrencies() ([]models.ConvertCurrency, error) {
	var CurrencyArr models.CurrencyArr
	var ConvertCurrencySl []models.ConvertCurrency

	client := &http.Client{}
	//	log.Println(reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/currency_list"
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return ConvertCurrencySl, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ConvertCurrencySl, err
	}
	//log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return ConvertCurrencySl, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &CurrencyArr)
	if err != nil {
		log.Println(err)
		return ConvertCurrencySl, err
	}
	for _, value := range CurrencyArr.CurrencyArr {
		convertCur := models.ConvertCurrency{
			CurrencyName: value.CurrencyName,
			CurrencyCode: value.CurrencyCode,
		}
		ConvertCurrencySl = append(ConvertCurrencySl, convertCur)
	}

	return ConvertCurrencySl, nil

}

func GetCountries() (models.Country, error) {
	countries := models.Country{}
	client := &http.Client{}
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/countrylist"
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return countries, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return countries, err
	}

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return countries, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &countries)
	if err != nil {
		log.Println(err)
		return models.Country{}, err
	}

	fmt.Println(string(body))
	return countries, nil

}

func GetPriceType(bin string) ([]models.PriceTypeAndCode, error) {
	var priceType models.RespPriceType
	priceAndCodeMap := map[string]string{}

	var priceAndCodeSl []models.PriceTypeAndCode

	date := models.ReqBrand{
		ClientBin: bin,
	}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&date)
	fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{}
	log.Println(reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/pricetype"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return priceAndCodeSl, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return priceAndCodeSl, err
	}
	log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return priceAndCodeSl, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &priceType)
	if err != nil {
		log.Println(err)
		return priceAndCodeSl, err
	}

	for _, code := range priceType.PricetypeArr {
		priceAndCodeMap[code.PricetypeCode] = code.PricetypeName
	}
	for key, value := range priceAndCodeMap {
		priceAndCode := models.PriceTypeAndCode{
			PricetypeName: value,
			PricetypeCode: key,
		}
		priceAndCodeSl = append(priceAndCodeSl, priceAndCode)

	}

	return priceAndCodeSl, nil

}

func CreatePriceType(payload models.PriceTypeCreate) (models.PriceTypeResponse, error) {
	var responsePriceType models.PriceTypeResponse

	//date := models.ReqBrand{
	//	ClientBin: bin,
	//}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&payload)
	fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{}
	log.Println(reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/create_pricetype"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return responsePriceType, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return responsePriceType, err
	}
	//log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return responsePriceType, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &responsePriceType)
	if err != nil {
		log.Println(err)
		return responsePriceType, err
	}

	return responsePriceType, nil

}

func CheckContractIn1C(bin string) (models.ResponseContractFrom1C, error) {
	var checkContractFrom1C models.ResponseContractFrom1C
	clientBin := models.BinPriceType{
		ClientBin: bin,
	}
	//date := models.ReqBrand{
	//	ClientBin: bin,
	//}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&clientBin)
	fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{}
	log.Println(reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getcontracts"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return checkContractFrom1C, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return checkContractFrom1C, err
	}
	//log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return checkContractFrom1C, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &checkContractFrom1C)
	if err != nil {
		log.Println(err)
		return checkContractFrom1C, err
	}

	if checkContractFrom1C.ContractArr == nil {
		return models.ResponseContractFrom1C{}, errors.New("Договор с таким бином нет")
	}

	return checkContractFrom1C, nil

}

func CheckContractNumber(contractFor1C models.ContractDTOFor1C) (code int, err error) {

	resp1C, err := CheckContractIn1C(contractFor1C.Requisites.BIN)
	if err != nil {
		return 0, err
	}
	for _, contractParam := range resp1C.ContractArr {
		if contractParam.ContractNumber == contractFor1C.ContractParameters.ContractNumber || contractParam.ContractName == contractFor1C.ContractParameters.ContractNumber {
			fmt.Println("TRUE")
			fmt.Println("DATA FROM 1C", contractParam)
			fmt.Println("OUR'RE BD", contractFor1C)
			err = repository.SaveContractExternalCodeByBIN(contractFor1C, contractParam.ContractCode)
			if err != nil {
				return 0, err
			}

			return 200, nil
		}
	}

	return 0, err
}

func GetRegionsFrom1C() (regions []models.Regions, err error) {
	regions1C := struct {
		RegionArr []models.Regions `json:"region_arr"`
	}{}

	//models.Region{}
	client := &http.Client{}
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/regions"
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &regions1C)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Println(string(body))
	return regions1C.RegionArr, nil
}


func GetListSuppliers() ([]models.DataClient, error)   {
	suppliers := struct {
		RegionArr []models.DataClient `json:"client_arr"`
	}{}
	//var respSupplier models.RespSupplier
	//var suppliers []models.DataClient


	//reqBodyBytes := new(bytes.Buffer)
	//json.NewEncoder(reqBodyBytes).Encode(&clientBin)
	//fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{

	}
	//log.Println(reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getsuppliers"
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &suppliers)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//



	return suppliers.RegionArr, nil
}


