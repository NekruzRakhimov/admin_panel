package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
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
