package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetStoreRegions() ([]models.StoreRegion, error) {
	storeRegions := struct {
		StoreRegionArr []models.StoreRegion `json:"store_region_arr"`
	}{}

	client := &http.Client{}
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getStoresRegions"
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

	err = json.Unmarshal(body, &storeRegions)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Println(string(body))
	return storeRegions.StoreRegionArr, nil
}

func GetMatrixExt(code string) ([]models.Matrix, error) {
	matrix := struct {
		MatrixArr []models.Matrix `json:"matrix_arr"`
	}{}

	client := &http.Client{}

	body, err := json.Marshal(struct {
		StoreCode string `json:"store_code"`
	}{
		StoreCode: code,
	})

	fmt.Printf("%v\n", string(body))
	if err != nil {
		log.Println("[repository.AddOperationExternalService]|[json.Marshal(&paymentRequest)] error is ", err.Error())
		return nil, err
	}

	req, err := http.NewRequest("POST", "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getmatrix", bytes.NewBuffer(body))
	if err != nil {
		log.Println("[repository.AddOperationExternalService]|[http.NewRequest] error is ", err.Error())
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth("http_client", "123456")
	//req.Header.Add("Authorization", "Basic "+basicAuth("http_client", "123456"))
	//req.SetBasicAuth("http_client", "123456" )
	//req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("[client.Do] error is ", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	//handle response
	responseSTR, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ioutil.ReadAll(resp.Body)] error is ", err.Error())
		return nil, err
	}

	log.Println("[ioutil.ReadAll(resp.Body)] = ", string(responseSTR))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	responseSTR = bytes.TrimPrefix(responseSTR, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	if resp.StatusCode != 200 {
		log.Printf("[resp.StatusCode = %d] error is %s", resp.StatusCode, string(responseSTR))
		return nil, errors.New(string(responseSTR))
	}
	err = json.Unmarshal(responseSTR, &matrix)
	if err != nil {
		log.Println(">>")
		log.Println(string(responseSTR))
		log.Println(">>")
		log.Println("[json.Unmarshal(responseSTR, &MatrixArr)]", err.Error())
		return nil, err
	}

	return matrix.MatrixArr, nil
}

func CreateGraphic(graphic models.Graphic) error {
	return repository.CreateGraphic(graphic)
}

func GetAllGraphics() (graphics []models.Graphic, err error) {
	return repository.GetAllGraphics()
}

func GetGraphicByID(id int) (graphic models.Graphic, err error) {
	return repository.GetGraphicByID(id)
}

func EditGraphic(graphic models.Graphic) error {
	return repository.EditGraphic(graphic)
}

func DeleteGraphic(id int) error {
	return repository.DeleteGraphic(id)
}
