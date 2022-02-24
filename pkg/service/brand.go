package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	TempDateCompleter = " 0:00:00"
	TempDateEnd       = " 23:59:59"
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

func GetSales(reqBrand model.ReqBrand) (model.Sales, error) {
	var sales model.Sales

	date := model.ReqBrand{
		ClientBin:      reqBrand.ClientBin,
		DateStart:      reqBrand.DateStart + TempDateCompleter,
		DateEnd:        reqBrand.DateEnd + TempDateEnd,
		Type:           "sales",
		TypeValue:      "",
		TypeParameters: nil,
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
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
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
	log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", body)

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

func GetBrandInfo(bin string) ([]model.BrandInfo, error) {
	return repository.GetBrandInfo(bin)

}

func GetSalesBrand(reqBrand model.ReqBrand, brandInfo []model.BrandInfo) (model.Sales, error) {
	var sales model.Sales

	date := model.ReqBrand{
		ClientBin:      reqBrand.ClientBin,
		DateStart:      reqBrand.DateStart + TempDateCompleter,
		DateEnd:        reqBrand.DateEnd + TempDateEnd,
		Type:           "sales",
		TypeValue:      "",
		TypeParameters: nil,
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
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
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
	log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", body)

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

func GetBrandSales(reqBrand model.ReqBrand) (model.Sales, error) {
	var sales model.Sales

	date := model.ReqBrand{
		ClientBin:      reqBrand.ClientBin,
		DateStart:      reqBrand.DateStart + TempDateCompleter,
		DateEnd:        reqBrand.DateEnd + TempDateEnd,
		Type:           "sales",
		TypeValue:      reqBrand.TypeValue,
		TypeParameters: reqBrand.TypeParameters,
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
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
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
	log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

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

func FoundBrandDiscount(reqBrand model.ReqBrand) []model.RbDTO {
	var rbBrands []model.RbDTO
	var rbBrand model.RbDTO
	var totalBrandsDiscount []model.TotalBrandDiscount
	var BrandTotal model.TotalBrandDiscount

	f, err := excelize.OpenFile("files/reports/rb/rb_report.xlsx")
	if err != nil {
		fmt.Println(err)
		return rbBrands
	}




	// берем бренды по Бину

	// тут возвращаю
	//brand AS brand_name, discount_percent, contract_id - 3 fields
	dataBrand, contractNumber:= repository.GetIDBYBIN(reqBrand.ClientBin)
	log.Println("ДАННЫЕ БРЕНДОВ", dataBrand)


	// тут что ищем?
	// dataBrand - ID, discount, brand


	for _, value := range dataBrand {
		var floatPercent float32
		valuePercent, err := strconv.ParseFloat(value.DiscountPercent, 32)
		if err != nil {
			// do something sensible
		}
		floatPercent = float32(valuePercent)






		// тут добавляю бренды в TypeParameters
		reqBrand.TypeParameters = append(reqBrand.TypeParameters, value.BrandName)
		//ID              int     `json:"id"` -- отдадим сами
		//ContractNumber  string  `json:"contract_number"`  -- взял
		//StartDate       string  `json:"start_date"`  -- возьмем из запроса
		//EndDate         string  `json:"end_date"` -- возьмем из запроса
		//BrandName       string  `json:"brand_name,omitempty"` -- отдадим сами
		//DiscountPercent float32 `json:"discount_percent"` -- отдадим сами
		//DiscountAmount  float32 `json:"discount_amount"` --  отдадим тоже сами
		BrandTotal.ContractNumber = contractNumber
		BrandTotal.BrandName = value.BrandName
		BrandTotal.DiscountPercent = floatPercent
		BrandTotal.Id, _ = strconv.Atoi(value.ContractID)

		totalBrandsDiscount = append(totalBrandsDiscount, BrandTotal)
	}

	log.Println("reqBrand.TypeParameters", reqBrand.TypeParameters)


	// BrandName string  `json:"brand_name"`
	//	Amount    float32 `json:"amount"` - сумма чего тогда


	// reqBrand -> он дает массив брендов

	sales, _ := GetBrandSales(reqBrand)

	// Берет определенные бренды из 1С:
	counter := 1
	for _, sale := range sales.SalesArr {

		//   {
		//            "product_name": "7Stick жевательная резинка Арбуз 14,5 г ",
		//            "product_code": "00000074577",
		//            "total": 3600,
		//            "qnt_total": 36,
		//            "date": "2022-01-03T00:00:00",
		//            "brand_code": "000000137",
		//            "brand_name": "7Stick"
		//        },



		count := float32(0)
		//	тут будет список брендов

		//TODO: тут ты можешь сразу записать в экселе наименование,кол-во,код, сумму, бренд,
		//sale.BrandName

		//TODO: после того как мы записали все бренды, мы должны посчитать от него общую сумму
		for _, brand := range reqBrand.TypeParameters {
			log.Println("BRAND", brand)

			// мы нашли схожие бренды, что мы должны сделать?
			if sale.BrandName == brand {
				count += sale.Total // - это сколько было продано по данному товару
				//TODO: ты итог должен записать и после чего какой процент и только потом сумму скидки

				//TOTAL
				// Percent

				counter += 1
				strCount := fmt.Sprint(counter)

				//log.Println("Наименование: ", sale.ProductName)

				//log.Println("Количество:", sale.QntTotal)
				//log.Println("Номер продукта:", sale.ProductCode)
				//log.Println("Стоимость: ", sale.Total)
				//log.Println("БРЕНД:", sale.BrandName)
				f.SetCellValue("Sheet1", "A"+strCount, sale.ProductName)
				f.SetCellValue("Sheet1", "B"+strCount, sale.QntTotal)
				f.SetCellValue("Sheet1", "C"+strCount, sale.ProductCode)
				f.SetCellValue("Sheet1", "D"+strCount, sale.Total)
				f.SetCellValue("Sheet1", "E"+strCount, sale.BrandName)
				if err := f.SaveAs("reportRB_brand1.xlsx"); err != nil {
					fmt.Println("ERRRRRRRRRRRRRRRRRRRRRRRRRRRRROOOOOOOOOOORRRRRRRRRRR", err)

				}


				// Тут он пустой
				fmt.Println("totalBrandsDiscount", totalBrandsDiscount)
				fmt.Println("LEN:", len(totalBrandsDiscount))
				if len(totalBrandsDiscount) == 0 {

					//TODO: тут записываем каждый бренд и его общую сумму, это и есть TOTAL но по скидкам

					log.Println("сРАБАОТЛО")
					BrandTotal.BrandName = brand
					BrandTotal.Amount = count

					totalBrandsDiscount = append(totalBrandsDiscount, BrandTotal)

				}


				//  в начале так как массив длина его 0 - поэтому мы заранее добавили 1 бренд
				// после чего делаем проверку, если нашли схожий бренд - просто делаем + суммы к этому бренду
				for i, check := range totalBrandsDiscount {
					log.Println("___________________________________________________________________")

					if brand != check.BrandName {
						BrandTotal.BrandName = brand
						BrandTotal.Amount = count
						fmt.Println("записался в массив")
						log.Println("записался в массив")
						totalBrandsDiscount = append(totalBrandsDiscount, BrandTotal)
					} else if brand == check.BrandName {
						fmt.Println("обновили  массив")
						log.Println("обновили  массив")
						totalBrandsDiscount[i].Amount += count
					}

				}

				// когда мы посчитали общую сумму данного бренда мы должны найти его скидку

			}

		}

	}

	for _, value := range totalBrandsDiscount {
		fmt.Println("ИтогСуммы: ", value.Amount)
		TotalPercent := (value.Amount * value.DiscountPercent) / 100
		log.Println("Сумма скидки: ", TotalPercent)
		log.Println("Скидка: ", value.DiscountPercent)
		//ID              int     `json:"id"`
		//ContractNumber  string  `json:"contract_number"`
		//StartDate       string  `json:"start_date"`
		//EndDate         string  `json:"end_date"`
		//BrandName       string  `json:"brand_name,omitempty"`
		//DiscountPercent float32 `json:"discount_percent"`
		//DiscountAmount  float32 `json:"discount_amount"`

		rbBrand.ID = value.Id
		rbBrand.ContractNumber = value.ContractNumber
		rbBrand.StartDate = reqBrand.DateStart
		rbBrand.EndDate = reqBrand.DateEnd
		rbBrand.DiscountPercent = value.DiscountPercent
		rbBrand.DiscountAmount = TotalPercent




	}







	//ID              int     `json:"id"`
	//ContractNumber  string  `json:"contract_number"`
	//StartDate       string  `json:"start_date"`
	//EndDate         string  `json:"end_date"`
	//BrandName       string  `json:"brand_name,omitempty"`
	//DiscountPercent float32 `json:"discount_percent"`
	//DiscountAmount  float32 `json:"discount_amount"`
	//TODO: вернуть даннные:
	// 1. процент
	// 2. Сумму скидки
	// 3. Имя Бренда
	// 4.  номер договора??
	// 5. ID Договора


	log.Println(rbBrands, "ОТВЕТ ТВОЕЙ МОДЕЛИ")

	return rbBrands

}

const (
	rb2Mock = `{
    "sales_arr": [
        {
            "product_name": "7Stick жевательная резинка Арбуз 14,5 г ",
            "product_code": "00000074577",
            "total": 3600,
            "qnt_total": 36,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Арбуз 14,5 г ",
            "product_code": "00000074577",
            "total": 13900,
            "qnt_total": 139,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Арбуз 14,5 г ",
            "product_code": "00000074577",
            "total": 3400,
            "qnt_total": 34,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Арбуз 14,5 г ",
            "product_code": "00000074577",
            "total": 600,
            "qnt_total": 6,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Арбуз 14,5 г ",
            "product_code": "00000074577",
            "total": 4000,
            "qnt_total": 40,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Арбуз 14,5 г ",
            "product_code": "00000074577",
            "total": 1300,
            "qnt_total": 13,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Арбуз 14,5 г ",
            "product_code": "00000074577",
            "total": 1000,
            "qnt_total": 10,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Дыня 14,5 г ",
            "product_code": "00000120965",
            "total": 3200,
            "qnt_total": 32,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Дыня 14,5 г ",
            "product_code": "00000120965",
            "total": 14500,
            "qnt_total": 145,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Дыня 14,5 г ",
            "product_code": "00000120965",
            "total": 4000,
            "qnt_total": 40,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Дыня 14,5 г ",
            "product_code": "00000120965",
            "total": 2300,
            "qnt_total": 23,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Дыня 14,5 г ",
            "product_code": "00000120965",
            "total": 1000,
            "qnt_total": 10,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Дыня 14,5 г ",
            "product_code": "00000120965",
            "total": 900,
            "qnt_total": 9,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Клубника 14,5 г ",
            "product_code": "00000074579",
            "total": 3100,
            "qnt_total": 31,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Клубника 14,5 г ",
            "product_code": "00000074579",
            "total": 14800,
            "qnt_total": 148,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Клубника 14,5 г ",
            "product_code": "00000074579",
            "total": 4000,
            "qnt_total": 40,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Клубника 14,5 г ",
            "product_code": "00000074579",
            "total": 300,
            "qnt_total": 3,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Клубника 14,5 г ",
            "product_code": "00000074579",
            "total": 2600,
            "qnt_total": 26,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Клубника 14,5 г ",
            "product_code": "00000074579",
            "total": 1600,
            "qnt_total": 16,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Клубника 14,5 г ",
            "product_code": "00000074579",
            "total": 1300,
            "qnt_total": 13,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Кола 14,5 г ",
            "product_code": "00000074581",
            "total": 2900,
            "qnt_total": 29,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Кола 14,5 г ",
            "product_code": "00000074581",
            "total": 15100,
            "qnt_total": 151,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Кола 14,5 г ",
            "product_code": "00000074581",
            "total": 5200,
            "qnt_total": 52,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Кола 14,5 г ",
            "product_code": "00000074581",
            "total": 1400,
            "qnt_total": 14,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Кола 14,5 г ",
            "product_code": "00000074581",
            "total": 900,
            "qnt_total": 9,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Кола 14,5 г ",
            "product_code": "00000074581",
            "total": 1400,
            "qnt_total": 14,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Лимон 14,5 г ",
            "product_code": "00000120964",
            "total": 3200,
            "qnt_total": 32,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Лимон 14,5 г ",
            "product_code": "00000120964",
            "total": 14900,
            "qnt_total": 149,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Лимон 14,5 г ",
            "product_code": "00000120964",
            "total": 2400,
            "qnt_total": 24,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Лимон 14,5 г ",
            "product_code": "00000120964",
            "total": 600,
            "qnt_total": 6,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Лимон 14,5 г ",
            "product_code": "00000120964",
            "total": 500,
            "qnt_total": 5,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Лимон 14,5 г ",
            "product_code": "00000120964",
            "total": 700,
            "qnt_total": 7,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Микс фрукт 14,5 г ",
            "product_code": "00000074575",
            "total": 3900,
            "qnt_total": 39,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Микс фрукт 14,5 г ",
            "product_code": "00000074575",
            "total": 12100,
            "qnt_total": 121,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Микс фрукт 14,5 г ",
            "product_code": "00000074575",
            "total": 3000,
            "qnt_total": 30,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Микс фрукт 14,5 г ",
            "product_code": "00000074575",
            "total": 300,
            "qnt_total": 3,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Микс фрукт 14,5 г ",
            "product_code": "00000074575",
            "total": 2800,
            "qnt_total": 28,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Микс фрукт 14,5 г ",
            "product_code": "00000074575",
            "total": 800,
            "qnt_total": 8,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Микс фрукт 14,5 г ",
            "product_code": "00000074575",
            "total": 700,
            "qnt_total": 7,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Мята 14,5 г ",
            "product_code": "00000074576",
            "total": 3500,
            "qnt_total": 35,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Мята 14,5 г ",
            "product_code": "00000074576",
            "total": 11000,
            "qnt_total": 110,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Мята 14,5 г ",
            "product_code": "00000074576",
            "total": 3400,
            "qnt_total": 34,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Мята 14,5 г ",
            "product_code": "00000074576",
            "total": 1600,
            "qnt_total": 16,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Мята 14,5 г ",
            "product_code": "00000074576",
            "total": 2200,
            "qnt_total": 22,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Мята 14,5 г ",
            "product_code": "00000074576",
            "total": 1900,
            "qnt_total": 19,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Перечная мята 14,5 г ",
            "product_code": "00000074582",
            "total": 3500,
            "qnt_total": 35,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Перечная мята 14,5 г ",
            "product_code": "00000074582",
            "total": 13700,
            "qnt_total": 137,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Перечная мята 14,5 г ",
            "product_code": "00000074582",
            "total": 4400,
            "qnt_total": 44,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Перечная мята 14,5 г ",
            "product_code": "00000074582",
            "total": 1900,
            "qnt_total": 19,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Перечная мята 14,5 г ",
            "product_code": "00000074582",
            "total": 2000,
            "qnt_total": 20,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Перечная мята 14,5 г ",
            "product_code": "00000074582",
            "total": 2700,
            "qnt_total": 27,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Энерджи 14,5 г ",
            "product_code": "00000074580",
            "total": 3300,
            "qnt_total": 33,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Энерджи 14,5 г ",
            "product_code": "00000074580",
            "total": 15800,
            "qnt_total": 158,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Энерджи 14,5 г ",
            "product_code": "00000074580",
            "total": 4400,
            "qnt_total": 44,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Энерджи 14,5 г ",
            "product_code": "00000074580",
            "total": 100,
            "qnt_total": 1,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Энерджи 14,5 г ",
            "product_code": "00000074580",
            "total": 1300,
            "qnt_total": 13,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Энерджи 14,5 г ",
            "product_code": "00000074580",
            "total": 800,
            "qnt_total": 8,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "7Stick жевательная резинка Энерджи 14,5 г ",
            "product_code": "00000074580",
            "total": 1400,
            "qnt_total": 14,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000137",
            "brand_name": "7Stick"
        },
        {
            "product_name": "911 бадяга 100 мл. гель ",
            "product_code": "00000021739",
            "total": 835.59,
            "qnt_total": 3,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 грибкосепт 100 мл, гель, для рук и ног ",
            "product_code": "00000020603",
            "total": 310.84,
            "qnt_total": 1,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 грибкосепт 100 мл, гель, для рук и ног ",
            "product_code": "00000020603",
            "total": 310.84,
            "qnt_total": 1,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 пиявок экстракт 100 мл, гель-бальзам, для ног ",
            "product_code": "00000011961",
            "total": 380.78,
            "qnt_total": 1,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 пиявок экстракт 100 мл, гель-бальзам, для ног ",
            "product_code": "00000011961",
            "total": 1615.68,
            "qnt_total": 4,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 пчелиный яд 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000036591",
            "total": 321.12,
            "qnt_total": 1,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 пчелиный яд 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000036591",
            "total": 1602.46,
            "qnt_total": 5,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 пчелиный яд 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000036591",
            "total": 644.8,
            "qnt_total": 2,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 пчелиный яд 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000036591",
            "total": 644.8,
            "qnt_total": 2,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 сабельник 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000011963",
            "total": 882.12,
            "qnt_total": 3,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 сабельник 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000011963",
            "total": 819.03,
            "qnt_total": 3,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 сабельник 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000011963",
            "total": 546.06,
            "qnt_total": 2,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 сабельник 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000011963",
            "total": 819.15,
            "qnt_total": 3,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 сабельник 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000011963",
            "total": 819.15,
            "qnt_total": 3,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 угрисепт 100 мл, гель, для лица ",
            "product_code": "00000011965",
            "total": 240.63,
            "qnt_total": 1,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 угрисепт 100 мл, гель, для лица ",
            "product_code": "00000011965",
            "total": 721.89,
            "qnt_total": 3,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 угрисепт 100 мл, гель, для лица ",
            "product_code": "00000011965",
            "total": 516.41,
            "qnt_total": 2,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 угрисепт 100 мл, гель, для лица ",
            "product_code": "00000011965",
            "total": 269.97,
            "qnt_total": 1,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 хондроитин 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000036592",
            "total": 1386.09,
            "qnt_total": 4,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 хондроитин 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000036592",
            "total": 1386.09,
            "qnt_total": 4,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 хондроитин 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000036592",
            "total": 693.05,
            "qnt_total": 2,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 хондроитин 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000036592",
            "total": 1384.44,
            "qnt_total": 4,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "911 чага 100 мл, гель-бальзам, для суставов ",
            "product_code": "00000029812",
            "total": 388.1,
            "qnt_total": 1,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000386",
            "brand_name": "911"
        },
        {
            "product_name": "Always прокладки Ultra Night Deo гигиенические № 12 шт ",
            "product_code": "00000053546",
            "total": 957.24,
            "qnt_total": 1,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Deo гигиенические № 12 шт ",
            "product_code": "00000053546",
            "total": 6700.68,
            "qnt_total": 7,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Deo гигиенические № 12 шт ",
            "product_code": "00000053546",
            "total": 3828.96,
            "qnt_total": 4,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Deo гигиенические № 12 шт ",
            "product_code": "00000053546",
            "total": 957.24,
            "qnt_total": 1,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Duo № 14 шт ",
            "product_code": "00000038961",
            "total": 15548.2,
            "qnt_total": 17,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Duo № 14 шт ",
            "product_code": "00000038961",
            "total": 48473.8,
            "qnt_total": 53,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Duo № 14 шт ",
            "product_code": "00000038961",
            "total": 24694.2,
            "qnt_total": 27,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Duo № 14 шт ",
            "product_code": "00000038961",
            "total": 911.68,
            "qnt_total": 1,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Quatro № 26 шт ",
            "product_code": "00000057854",
            "total": 5703.76,
            "qnt_total": 4,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Quatro № 26 шт ",
            "product_code": "00000057854",
            "total": 11407.52,
            "qnt_total": 8,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Quatro № 26 шт ",
            "product_code": "00000057854",
            "total": 9981.58,
            "qnt_total": 7,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Quatro № 26 шт ",
            "product_code": "00000057854",
            "total": 1425.94,
            "qnt_total": 1,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Single № 7 шт ",
            "product_code": "00000008406",
            "total": 4151.6,
            "qnt_total": 8,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Single № 7 шт ",
            "product_code": "00000008406",
            "total": 17644.3,
            "qnt_total": 34,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Single № 7 шт ",
            "product_code": "00000008406",
            "total": 518.95,
            "qnt_total": 1,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Night Single № 7 шт ",
            "product_code": "00000008406",
            "total": 518.95,
            "qnt_total": 1,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Duo № 20 шт ",
            "product_code": "00000038963",
            "total": 7147.56,
            "qnt_total": 7,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Duo № 20 шт ",
            "product_code": "00000038963",
            "total": 1021.08,
            "qnt_total": 1,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Duo № 20 шт ",
            "product_code": "00000038963",
            "total": 9189.72,
            "qnt_total": 9,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Duo № 20 шт ",
            "product_code": "00000038963",
            "total": 2042.16,
            "qnt_total": 2,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Duo № 20 шт ",
            "product_code": "00000038963",
            "total": 3063.24,
            "qnt_total": 3,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Duo № 20 шт ",
            "product_code": "00000038963",
            "total": 8168.64,
            "qnt_total": 8,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Duo № 20 шт ",
            "product_code": "00000038963",
            "total": 7147.56,
            "qnt_total": 7,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Quatro № 36 шт ",
            "product_code": "00000057872",
            "total": 1416.07,
            "qnt_total": 1,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Quatro № 36 шт ",
            "product_code": "00000057872",
            "total": 9912.49,
            "qnt_total": 7,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Quatro № 36 шт ",
            "product_code": "00000057872",
            "total": 5664.28,
            "qnt_total": 4,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Quatro № 36 шт ",
            "product_code": "00000057872",
            "total": 1416.07,
            "qnt_total": 1,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Quatro № 36 шт ",
            "product_code": "00000057872",
            "total": 9912.49,
            "qnt_total": 7,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Quatro № 36 шт ",
            "product_code": "00000057872",
            "total": 9912.49,
            "qnt_total": 7,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Single № 10 шт ",
            "product_code": "00000038997",
            "total": 3801.28,
            "qnt_total": 7,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Single № 10 шт ",
            "product_code": "00000038997",
            "total": 5430.4,
            "qnt_total": 10,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Single № 10 шт ",
            "product_code": "00000038997",
            "total": 3801.28,
            "qnt_total": 7,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Single № 10 шт ",
            "product_code": "00000038997",
            "total": 543.04,
            "qnt_total": 1,
            "date": "2022-01-06T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Single № 10 шт ",
            "product_code": "00000038997",
            "total": 1086.08,
            "qnt_total": 2,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Normal Plus Single № 10 шт ",
            "product_code": "00000038997",
            "total": 1629.12,
            "qnt_total": 3,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Plat Normal PI № 8 шт ",
            "product_code": "00000059779",
            "total": 510.44,
            "qnt_total": 1,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Plat Normal PI № 8 шт ",
            "product_code": "00000059779",
            "total": 1531.32,
            "qnt_total": 3,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Plat Normal PI № 8 шт ",
            "product_code": "00000059779",
            "total": 7656.6,
            "qnt_total": 15,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Plat Normal PI № 8 шт ",
            "product_code": "00000059779",
            "total": 510.44,
            "qnt_total": 1,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Collection Super Plus Duo ежедневные № 14 шт ",
            "product_code": "00000079839",
            "total": 1915.6,
            "qnt_total": 2,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Collection Super Plus Duo ежедневные № 14 шт ",
            "product_code": "00000079839",
            "total": 5746.8,
            "qnt_total": 6,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Collection Super Plus Duo ежедневные № 14 шт ",
            "product_code": "00000079839",
            "total": 3831.2,
            "qnt_total": 4,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Collection Super Plus Duo ежедневные № 14 шт ",
            "product_code": "00000079839",
            "total": 957.8,
            "qnt_total": 1,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Night  Duo ежедневные № 12 шт ",
            "product_code": "00000079840",
            "total": 10940.04,
            "qnt_total": 12,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Night  Duo ежедневные № 12 шт ",
            "product_code": "00000079840",
            "total": 7293.36,
            "qnt_total": 8,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Night  Duo ежедневные № 12 шт ",
            "product_code": "00000079840",
            "total": 911.67,
            "qnt_total": 1,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Night  Duo ежедневные № 12 шт ",
            "product_code": "00000079840",
            "total": 1823.34,
            "qnt_total": 2,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Night  Duo ежедневные № 12 шт ",
            "product_code": "00000079840",
            "total": 2735.01,
            "qnt_total": 3,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Night  Duo ежедневные № 12 шт ",
            "product_code": "00000079840",
            "total": 7293.36,
            "qnt_total": 8,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Night  Duo ежедневные № 12 шт ",
            "product_code": "00000079840",
            "total": 1823.34,
            "qnt_total": 2,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Night Single № 6 шт ",
            "product_code": "00000059777",
            "total": 1568.01,
            "qnt_total": 3,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Night Single № 6 шт ",
            "product_code": "00000059777",
            "total": 2090.68,
            "qnt_total": 4,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Night Single № 6 шт ",
            "product_code": "00000059777",
            "total": 1568.01,
            "qnt_total": 3,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Super Plus Single № 7 шт ",
            "product_code": "00000059778",
            "total": 3105.48,
            "qnt_total": 6,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Super Plus Single № 7 шт ",
            "product_code": "00000059778",
            "total": 2070.32,
            "qnt_total": 4,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Super Plus Single № 7 шт ",
            "product_code": "00000059778",
            "total": 2070.32,
            "qnt_total": 4,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Super Plus Single № 7 шт ",
            "product_code": "00000059778",
            "total": 1086.08,
            "qnt_total": 2,
            "date": "2022-01-06T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Platinum Super Plus Single № 7 шт ",
            "product_code": "00000059778",
            "total": 543.04,
            "qnt_total": 1,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super Duo Action  № 16 шт ",
            "product_code": "00000038964",
            "total": 10028.37,
            "qnt_total": 11,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super Duo Action  № 16 шт ",
            "product_code": "00000038964",
            "total": 12763.38,
            "qnt_total": 14,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super Duo Action  № 16 шт ",
            "product_code": "00000038964",
            "total": 3658.84,
            "qnt_total": 4,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super Duo Action  № 16 шт ",
            "product_code": "00000038964",
            "total": 6466.81,
            "qnt_total": 7,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super Duo Action  № 16 шт ",
            "product_code": "00000038964",
            "total": 923.83,
            "qnt_total": 1,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super Plus Single № 8 шт ",
            "product_code": "00000038967",
            "total": 4717.66,
            "qnt_total": 9,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super Plus Single № 8 шт ",
            "product_code": "00000038967",
            "total": 3801.28,
            "qnt_total": 7,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super Plus Single № 8 шт ",
            "product_code": "00000038967",
            "total": 3258.24,
            "qnt_total": 6,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super Plus Single № 8 шт ",
            "product_code": "00000038967",
            "total": 1086.08,
            "qnt_total": 2,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super QP Prospera № 30 шт ",
            "product_code": "00000057855",
            "total": 7080.4,
            "qnt_total": 5,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super QP Prospera № 30 шт ",
            "product_code": "00000057855",
            "total": 8496.48,
            "qnt_total": 6,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super QP Prospera № 30 шт ",
            "product_code": "00000057855",
            "total": 4248.24,
            "qnt_total": 3,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super QP Prospera № 30 шт ",
            "product_code": "00000057855",
            "total": 2832.16,
            "qnt_total": 2,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки Ultra Super QP Prospera № 30 шт ",
            "product_code": "00000057855",
            "total": 8496.48,
            "qnt_total": 6,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки гигиенические ультратонкие ночные экстра защита Ultra Platinum Secure Night Single ",
            "product_code": "00000122925",
            "total": 2158.54,
            "qnt_total": 4,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки гигиенические ультратонкие ночные экстра защита Ultra Platinum Secure Night Single ",
            "product_code": "00000122925",
            "total": 2545.45,
            "qnt_total": 5,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки гигиенические ультратонкие ночные экстра защита Ultra Platinum Secure Night Single ",
            "product_code": "00000122925",
            "total": 509.09,
            "qnt_total": 1,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки гигиенические ультратонкие ночные экстра защита Ultra Platinum Secure Night Single ",
            "product_code": "00000122925",
            "total": 1527.27,
            "qnt_total": 3,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки гигиенические ультратонкие ночные экстра защита Ultra Platinum Secure Night Single ",
            "product_code": "00000122925",
            "total": 3563.63,
            "qnt_total": 7,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки гигиенические ультратонкие ночные экстра защита Ultra Platinum Secure Night Single ",
            "product_code": "00000122925",
            "total": 3054.54,
            "qnt_total": 6,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки гигиенические ультратонкие ночные экстра защита Ultra Platinum Secure Night Single ",
            "product_code": "00000122925",
            "total": 509.09,
            "qnt_total": 1,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки незаметная защита ежедневные № 16 шт удлиненные ",
            "product_code": "00000077162",
            "total": 5903.1,
            "qnt_total": 10,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки незаметная защита ежедневные № 16 шт удлиненные ",
            "product_code": "00000077162",
            "total": 2951.55,
            "qnt_total": 5,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки незаметная защита ежедневные № 18 шт экстраудлиненные ",
            "product_code": "00000077160",
            "total": 1775.32,
            "qnt_total": 2,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки незаметная защита ежедневные № 18 шт экстраудлиненные ",
            "product_code": "00000077160",
            "total": 887.66,
            "qnt_total": 1,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки незаметная защита ежедневные № 18 шт экстраудлиненные ",
            "product_code": "00000077160",
            "total": 887.66,
            "qnt_total": 1,
            "date": "2022-01-06T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки незаметная защита ежедневные № 18 шт экстраудлиненные ",
            "product_code": "00000077160",
            "total": 887.66,
            "qnt_total": 1,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки незаметная защита ежедневные № 18 шт экстраудлиненные ",
            "product_code": "00000077160",
            "total": 1775.32,
            "qnt_total": 2,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки незаметная защита ежедневные № 20 шт ",
            "product_code": "00000077163",
            "total": 590.32,
            "qnt_total": 1,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки незаметная защита ежедневные № 20 шт ",
            "product_code": "00000077163",
            "total": 590.32,
            "qnt_total": 1,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Always прокладки незаметная защита ежедневные № 20 шт ",
            "product_code": "00000077163",
            "total": 590.32,
            "qnt_total": 1,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000022",
            "brand_name": "Always"
        },
        {
            "product_name": "Aura  Гель антибактериальный изопропиловый спирт туба  для рук 40 мл ",
            "product_code": "00000114746",
            "total": 2594.64,
            "qnt_total": 8,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura  Гель антибактериальный изопропиловый спирт туба  для рук 40 мл ",
            "product_code": "00000114746",
            "total": 324.33,
            "qnt_total": 1,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura  крем-мыло антибактериальное Derma Protect 2в1 250 мл ",
            "product_code": "00000078693",
            "total": 328.05,
            "qnt_total": 1,
            "date": "2022-01-10T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura  мыло мягкий уход для мам и малышей 250 мл ",
            "product_code": "00000053689",
            "total": 224.95,
            "qnt_total": 1,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura  мыло мягкий уход для мам и малышей 250 мл ",
            "product_code": "00000053689",
            "total": 341.96,
            "qnt_total": 1,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura beauty ватные диски Cotton Pads № 150 шт ",
            "product_code": "00000035832",
            "total": 769.96,
            "qnt_total": 2,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura beauty ватные диски Cotton Pads № 50 шт ",
            "product_code": "00000027503",
            "total": 178.57,
            "qnt_total": 1,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura beauty ватные диски Cotton Pads № 80 шт ",
            "product_code": "00000027504",
            "total": 243.75,
            "qnt_total": 1,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura beauty ватные диски Cotton Pads № 80 шт ",
            "product_code": "00000027504",
            "total": 975,
            "qnt_total": 4,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura beauty ватные диски Cotton Pads № 80 шт ",
            "product_code": "00000027504",
            "total": 487.5,
            "qnt_total": 2,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura beauty ватные диски Cotton Pads № 80 шт ",
            "product_code": "00000027504",
            "total": 243.75,
            "qnt_total": 1,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura beauty крем для рук  с маслом ши туба  КК/24 восстанавливающий  75 мл ",
            "product_code": "00000064203",
            "total": 4968,
            "qnt_total": 18,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura beauty крем для рук  с маслом ши туба  КК/24 восстанавливающий  75 мл ",
            "product_code": "00000064203",
            "total": 3036,
            "qnt_total": 11,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura beauty крем для рук  с маслом ши туба  КК/24 восстанавливающий  75 мл ",
            "product_code": "00000064203",
            "total": 828,
            "qnt_total": 3,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura ultra comfort влажные салфетки детские 0+ с экст алоэ и вит Е big-pack с крышкой № 100 шт ",
            "product_code": "00000053686",
            "total": 1093.47,
            "qnt_total": 3,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura ultra comfort влажные салфетки детские 0+ с экст алоэ и вит Е big-pack с крышкой № 100 шт ",
            "product_code": "00000053686",
            "total": 1093.47,
            "qnt_total": 3,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Ultra comfort влажные салфетки детские 0+ с экст алоэ и вит Е big-pack с крышкой № 120 шт ",
            "product_code": "00000064888",
            "total": 2118,
            "qnt_total": 5,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura бумага влажная туалетная №20 ",
            "product_code": "00000031281",
            "total": 250,
            "qnt_total": 1,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura бумага влажная туалетная №20 ",
            "product_code": "00000031281",
            "total": 250,
            "qnt_total": 1,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura ватные палочки Cotton Buds п/эт пакет № 200 шт ",
            "product_code": "00000024283",
            "total": 259,
            "qnt_total": 1,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura влажные носовые платочки Antibacterial pocket-pack КК/26 № 10 шт ",
            "product_code": "00000078639",
            "total": 321.44,
            "qnt_total": 4,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura влажные носовые платочки Antibacterial pocket-pack КК/26 № 10 шт ",
            "product_code": "00000078639",
            "total": 80.36,
            "qnt_total": 1,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura влажные салфетки Ромашка стикер рука pocket-pack  антибактериальные № 15 шт ",
            "product_code": "00000046971",
            "total": 826.45,
            "qnt_total": 5,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura гель antibacterial с ароматом клубники и экстр алоэ для рук 50 мл ",
            "product_code": "00000078641",
            "total": 403.2,
            "qnt_total": 1,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura гель antibacterial с ароматом клубники и экстр алоэ для рук 50 мл ",
            "product_code": "00000078641",
            "total": 806.4,
            "qnt_total": 2,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura гель antibacterial с ароматом клубники и экстр алоэ для рук 50 мл ",
            "product_code": "00000078641",
            "total": 403.2,
            "qnt_total": 1,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura гель antibacterial с ароматом клубники и экстр алоэ для рук 50 мл ",
            "product_code": "00000078641",
            "total": 403.2,
            "qnt_total": 1,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Гель для рук антибактериальный Fresh изопропиловый спирт ПРОМО для рук 50+50мл ",
            "product_code": "00000122501",
            "total": 660,
            "qnt_total": 2,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Гель для рук антибактериальный Fresh изопропиловый спирт ПРОМО для рук 50+50мл ",
            "product_code": "00000122501",
            "total": 1650,
            "qnt_total": 5,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Гель для рук антибактериальный Fresh изопропиловый спирт ПРОМО для рук 50+50мл ",
            "product_code": "00000122501",
            "total": 990,
            "qnt_total": 3,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Гель для рук антибактериальный Fresh изопропиловый спирт ПРОМО для рук 50+50мл ",
            "product_code": "00000122501",
            "total": 1320,
            "qnt_total": 4,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Гель для рук антибактериальный Fresh изопропиловый спирт ПРОМО для рук 50+50мл ",
            "product_code": "00000122501",
            "total": 660,
            "qnt_total": 2,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "AURA жидкое мыло антибактериальное  Алое вера ультра 300мл",
            "product_code": "00000045780",
            "total": 3919.6,
            "qnt_total": 10,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "AURA жидкое мыло антибактериальное  Алое вера ультра 300мл",
            "product_code": "00000045780",
            "total": 1567.84,
            "qnt_total": 4,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "AURA жидкое мыло антибактериальное  Алое вера ультра 300мл",
            "product_code": "00000045780",
            "total": 783.92,
            "qnt_total": 2,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "AURA жидкое мыло антибактериальное  Алое вера ультра 300мл",
            "product_code": "00000045780",
            "total": 2351.76,
            "qnt_total": 6,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Жидкое мыло Ромашка дой-пак  антибактериальный эф-т 500 мл ",
            "product_code": "00000121390",
            "total": 512.32,
            "qnt_total": 2,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura жидкое мыло с антибактериальным эффектом ромашка КК/12\u0009 300 мл ",
            "product_code": "00000045781",
            "total": 249.68,
            "qnt_total": 1,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura жидкое мыло с антибактериальным эффектом ромашка КК/12\u0009 300 мл ",
            "product_code": "00000045781",
            "total": 488.85,
            "qnt_total": 2,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura жидкое мыло с антибактериальным эффектом ромашка КК/12\u0009 300 мл ",
            "product_code": "00000045781",
            "total": 783.92,
            "qnt_total": 2,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura жидкое мыло с антибактериальным эффектом ромашка КК/12\u0009 300 мл ",
            "product_code": "00000045781",
            "total": 391.96,
            "qnt_total": 1,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura крем для рук Beauty увлажняющий в тубе 75 мл с глицерином и экстрактом алоэ ",
            "product_code": "00000064343",
            "total": 246.43,
            "qnt_total": 1,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura крем для рук Beauty увлажняющий в тубе 75 мл с глицерином и экстрактом алоэ ",
            "product_code": "00000064343",
            "total": 2464.3,
            "qnt_total": 10,
            "date": "2022-01-11T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura крем для рук Beauty увлажняющий в тубе 75 мл с глицерином и экстрактом алоэ ",
            "product_code": "00000064343",
            "total": 3450.02,
            "qnt_total": 14,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura крем для рук Beauty увлажняющий в тубе 75 мл с глицерином и экстрактом алоэ ",
            "product_code": "00000064343",
            "total": 246.43,
            "qnt_total": 1,
            "date": "2022-01-13T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura крем-мыло антибактериальное Derma Protect 500 мл ",
            "product_code": "00000078649",
            "total": 926.78,
            "qnt_total": 2,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura крем-мыло антибактериальное Derma Protect 500 мл ",
            "product_code": "00000078649",
            "total": 463.39,
            "qnt_total": 1,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Крем-мыло антибактериальное KIDS флакон/дозатор 250 мл ",
            "product_code": "00000121391",
            "total": 451.79,
            "qnt_total": 1,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Крем-мыло антибактериальное KIDS флакон/дозатор 250 мл ",
            "product_code": "00000121391",
            "total": 903.58,
            "qnt_total": 2,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Одноразовые покрытия для унитаза №10 ",
            "product_code": "00000026359",
            "total": 847.4,
            "qnt_total": 4,
            "date": "2022-01-03T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Одноразовые покрытия для унитаза №10 ",
            "product_code": "00000026359",
            "total": 1059.25,
            "qnt_total": 5,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Одноразовые покрытия для унитаза №10 ",
            "product_code": "00000026359",
            "total": 423.7,
            "qnt_total": 2,
            "date": "2022-01-05T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura Одноразовые покрытия для унитаза №10 ",
            "product_code": "00000026359",
            "total": 423.7,
            "qnt_total": 2,
            "date": "2022-01-12T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        },
        {
            "product_name": "Aura туалетная бумага влажная с экстрактом ромашки № 80 шт ",
            "product_code": "00000059696",
            "total": 4018.77,
            "qnt_total": 7,
            "date": "2022-01-04T00:00:00",
            "brand_code": "000000313",
            "brand_name": "Aura"
        }
    ]
}`
)
