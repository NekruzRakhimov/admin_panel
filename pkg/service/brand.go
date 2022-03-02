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
	dataBrand, contractNumber := repository.GetIDBYBIN(reqBrand.ClientBin)
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

func GetSalesSKU(reqBrand model.ReqBrand) (model.Sales, error) {
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

func PresentationDiscount(rbReq model.RBRequest) (model.Purchase, error) {
	var purchase model.Purchase

	date := model.ReqBrand{
		ClientBin:      rbReq.BIN,
		DateStart:      rbReq.PeriodFrom + TempDateCompleter,
		DateEnd:        rbReq.PeriodTo + TempDateEnd,
		Type:           "purchase",
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
		return purchase, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return purchase, err
	}
	log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return purchase, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &purchase)
	if err != nil {
		log.Println(err)
		return purchase, err
	}

	return purchase, nil

}

func InfoPresentationDiscount(rbReq model.RBRequest) []model.RbDTO {
	//..var rbBrands []model.RbDTO

	//ID                   int     `json:"id"`
	//ContractNumber       string  `json:"contract_number"`
	//StartDate            string  `json:"start_date"`
	//EndDate              string  `json:"end_date"`
	//BrandName            string  `json:"brand_name,omitempty"`
	//ProductCode          string  `json:"product_code,omitempty"`
	//DiscountPercent      float32 `json:"discount_percent"`
	//DiscountAmount       float32 `json:"discount_amount"`
	//TotalWithoutDicsount float32 `json:"TotalWithoutDiscount"`
	//LeasePlan            float32 `json:"lease_plan"`
	//RewardAmount         float32 `json:"reward_amount"`
	rbDTO := []model.RbDTO{
		{
			ContractNumber:       "9898989211",
			StartDate:            "01.01.2022",
			EndDate:              "01.02.2022",
			BrandName:            "Colgate",
			ProductCode:          "00002313",
			DiscountPercent:      5,
			DiscountAmount:       500,
			TotalWithoutDicsount: 100000,
		},
		{
			ContractNumber:       "9898989211",
			StartDate:            "01.01.2022",
			EndDate:              "01.02.2022",
			BrandName:            "Bella",
			ProductCode:          "5545454",
			DiscountPercent:      10,
			DiscountAmount:       500_000,
			TotalWithoutDicsount: 5_000_000,
		},
		{
			ContractNumber:       "11255656565",
			StartDate:            "01.02.2022",
			EndDate:              "01.04.2022",
			BrandName:            "Seni",
			ProductCode:          "065655",
			DiscountPercent:      7,
			DiscountAmount:       70_000,
			TotalWithoutDicsount: 1_000_000,
		},
	}
	return rbDTO

	//TOOD: Доработать от сюда

	////ID, contract_number, discount, bin
	//infoPresentationDiscounts := repository.GetPurchase(rbReq.BIN)
	//
	//layoutISO := "02.1.2006"
	//
	//for _, value := range infoPresentationDiscounts {
	//	if len(infoPresentationDiscounts) > 2 {
	//		//  value.StartDate - эта дата, которую мы взяли из бд
	//		// 01
	//		// 21
	//		// 10
	//		timeDB, err := time.Parse(layoutISO, value.StartDate)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		timeReq, err := time.Parse(layoutISO, value.StartDate)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		if timeDB.Before(timeReq) {
	//			//TODO: например были созданы договора за 21 и за за 25 - а ты берешь дату от 26 числа
	//			//  то да, она будет считать за 21, после этого он будет считать за 25
	//			// но есть нюансы, что сперва выпадает тебе 25 число, а не 21
	//			// ты должен учесть этот момент
	//		}
	//
	//	}
	//}
	//
	////внутри него массив
	//presentationDiscount, err := PresentationDiscount(rbReq)
	//if err != nil {
	//	log.Println(err)
	//	return rbBrand
	//}
	//totalAmount := 0
	//totalWithDiscount := 0
	//
	//for _, value := range presentationDiscount.PurchaseArr {
	//	//rbBrand.ContractNumber = infoPresentationDiscounts.ContractNumber
	//	//rbBrand.StartDate = rbReq.PeriodFrom
	//	//rbBrand.EndDate = rbReq.PeriodTo
	//	//rbBrand.BrandName = value.BrandName
	//	//rbBrand.ProductCode = value.ProductCode
	//	//rbBrand.DiscountPercent = 10
	//
	//	//TODO: подсчет общей суммы
	//	totalAmount += value.Total
	//
	//}
	//totalWithDiscount = (totalAmount * 10) / 100
	//
	////rbBrand.ContractNumber = infoPresentationDiscounts.ContractNumber
	//rbBrand.StartDate = rbReq.PeriodFrom
	//rbBrand.EndDate = rbReq.PeriodTo
	////rbBrand.BrandName = value.BrandName
	////rbBrand.ProductCode = value.ProductCode
	//rbBrand.DiscountPercent = 10
	//rbBrand.TotalWithoutDicsount = float32(totalAmount)
	//rbBrand.DiscountAmount = float32(totalWithDiscount)

	//return rbBrand

}
