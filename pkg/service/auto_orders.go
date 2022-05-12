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

func GetMatrix(code string) ([]models.Matrix, error) {
	matrix := struct {
		MatrixArr []models.Matrix `json:"matrix_arr"`
	}{}

	client := &http.Client{}

	body, err := json.Marshal(struct {
		StoreCode string `json:"store_code"`
	}{
		StoreCode: code,
	})

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

	//log.Println("[ioutil.ReadAll(resp.Body)] = ", string(responseSTR))

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
	/*
	   	responseSTR = []byte(`{
	       "matrix_arr": [
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex тампоны Active Normal № 8 шт ",
	               "product_code": "00000064187",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex тампоны Active Super № 8 шт ",
	               "product_code": "00000064188",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex тампоны Active Normal 24* № 16 шт ",
	               "product_code": "00000065247",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex тампоны Active Super № 16 шт ",
	               "product_code": "00000065248",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики 6 (15-22 кг) для мальчиков № 44 шт ",
	               "product_code": "00000065727",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики 6 (15-22 кг) для детей 6-10 кг № 44 шт ",
	               "product_code": "00000065728",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Elite Soft 3 (6-11 кг) № 25 шт ",
	               "product_code": "00000067262",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Elite Soft 4 (9-14кг) № 21 шт ",
	               "product_code": "00000067263",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Elite Soft 4 (9-14кг) №42 ",
	               "product_code": "00000067264",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Elite Soft 5 (12-17 кг) № 19 шт ",
	               "product_code": "00000067265",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Elite Soft 5 (12-17 кг) №38 ",
	               "product_code": "00000067266",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Elite soft 3 (6-11кг)  №54 ",
	               "product_code": "00000068031",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Elite Soft 1 (3-5 кг)  50 шт ",
	               "product_code": "00000077746",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Elite soft 3 (6-11кг)  №23 ночные ",
	               "product_code": "00000079918",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Elite soft 4 (9-14кг)  №19 ночные ",
	               "product_code": "00000079919",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Elite soft 5 (12-17кг)  №17 ночные ",
	               "product_code": "00000079920",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Elite soft 6 (15-25кг)  №16 ночные ",
	               "product_code": "00000079921",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kleenex салфетки влажные антибактериальные 80 № 10 шт ",
	               "product_code": "00000079962",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Normal ежедневные № 56 шт ",
	               "product_code": "00000080181",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки ультратонкие ежедневные № 56 шт ",
	               "product_code": "00000080183",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Deo Normal ежедневные № 56 шт ",
	               "product_code": "00000080185",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Deo Super тонкие ежедневные № 56 шт ",
	               "product_code": "00000080186",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Super тонкие ежедневные № 20 шт ",
	               "product_code": "00000080285",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Normal ежедневные № 20 шт ",
	               "product_code": "00000080286",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "KIeenex салфетки Box Balzam № 72 шт ",
	               "product_code": "00000074126",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "KIeenex салфетки Box Family № 150 шт ",
	               "product_code": "00000074127",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Active экстратонкие ежедневные № 16 шт ",
	               "product_code": "00000074128",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Active экстратонкие ежедневные № 48 шт ",
	               "product_code": "00000074129",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Active Liners Deo ежедневные № 16 шт ",
	               "product_code": "00000074130",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Active нормал плюс гигиенические № 8 шт ",
	               "product_code": "00000074132",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Active супер плюс гигиенические № 7 шт ",
	               "product_code": "00000074133",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Elite Soft Pants 6 (15-25 кг) Mega № 32 шт ",
	               "product_code": "00000074220",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Elite Soft 0+ (до 3.5 кг) № 25 шт ",
	               "product_code": "00000076278",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Elite Soft 1 (3-5 кг) № 25 шт ",
	               "product_code": "00000076291",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Elite Soft 2 (4-6 кг) № 25 шт ",
	               "product_code": "00000076292",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Elite Soft 2 (4-6 кг)  82 шт ",
	               "product_code": "00000076293",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Elite Soft 2 № 50 шт ",
	               "product_code": "00000076296",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra ночные гигиенические № 24 шт ",
	               "product_code": "00000076307",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex женские прокладки natural нормал  гигиенические № 8 шт ",
	               "product_code": "00000115150",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex женские  прокладки natural супер  гигиенические № 7 шт ",
	               "product_code": "00000115151",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex женские прокладки  Natural ночные  гигиенические № 6 шт ",
	               "product_code": "00000115152",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex женские  ежедневные прокладки Natural нормал  гигиенические № 40 шт ",
	               "product_code": "00000115154",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies салфетки Elite Soft Triplo влажные № 56 шт ",
	               "product_code": "00000115155",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex женские ежедневные прокладки Natural нормал  гигиенические № 20 шт ",
	               "product_code": "00000116806",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies салфетки детские Huggies BW Classic Triplo влажные № 168 шт ",
	               "product_code": "00000116807",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies салфетки детские Huggies BW Ultra Comfort Aloe Triplo влажные № 56 шт ",
	               "product_code": "00000116816",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies салфетки Wipes Tier classic   56 шт ",
	               "product_code": "00000117417",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies салфетки BW Ultra Comfort Aloe   56 шт ",
	               "product_code": "00000117418",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies салфетки детские Elite Soft влажные № 56 шт ",
	               "product_code": "00000118152",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики DryNiltes 8-15 лет ночные для мальчиков № 9 шт ",
	               "product_code": "00000030494",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra Soft нормал гигиенические № 10 шт ",
	               "product_code": "00000045951",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Mega 5 (12-17кг) Girl № 48 шт ",
	               "product_code": "00000048987",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Elite Soft 3 (5-9 кг) № 80 шт ",
	               "product_code": "00000049318",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Elite Soft 4 (8-14 кг) № 66 шт ",
	               "product_code": "00000049319",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Elite Soft 1 (3-5 кг)  84 шт ",
	               "product_code": "00000056042",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex тампоны Super № 32 шт ",
	               "product_code": "00000059755",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики DryNiltes 8-15 лет ночные для девочек № 9 шт ",
	               "product_code": "00000060396",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики 4 (9-14 кг) для девочек № 17 шт ",
	               "product_code": "00000043024",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики 5 (12-17кг) для детей 6-10 кг № 15 шт ",
	               "product_code": "00000043029",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики 4 (9-14кг) для мальчиков № 17 шт ",
	               "product_code": "00000043030",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 4+(10-16 кг) для девочек № 68 шт ",
	               "product_code": "00000007422",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 5 (12-22кг) 15*8 для мальчиков ",
	               "product_code": "00000007481",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort Mega 5 (12-22кг) 56*2 для мальчиков ",
	               "product_code": "00000007525",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 4+ (10-16 кг) № 17 шт ",
	               "product_code": "00000007595",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 4 (8-14кг) для мальчиков № 66 шт ",
	               "product_code": "00000007611",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 4 (8-14кг) для девочек № 80 шт ",
	               "product_code": "00000007635",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 4+ (10-16кг) для мальчиков № 68 шт ",
	               "product_code": "00000007657",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 5 (12-22кг) для мальчиков № 64 шт ",
	               "product_code": "00000007661",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 5 (12-22 кг) для девочек № 64 шт ",
	               "product_code": "00000007662",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 4 (8-14кг) для девочек № 19 шт ",
	               "product_code": "00000007694",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 3 (5-9кг) для мальчиков № 94 шт ",
	               "product_code": "00000007720",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 3 (5-9 кг) для мальчиков № 21 шт ",
	               "product_code": "00000007725",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 3 (5-9 кг) для девочек № 80 шт ",
	               "product_code": "00000007765",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 5 (12-22 кг) для девочек № 15 шт ",
	               "product_code": "00000007770",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 4+(10-16кг) для девочек № 60 шт ",
	               "product_code": "00000007774",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 4 (8-14кг) для мальчиков № 19 шт ",
	               "product_code": "00000007818",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 4 (8-14кг) № 66 шт ",
	               "product_code": "00000007862",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 4 (8-14кг) для мальчиков № 80 шт ",
	               "product_code": "00000007863",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 3 (5-9 кг) для девочек № 21 шт ",
	               "product_code": "00000007873",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 3 (5-9кг) для мальчиков № 80 шт ",
	               "product_code": "00000007884",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort Mega 5 (12-22кг) 56*2 для девочек ",
	               "product_code": "00000008061",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 3 (5-9 кг) для девочек № 94 шт ",
	               "product_code": "00000008067",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики 5 (12-17 кг) для мальчиков № 15 шт ",
	               "product_code": "00000008090",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 4+(10-16кг) для мальчиков № 60 шт ",
	               "product_code": "00000008141",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Ultra Comfort 4+ (10-16 кг) для девочек № 17 шт ",
	               "product_code": "00000008144",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex тампоны супер № 8 шт ",
	               "product_code": "00000008353",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra Soft супер гигиенические № 8 шт ",
	               "product_code": "00000008362",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra Soft супер гигиенические № 16 шт ",
	               "product_code": "00000008364",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra нормал гигиенические № 10 шт ",
	               "product_code": "00000008385",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex тампоны super № 16 шт ",
	               "product_code": "00000008401",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra супер гигиенические № 16 шт ",
	               "product_code": "00000008407",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex тампоны Normal № 16 шт ",
	               "product_code": "00000008451",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Young нормал  гигиенические № 10 шт ",
	               "product_code": "00000008510",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Deo Normal ежедневные № 20 шт ",
	               "product_code": "00000008546",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra Normal сетчатые № 20 шт ",
	               "product_code": "00000008560",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки ультратонкие ежедневные № 20 шт ",
	               "product_code": "00000008563",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra ночные гигиенические № 7 шт ",
	               "product_code": "00000008657",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex тампоны мини № 16 шт ",
	               "product_code": "00000008702",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra супер гигиенические № 8 шт ",
	               "product_code": "00000008726",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra Soft нормал гигиенические № 20 шт ",
	               "product_code": "00000008824",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex тампоны нормал № 8 шт ",
	               "product_code": "00000008830",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex тампоны мини № 8 шт ",
	               "product_code": "00000018377",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Deo ультратонкие ежедневные № 20 шт ",
	               "product_code": "00000043343",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra нормал гигиенические № 40 шт ",
	               "product_code": "00000048110",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra супер гигиенические № 32 шт ",
	               "product_code": "00000048111",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Kotex прокладки Ultra ночные гигиенические № 14 шт ",
	               "product_code": "00000047752",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Elite Soft 4 (8-14 кг) № 19 шт ",
	               "product_code": "00000048981",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Elite Soft 3 (5-9 кг) № 21 шт ",
	               "product_code": "00000048982",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies подгузники Elite Soft 5 (12-22 кг) № 56 шт ",
	               "product_code": "00000048983",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           },
	           {
	               "store_name": "Аптека № 2, Шымкент, (Городской Акимат)",
	               "store_code": "A0000120 ",
	               "region_name": "",
	               "region_code": "",
	               "product_name": "Huggies трусики Mega 5 (12-17кг) Boy № 48 шт ",
	               "product_code": "00000048986",
	               "supplier_name": "Фиркан ТОО",
	               "supplier_code": "000001770",
	               "format": "small2",
	               "min": "0",
	               "max": "0",
	               "import": "false",
	               "defect": "false"
	           }
	       ]
	   }`)
	*/
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

func CreateNecessity() {
	log.Println("|#######################################|", "##########|", "###############################################|", "################|", "###########|")
	log.Println("|Наименование аптеки                    |", "код аптеки|", "Наименование номенклатуры                      |", "Код номенклатуры|", "Потребность|")
	log.Println("|#######################################|", "##########|", "###############################################|", "################|", "###########|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Kotex тампоны Active Normal № 8 шт             |", "00000064187     |", "100        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Kotex тампоны Active Normal 24* № 16 шт        |", "00000065247     |", "120        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Kotex тампоны Active Super № 16 шт             |", "00000065248     |", "175        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Huggies трусики Elite Soft 3 (6-11 кг) № 25 шт |", "00000067262     |", "155        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Huggies трусики Elite Soft 4 (9-14кг) № 21 шт  |", "00000067263     |", "190        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Huggies трусики Elite Soft 4 (9-14кг) №42      |", "00000067264     |", "130        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Huggies трусики Elite Soft 5 (12-17 кг) № 19 шт|", "00000067265     |", "140        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Huggies трусики Elite Soft 5 (12-17 кг) №38    |", "00000067266     |", "200        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")

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
