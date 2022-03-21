package service

import (
	"admin_panel/models"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
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
