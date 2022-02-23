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
		Type:           reqBrand.Type,
		TypeValue:      "",
		TypeParameters: nil,
	}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&date)
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

func FoundBrandDiscount(reqBrand model.ReqBrand) {
	f, err := excelize.OpenFile("files/reports/rb/rb_report.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// берем бренды по Бину
	dataBrand := repository.GetIDBYBIN(reqBrand.ClientBin)
	log.Println("ДАННЫЕ БРЕНДОВ", dataBrand)

	for _, value := range dataBrand {
		reqBrand.TypeParameters = append(reqBrand.TypeParameters, value.BrandName)
	}

	log.Println("reqBrand.TypeParameters", reqBrand.TypeParameters)

	var totalBrandsDiscount []model.TotalBrandDiscount
	var BrandTotal model.TotalBrandDiscount

	// reqBrand -> он дает массив брендов

	sales, _ := GetBrandSales(reqBrand)

	// Берет определенные бренды из 1С:
	counter := 1
	for _, sale := range sales.SalesArr {

		count := 0
		//	тут будет список брендов

		//TODO: тут ты можешь сразу записать в экселе наименование,кол-во,код, сумму, бренд,
		//sale.BrandName

		//TODO: после того как мы записали все бренды, мы должны посчитать от него общую сумму
		for _, brand := range reqBrand.TypeParameters {
			log.Println("BRAND", brand)

			// мы нашли схожие бренды, что мы должны сделать?
			if sale.BrandName == brand {
				count += sale.Total
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

				fmt.Println("totalBrandsDiscount", totalBrandsDiscount)
				fmt.Println("LEN:", len(totalBrandsDiscount))
				if len(totalBrandsDiscount) == 0 {
					log.Println("сРАБАОТЛО")
					BrandTotal.BrandName = brand
					BrandTotal.Amount = count
					totalBrandsDiscount = append(totalBrandsDiscount, BrandTotal)

				}

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

		TotalPercent := (value.Amount * 5) / 100
		log.Println("Сумма скидки: ", TotalPercent)
		log.Println("Скидка: ", 5)
	}

}
