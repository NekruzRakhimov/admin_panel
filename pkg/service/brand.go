package service

import (
	"admin_panel/model"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	TempDateCompleter = " 0:02:09"
)

func GetBrands() (model.Brand, error) {
	brand := model.Brand{}
	client := &http.Client{}
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/brandlist"
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return brand, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &brand)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	return brand, nil

}

func GetSales(dateStart, DateEnd string, clientBin string) (model.Sales, error) {
	date := model.DateSales{
		//Datestart: "01.01.2022",
		//Dateend:   "01.01.2022",
		Datestart: dateStart,
		Dateend:   DateEnd,
		ClientBin: clientBin,
	}
	sales := model.Sales{}
	//parm := url.Values{}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&date)
	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{}
	log.Println(reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getsales"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return sales, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &sales)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	return sales, nil

}



//service
func AddBrand(brandName string) (model.AddBrand, error) {
	brand := model.AddBrand{BrandName: brandName}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&brand)

	client := &http.Client{}
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/createbrand"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return brand, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}
	log.Println(string(body))
	err = json.Unmarshal(body, &brand)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	return brand, nil

}