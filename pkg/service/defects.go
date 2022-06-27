package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
)

const defectsSheet = "TDSheet"

const (
	getAllMatrix   = "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getallmatrix"
	getDefectsInfo = "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdefectinfo"
)

func GetDefectsExt(req models.DefectsRequest) (defects []models.Defect, err error) {
	//var binOrganizationAKNIET = "060540001442"

	//req = models.DefectsRequest{
	//	Startdate: fmt.Sprintf("%s 00:00:00", req.Startdate),
	//	Enddate:   fmt.Sprintf("%s 23:59:59", req.Enddate),
	//}

	req.Startdate += " 00:00:00"
	req.Enddate += " 23:59:59"

	bodyBin := new(bytes.Buffer)
	err = json.NewEncoder(bodyBin).Encode(&req)
	if err != nil {
		return nil, err
	}
	fmt.Println("BODY", bodyBin)

	client := &http.Client{}
	endpoint := fmt.Sprintf("http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdefect")
	r, err := http.NewRequest("POST", endpoint, bodyBin) // URL-encoded payload
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	// надо логин и пароль добавить в конфиг
	r.SetBasicAuth("http_client", "123456")

	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// ----------> часть Unmarshall json ->
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	defectsArr := struct {
		DefectArr []models.Defect `json:"defect_arr"`
	}{}

	err = json.Unmarshal(body, &defectsArr)
	if err != nil {
		return nil, err
	}

	return defectsArr.DefectArr, nil
}

func GetDefectsInfoExt(req models.DefectsInfoReq) (defectsInfo []models.DefectsInfo, err error) {
	//var binOrganizationAKNIET = "060540001442"

	//req = models.DefectsRequest{
	//	Startdate: fmt.Sprintf("%s 00:00:00", req.Startdate),
	//	Enddate:   fmt.Sprintf("%s 23:59:59", req.Enddate),
	//}

	req.Startdate += " 00:00:00"
	req.Enddate += " 00:00:00"

	bodyBin := new(bytes.Buffer)
	err = json.NewEncoder(bodyBin).Encode(&req)
	if err != nil {
		return nil, err
	}
	fmt.Println("BODY", bodyBin)

	client := &http.Client{}
	endpoint := fmt.Sprintf(getDefectsInfo)
	r, err := http.NewRequest("POST", endpoint, bodyBin) // URL-encoded payload
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	// надо логин и пароль добавить в конфиг
	r.SetBasicAuth("http_client", "123456")

	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// ----------> часть Unmarshall json ->
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	defectsArr := struct {
		DefectArr []models.DefectsInfo `json:"defect_arr"`
	}{}

	err = json.Unmarshal(body, &defectsArr)
	if err != nil {
		return nil, err
	}

	return defectsArr.DefectArr, nil
}

func GetAllMatrixExt() (arr []models.MatrixInfoFrom1C, err error) {
	//var binOrganizationAKNIET = "060540001442"

	//req = models.DefectsRequest{
	//	Startdate: fmt.Sprintf("%s 00:00:00", req.Startdate),
	//	Enddate:   fmt.Sprintf("%s 23:59:59", req.Enddate),
	//}

	client := &http.Client{}
	endpoint := fmt.Sprintf(getAllMatrix)
	r, err := http.NewRequest("GET", endpoint, nil) // URL-encoded payload
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	// надо логин и пароль добавить в конфиг
	r.SetBasicAuth("http_client", "123456")

	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// ----------> часть Unmarshall json ->
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	defectsArr := struct {
		MatrixArr []models.MatrixInfoFrom1C `json:"matrix_arr"`
	}{}

	err = json.Unmarshal(body, &defectsArr)
	if err != nil {
		return nil, err
	}

	return defectsArr.MatrixArr, nil
}

func SaveAllMatrixFrom1C(matrix []models.MatrixInfoFrom1C) error {
	//matrix, err := GetAllMatrixExt()
	//if err != nil {
	//	return err
	//}

	var addedStores []string
	var stores []models.Store
	for _, m := range matrix {
		if !stringContains(addedStores, m.StoreCode) {
			addedStores = append(addedStores, m.StoreCode)
			stores = append(stores, models.Store{
				StoreName: m.StoreName,
				StoreCode: m.StoreCode,
			})
		}
	}

	return repository.SaveAllStoresFrom1C(stores)
}

func GetDefectsLS(req models.DefectsRequest) (filteredDefects []models.DefectsFiltered, err error) {
	//var filteredDefects []models.DefectsFiltered
	req.IsPF = false
	req.DaysCount = -60
	req.QueryType = "warehouse_defect"
	defects, err := GetDefectsExt(req)
	if err != nil {
		return nil, err
	}
	var checkedCodes []string

	for i, defect := range defects {
		//isPF, err := strconv.ParseBool(defect.PF)
		//if err != nil {
		//	return nil, err
		//}

		//if isPF {
		//	continue
		//}

		if !StrArrContainsEl(checkedCodes, defect.StoreCode) {
			defectsFiltered := models.DefectsFiltered{
				StoreCode: defect.StoreCode,
				StoreName: defect.StoreName,
			}
			for j := i; j < len(defects); j++ {
				if defect.StoreCode == defects[j].StoreCode {
					defectsFiltered.SubDefects = append(defectsFiltered.SubDefects, defects[j])
				}
			}
			checkedCodes = append(checkedCodes, defect.StoreCode)
			filteredDefects = append(filteredDefects, defectsFiltered)
		}
	}

	if err = FormExcelDefectsLS(req, filteredDefects); err != nil {
		return nil, err
	}

	return filteredDefects, nil
}

func StrArrContainsEl(arr []string, el string) bool {
	for _, s := range arr {
		if s == el {
			return true
		}
	}

	return false
}

func FormExcelDefectsLS(req models.DefectsRequest, filteredDefects []models.DefectsFiltered) error {
	f, err := excelize.OpenFile("files/defects/defects_pharmacy_template_ls.xlsx")
	if err != nil {
		return err
	}

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		NumFmt: 4,
	})
	moneyStyle, _ := f.NewStyle(`{"number_format": 4}`)
	moneyMainStyle, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#FFFAD9"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		NumFmt: 4,
	})

	//xlsx.SetCellStyle("Sheet1", "B4", "B4", style)

	//f.NewSheet(defectsSheet)
	f.SetCellValue(defectsSheet, "D1", req.Startdate) //Дата
	f.SetCellValue(defectsSheet, "H1", req.Startdate) //Дата
	var (
		globalDefectSum1 float64
		globalDefectSum2 float64
	)
	var i = 4
	for _, defect := range filteredDefects {
		storeIndex := i
		f.SetCellValue(defectsSheet, fmt.Sprintf("A%d", storeIndex), defect.StoreName) //Аптека
		var (
			storeMatrixSales     float64 // кол-во СКЮ продаваемых за 60 дней по матрице
			storeDefectQnt       int     //кол-во СКЮ в дефектуре
			storeDefectSum       float64 //сумма дефектуры
			storeFactSaleDefect  float64 //% дефектуры от факт продаж
			storeSkuQnt          int     //кол-во СКЮ входящих в АМ аптеки/магазина
			storeDefectSkuQnt    int     //кол-во СКЮ в дефектуре по АМ ПФ
			storeDefectSum2      float64 //сумма дефектуры
			storeDefectAM        float64 //% дефектуры от АМ
			storeStoreSaldoQnt   float64 // наличие продукции на складе
			storeStoreSaldoCount float64 // наличие продукции на складе - в суммарном выражении
		)
		i++
		for _, subDefect := range defect.SubDefects {
			f.SetCellValue(defectsSheet, fmt.Sprintf("A%d", i), defect.StoreName)
			storeSaldoQnt, _ := strconv.ParseFloat(subDefect.StoreSaldoQnt, 2)
			if err != nil {
				return err
			}

			defectQnt, _ := strconv.Atoi(subDefect.DefectQnt)
			if err != nil {
				return err
			}

			matrixProductQnt, _ := strconv.Atoi(subDefect.MatrixProductQnt)
			if err != nil {
				return err
			}

			price, _ := strconv.ParseFloat(subDefect.DefectPrice, 2)
			if err != nil {
				return err
			}

			matrixSales, _ := strconv.ParseFloat(subDefect.MatrixSales, 2)
			if err != nil {
				return err
			}

			f.SetCellValue(defectsSheet, fmt.Sprintf("B%d", i), subDefect.ProductName) //Наименование
			f.SetCellValue(defectsSheet, fmt.Sprintf("C%d", i), subDefect.ProductCode) //код 1С

			f.SetCellValue(defectsSheet, fmt.Sprintf("D%d", i), matrixSales) // кол-во СКЮ продаваемых за 60 дней по матрице
			storeMatrixSales += matrixSales

			f.SetCellValue(defectsSheet, fmt.Sprintf("E%d", i), defectQnt) //кол-во СКЮ в дефектуре
			storeDefectQnt += defectQnt

			f.SetCellValue(defectsSheet, fmt.Sprintf("F%d", i), float64(defectQnt)*price) //сумма дефектуры
			storeDefectSum += float64(defectQnt) * price

			if matrixSales == 0 {
				f.SetCellValue(defectsSheet, fmt.Sprintf("G%d", i), 0) //% дефектуры от факт продаж
			} else {
				f.SetCellValue(defectsSheet, fmt.Sprintf("G%d", i), float64(defectQnt)/matrixSales*100) //% дефектуры от факт продаж
			}
			storeFactSaleDefect += float64(defectQnt) / matrixSales * 100

			f.SetCellValue(defectsSheet, fmt.Sprintf("H%d", i), matrixProductQnt) //кол-во СКЮ входящих в АМ аптеки/магазина
			storeSkuQnt += matrixProductQnt

			f.SetCellValue(defectsSheet, fmt.Sprintf("I%d", i), len(defect.SubDefects)) //кол-во СКЮ в дефектуре по АМ ПФ
			storeDefectSkuQnt = len(defect.SubDefects)

			f.SetCellValue(defectsSheet, fmt.Sprintf("J%d", i), float64(len(defect.SubDefects))*price) //сумма дефектуры
			storeDefectSum2 += float64(len(defect.SubDefects)) * price

			f.SetCellValue(defectsSheet, fmt.Sprintf("K%d", i), float64(len(defect.SubDefects))/float64(matrixProductQnt)*100) //% дефектуры от АМ
			storeDefectAM += (float64(len(defect.SubDefects)) * price) * float64(matrixProductQnt)

			f.SetCellValue(defectsSheet, fmt.Sprintf("L%d", i), storeSaldoQnt) // наличие продукции на складе
			storeStoreSaldoQnt += storeSaldoQnt

			f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), storeSaldoQnt*price) // наличие продукции на складе - в суммарном выражении
			storeStoreSaldoCount += storeSaldoQnt * price

			f.SetCellValue(defectsSheet, fmt.Sprintf("N%d", i), price) // Закупочная цена

			f.SetCellStyle(defectsSheet, fmt.Sprintf("D%d", i), fmt.Sprintf("N%d", i), moneyStyle)
			i++
		}

		f.SetCellValue(defectsSheet, fmt.Sprintf("D%d", storeIndex), storeMatrixSales)                                    // кол-во СКЮ продаваемых за 60 дней по матрице
		f.SetCellValue(defectsSheet, fmt.Sprintf("E%d", storeIndex), storeDefectQnt)                                      //кол-во СКЮ в дефектуре
		f.SetCellValue(defectsSheet, fmt.Sprintf("F%d", storeIndex), storeDefectSum)                                      //сумма дефектуры
		f.SetCellValue(defectsSheet, fmt.Sprintf("G%d", storeIndex), float64(storeDefectQnt)/storeMatrixSales*100)        //% дефектуры от факт продаж
		f.SetCellValue(defectsSheet, fmt.Sprintf("H%d", storeIndex), storeSkuQnt)                                         //кол-во СКЮ входящих в АМ аптеки/магазина
		f.SetCellValue(defectsSheet, fmt.Sprintf("I%d", storeIndex), storeDefectSkuQnt)                                   //кол-во СКЮ в дефектуре по АМ ПФ
		f.SetCellValue(defectsSheet, fmt.Sprintf("J%d", storeIndex), storeDefectSum2)                                     //сумма дефектуры
		f.SetCellValue(defectsSheet, fmt.Sprintf("K%d", storeIndex), float64(storeDefectSkuQnt)/float64(storeSkuQnt)*100) //% дефектуры от АМ
		//if storeDefectSkuQnt != 0 {
		//	f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), float64(storeDefectSkuQnt)/float64(storeSkuQnt)) //% дефектуры от АМ
		//} else {
		//	f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), 0) //% дефектуры от АМ
		//}
		f.SetCellValue(defectsSheet, fmt.Sprintf("L%d", storeIndex), storeStoreSaldoQnt)   // наличие продукции на складе
		f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", storeIndex), storeStoreSaldoCount) // наличие продукции на складе - в суммарном выражении
		f.SetCellStyle(defectsSheet, fmt.Sprintf("A%d", storeIndex), fmt.Sprintf("M%d", storeIndex), style)
		//f.SetCellStyle(defectsSheet, fmt.Sprintf("D%d", storeIndex), fmt.Sprintf("N%d", storeIndex), moneyStyle)

		globalDefectSum1 += storeDefectSum
		globalDefectSum2 += storeDefectSum2
	}

	f.SetCellValue(defectsSheet, "F3", globalDefectSum1)
	f.SetCellStyle(defectsSheet, "F3", "F3", moneyMainStyle)
	f.SetCellValue(defectsSheet, "J3", globalDefectSum2)
	f.SetCellStyle(defectsSheet, "J3", "J3", moneyMainStyle)

	f.DeleteSheet("Sheet1")
	_ = f.SaveAs("files/defects/res_ls.xlsx")
	//f.Close()
	return nil
}

func GetSalesCountExt(req models.SalesCountRequest) (defects []models.SalesCount, err error) {
	//var binOrganizationAKNIET = "060540001442"

	response := struct {
		SalesCountArr []models.SalesCount `json:"sales_count_arr"`
	}{}

	fmt.Println("Started marshalling ext_req_body")

	bodyBin, err := json.Marshal(&req)
	//bodyBin := new(bytes.Buffer)
	//err = json.NewEncoder(bodyBin).Encode(&req)
	if err != nil {
		return nil, err
	}
	fmt.Println("body:", string(bodyBin))
	fmt.Println("Finished marshalling ext_req_body")
	//fmt.Println("BODY", bodyBin)

	fmt.Println("Started sending ext_req")
	client := &http.Client{}
	endpoint := fmt.Sprintf("http://89.218.153.38:8081/AQG_ULAN/hs/integration/salescount")
	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(bodyBin)) // URL-encoded payload
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	// надо логин и пароль добавить в конфиг
	r.SetBasicAuth("http_client", "123456")

	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	fmt.Println("Finished sending ext_req")

	defer res.Body.Close()
	fmt.Println("Started reading ext_response")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Finished reading ext_response")

	// ----------> часть Unmarshall json ->
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	fmt.Println("BODY", string(body))
	fmt.Println("Started unmarshalling ext_response")
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	fmt.Println("Finished unmarshalling ext_response")

	return response.SalesCountArr, nil
}

func GetNewDefects(req models.DefectsRequest) error {
	return NewFormExcelDefectsPF(req)
}

func NewFormExcelDefectsPF(req models.DefectsRequest) error {
	var restriction int
	f, err := excelize.OpenFile("files/defects/defects_pharmacy_template.xlsx")
	if err != nil {
		return err
	}

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		NumFmt: 4,
	})
	//moneyStyle, _ := f.NewStyle(`{"number_format": 4}`)
	moneyStyle, _ := f.NewStyle(&excelize.Style{
		//Fill: excelize.Fill{Type: "pattern", Color: []string{"#FFFAD9"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		NumFmt: 4,
	})
	moneyMainStyle, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#FFFAD9"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		NumFmt: 4,
	})

	f.NewSheet(defectsSheet)
	f.SetCellValue(defectsSheet, "D1", req.Startdate) //Дата
	f.SetCellValue(defectsSheet, "H1", req.Startdate) //Дата

	var (
		globalDefectSum1 float64
		globalDefectSum2 float64
	)
	var i = 4
	stores, err := repository.GetAllStores()
	if err != nil {
		return err
	}

	for _, store := range stores {

		storeIndex := i
		f.SetCellValue(defectsSheet, fmt.Sprintf("A%d", storeIndex), store.StoreName) //Аптека
		f.SetCellStyle(defectsSheet, fmt.Sprintf("A%d", storeIndex), fmt.Sprintf("A%d", storeIndex), style)
		var (
			storeMatrixSales float64 // кол-во СКЮ продаваемых за 60 дней по матрице
			storeDefectQnt   float64 //кол-во СКЮ в дефектуре
			storeDefectSum   float64 //сумма дефектуры
			//storeFactSaleDefect  float64 //% дефектуры от факт продаж
			storeSkuQnt       float64 //кол-во СКЮ входящих в АМ аптеки/магазина
			storeDefectSkuQnt float64 //кол-во СКЮ в дефектуре по АМ ПФ
			storeDefectSum2   float64 //сумма дефектуры
			//storeDefectAM        float64 //% дефектуры от АМ
			storeStoreSaldoQnt   float64 // наличие продукции на складе
			storeStoreSaldoCount float64 // наличие продукции на складе - в суммарном выражении
		)
		i++
		j := i

		fmt.Println("[GetMatrixExt]")
		matrix, err := GetMatrixExt(store.StoreCode)
		if err != nil {
			return err
		}

		fmt.Println("[GetSalesCountExt]")
		salesCountReq := models.SalesCountRequest{Startdate: req.Startdate + " 00:00:00", Enddate: req.Enddate + " 00:00:00", StoreCode: store.StoreCode}
		sales, err := GetSalesCountExt(salesCountReq)
		if err != nil {
			return err
		}

		fmt.Println("[GetDefectsInfoExt]")
		reqDefectsInfo := models.DefectsInfoReq{Startdate: req.Startdate, Enddate: req.Enddate, StoreCode: store.StoreCode, Products: nil}
		defectsInfo, err := GetDefectsInfoExt(reqDefectsInfo)
		if err != nil {
			return err
		}

		var productFullInfo []models.ProductFullInfo
		for _, m := range matrix {
			defectInfo := models.DefectsInfo{}
			for _, info := range defectsInfo {
				if info.ProductCode == m.ProductCode {
					defectInfo = info
				}
			}

			sale := models.SalesCount{}
			for _, saleCount := range sales {
				if saleCount.ProductCode == m.ProductCode {
					sale = saleCount
				}
			}

			productFullInfo = append(productFullInfo, models.ProductFullInfo{
				Product:     m.ProductCode,
				Matrix:      m,
				SalesCount:  sale,
				DefectsInfo: defectInfo,
			})
		}

		for _, fullInfo := range productFullInfo {
			m := fullInfo.Matrix

			defectInfo := fullInfo.DefectsInfo

			sale := fullInfo.SalesCount

			matrixSales, _ := strconv.ParseFloat(defectInfo.MatrixSalesQnt, 2)
			//saleCount, _ := strconv.ParseFloat(sale.SalesCount, 2)
			salesDayCount, _ := strconv.ParseFloat(sale.SalesDayCount, 2)
			totalStoreCount, _ := strconv.ParseFloat(sale.TotalStoreCount, 2)
			min, _ := strconv.ParseFloat(m.Min, 2)
			max, _ := strconv.ParseFloat(m.Max, 2)
			price, _ := strconv.ParseFloat(defectInfo.DefectPrice, 2)
			storeSaldoQnt, _ := strconv.ParseFloat(defectInfo.StoreSaldoQnt, 2)

			//Аптека
			f.SetCellValue(defectsSheet, fmt.Sprintf("A%d", i), store.StoreName)

			//Наименование
			f.SetCellValue(defectsSheet, fmt.Sprintf("B%d", i), m.ProductName)

			//код 1С
			f.SetCellValue(defectsSheet, fmt.Sprintf("C%d", i), m.ProductCode)

			// кол-во СКЮ продаваемых за 60 дней по матрице
			f.SetCellValue(defectsSheet, fmt.Sprintf("D%d", i), matrixSales)
			if matrixSales != 0 {
				storeMatrixSales++
			}

			//кол-во СКЮ в дефектуре
			defect := 0.00
			if max != 0 {
				defect = max - totalStoreCount
			} else {
				if salesDayCount == 0 {
					defect = matrixSales*15 + min - totalStoreCount
				} else {
					defect = matrixSales/salesDayCount*15 + min - totalStoreCount
				}
			}
			f.SetCellValue(defectsSheet, fmt.Sprintf("E%d", i), defect)
			if m.ProductCode == "00000076486" {
				log.Println("[start_analyze]")
				log.Printf("[00000076486] %v\n", m)
				log.Printf("[matrixSales] %v\n", matrixSales)
				log.Printf("[salesDayCount] %v\n", salesDayCount)
				log.Printf("[min] %v\n", min)
				log.Printf("[totalStoreCount] %v\n", totalStoreCount)
				log.Printf("[defect] %v\n", defect)
				log.Println("[finish_analyze]")
			}
			if defect > 0 {
				storeDefectQnt++
			}

			//сумма дефектуры
			f.SetCellValue(defectsSheet, fmt.Sprintf("F%d", i), defect*price)
			storeDefectSum += defect * price

			//кол-во СКЮ входящих в АМ аптеки/магазина
			f.SetCellValue(defectsSheet, fmt.Sprintf("H%d", i), 1)
			storeSkuQnt++

			//кол-во СКЮ в дефектуре по АМ ПФ
			hasDefect := 0
			if defect > 0 {
				storeDefectSkuQnt++
				hasDefect = 1
			}
			f.SetCellValue(defectsSheet, fmt.Sprintf("I%d", i), hasDefect)

			//сумма дефектуры
			f.SetCellValue(defectsSheet, fmt.Sprintf("J%d", i), defect*price)
			if hasDefect != 0 {
				storeDefectSum2 += price
			}

			// наличие продукции на складе
			f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), storeSaldoQnt)
			storeStoreSaldoQnt += storeSaldoQnt

			// наличие продукции на складе - в суммарном выражении
			f.SetCellValue(defectsSheet, fmt.Sprintf("N%d", i), storeSaldoQnt*price)
			storeStoreSaldoCount += storeSaldoQnt * price

			// Закупочная цена
			f.SetCellValue(defectsSheet, fmt.Sprintf("O%d", i), price)

			// min
			f.SetCellValue(defectsSheet, fmt.Sprintf("P%d", i), min)

			// max
			f.SetCellValue(defectsSheet, fmt.Sprintf("Q%d", i), max)

			//Остаток
			f.SetCellValue(defectsSheet, fmt.Sprintf("R%d", i), totalStoreCount)

			// Дни продаж
			f.SetCellValue(defectsSheet, fmt.Sprintf("S%d", i), salesDayCount)

			f.SetCellStyle(defectsSheet, fmt.Sprintf("D%d", i), fmt.Sprintf("K%d", i), moneyStyle)
			f.SetCellStyle(defectsSheet, fmt.Sprintf("M%d", i), fmt.Sprintf("S%d", i), moneyStyle)
			i++
		}

		//% дефектуры от факт продаж
		for _, fullInfo := range productFullInfo {
			m := fullInfo.Matrix

			defectInfo := fullInfo.DefectsInfo

			sale := fullInfo.SalesCount

			saleCount, _ := strconv.ParseFloat(sale.SalesCount, 2)
			salesDayCount, _ := strconv.ParseFloat(sale.SalesDayCount, 2)
			totalStoreCount, _ := strconv.ParseFloat(sale.TotalStoreCount, 2)
			min, _ := strconv.ParseFloat(m.Min, 2)
			max, _ := strconv.ParseFloat(m.Max, 2)
			price, _ := strconv.ParseFloat(defectInfo.DefectPrice, 2)

			//кол-во СКЮ в дефектуре
			var defect float64
			if max != 0 {
				defect = max - totalStoreCount
			} else {
				if salesDayCount == 0 {
					salesDayCount++
				}
				defect = saleCount/salesDayCount*15 + min - totalStoreCount
			}

			if storeDefectQnt == 0 {
				storeDefectQnt++
			}
			//% дефектуры от факт продаж
			factSalePercent := fmt.Sprintf("%.2f", defect*price*100/storeDefectQnt)
			styledFactSalePercent, _ := strconv.ParseFloat(factSalePercent, 2)
			f.SetCellValue(defectsSheet, fmt.Sprintf("G%d", j), styledFactSalePercent) //% дефектуры от факт продаж

			if storeDefectSkuQnt == 0 {
				storeDefectSkuQnt++
			}
			defAm := fmt.Sprintf("%.2f", defect*price*100/storeDefectSkuQnt)
			styledDefAm, _ := strconv.ParseFloat(defAm, 2)
			f.SetCellValue(defectsSheet, fmt.Sprintf("K%d", j), styledDefAm) //% дефектуры от АМ

			j++
		}

		f.SetCellValue(defectsSheet, fmt.Sprintf("D%d", storeIndex), storeMatrixSales) // кол-во СКЮ продаваемых за 60 дней по матрице

		f.SetCellValue(defectsSheet, fmt.Sprintf("E%d", storeIndex), storeDefectQnt) //кол-во СКЮ в дефектуре

		f.SetCellValue(defectsSheet, fmt.Sprintf("F%d", storeIndex), storeDefectSum) //сумма дефектуры

		if storeMatrixSales == 0 {
			storeMatrixSales++
		}

		f.SetCellValue(defectsSheet, fmt.Sprintf("G%d", storeIndex), math.Abs(storeDefectQnt/storeMatrixSales*100)) //% дефектуры от факт продаж

		f.SetCellValue(defectsSheet, fmt.Sprintf("H%d", storeIndex), len(matrix)) //кол-во СКЮ входящих в АМ аптеки/магазина

		f.SetCellValue(defectsSheet, fmt.Sprintf("I%d", storeIndex), storeDefectSkuQnt) //кол-во СКЮ в дефектуре по АМ ПФ

		f.SetCellValue(defectsSheet, fmt.Sprintf("J%d", storeIndex), storeDefectSum2) //сумма дефектуры

		if storeSkuQnt == 0 {
			storeSkuQnt++
		}
		f.SetCellValue(defectsSheet, fmt.Sprintf("K%d", storeIndex), math.Abs(storeDefectSkuQnt/storeSkuQnt*100)) //% дефектуры от АМ
		//if storeDefectSkuQnt != 0 {
		//	f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), float64(storeDefectSkuQnt)/float64(storeSkuQnt)) //% дефектуры от АМ
		//} else {
		//	f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), 0) //% дефектуры от АМ
		//}
		f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", storeIndex), storeStoreSaldoQnt) // наличие продукции на складе
		//stream.SetRow(fmt.Sprintf("L%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeStoreSaldoQnt)}})

		f.SetCellValue(defectsSheet, fmt.Sprintf("N%d", storeIndex), storeStoreSaldoCount) // наличие продукции на складе - в суммарном выражении
		//stream.SetRow(fmt.Sprintf("M%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeStoreSaldoCount)}})

		//f.SetCellStyle(defectsSheet, fmt.Sprintf("A%d", storeIndex), fmt.Sprintf("M%d", storeIndex), style)
		f.SetCellStyle(defectsSheet, fmt.Sprintf("A%d", storeIndex), fmt.Sprintf("K%d", storeIndex), style)
		f.SetCellStyle(defectsSheet, fmt.Sprintf("M%d", storeIndex), fmt.Sprintf("S%d", storeIndex), style)
		globalDefectSum1 += storeDefectQnt
		globalDefectSum2 += storeDefectSkuQnt
		if restriction > 1 {
			break
		}
		restriction++
	}

	f.SetCellValue(defectsSheet, "E3", globalDefectSum1)
	f.SetCellStyle(defectsSheet, "E3", "E3", moneyMainStyle)
	f.SetCellValue(defectsSheet, "I3", globalDefectSum2)
	f.SetCellStyle(defectsSheet, "I3", "I3", moneyMainStyle)

	f.DeleteSheet("Sheet1")

	_ = f.SaveAs("files/defects/res.xlsx")
	return nil
}

func OrderDefectsPF(req models.DefectsRequest) error {
	order, err := repository.CreateDefectsOrder(models.DefectOrder{Date: req.Startdate, Status: "в процессе"})
	if err != nil {
		return err
	}

	order.FileName, err = FormExcelDefectsPF(req, order.ID)
	if err != nil {
		return err
	}

	order.Status = "сформирован"
	if err = repository.SaveFormedDefect(order); err != nil {
		return err
	}

	return err
}

func FormExcelDefectsPF(req models.DefectsRequest, id int) (fileName string, err error) {
	//var restriction int
	f, err := excelize.OpenFile("files/defects/defects_pharmacy_template.xlsx")
	if err != nil {
		return "", err
	}

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		NumFmt: 4,
	})
	//moneyStyle, _ := f.NewStyle(`{"number_format": 4}`)
	moneyStyle, _ := f.NewStyle(&excelize.Style{
		//Fill: excelize.Fill{Type: "pattern", Color: []string{"#FFFAD9"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		NumFmt: 4,
	})
	moneyMainStyle, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#FFFAD9"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		NumFmt: 4,
	})

	f.NewSheet(defectsSheet)
	f.SetCellValue(defectsSheet, "D1", req.Startdate) //Дата
	f.SetCellValue(defectsSheet, "H1", req.Startdate) //Дата

	var (
		globalDefectSum1 float64
		globalDefectSum2 float64
	)
	var i = 4
	stores, err := repository.GetAllStores()
	if err != nil {
		return "", err
	}

	for _, store := range stores {

		storeIndex := i
		f.SetCellValue(defectsSheet, fmt.Sprintf("A%d", storeIndex), store.StoreName) //Аптека
		f.SetCellStyle(defectsSheet, fmt.Sprintf("A%d", storeIndex), fmt.Sprintf("A%d", storeIndex), style)
		var (
			storeMatrixSales float64 // кол-во СКЮ продаваемых за 60 дней по матрице
			storeDefectQnt   float64 //кол-во СКЮ в дефектуре
			storeDefectSum   float64 //сумма дефектуры
			//storeFactSaleDefect  float64 //% дефектуры от факт продаж
			storeSkuQnt       float64 //кол-во СКЮ входящих в АМ аптеки/магазина
			storeDefectSkuQnt float64 //кол-во СКЮ в дефектуре по АМ ПФ
			storeDefectSum2   float64 //сумма дефектуры
			//storeDefectAM        float64 //% дефектуры от АМ
			storeStoreSaldoQnt   float64 // наличие продукции на складе
			storeStoreSaldoCount float64 // наличие продукции на складе - в суммарном выражении
		)
		i++
		j := i

		fmt.Println("[GetMatrixExt]")
		matrix, err := GetMatrixExt(store.StoreCode)
		if err != nil {
			return "", err
		}

		fmt.Println("[GetSalesCountExt]")
		salesCountReq := models.SalesCountRequest{Startdate: req.Startdate + " 00:00:00", Enddate: req.Enddate + " 00:00:00", StoreCode: store.StoreCode}
		sales, err := GetSalesCountExt(salesCountReq)
		if err != nil {
			return "", err
		}

		fmt.Println("[GetDefectsInfoExt]")
		reqDefectsInfo := models.DefectsInfoReq{Startdate: req.Startdate, Enddate: req.Enddate, StoreCode: store.StoreCode, Products: nil}
		defectsInfo, err := GetDefectsInfoExt(reqDefectsInfo)
		if err != nil {
			return "", err
		}

		var productFullInfo []models.ProductFullInfo
		for _, m := range matrix {
			defectInfo := models.DefectsInfo{}
			for _, info := range defectsInfo {
				if info.ProductCode == m.ProductCode {
					defectInfo = info
				}
			}

			sale := models.SalesCount{}
			for _, saleCount := range sales {
				if saleCount.ProductCode == m.ProductCode {
					sale = saleCount
				}
			}

			productFullInfo = append(productFullInfo, models.ProductFullInfo{
				Product:     m.ProductCode,
				Matrix:      m,
				SalesCount:  sale,
				DefectsInfo: defectInfo,
			})
		}

		for _, fullInfo := range productFullInfo {
			m := fullInfo.Matrix

			defectInfo := fullInfo.DefectsInfo

			sale := fullInfo.SalesCount

			matrixSales, _ := strconv.ParseFloat(defectInfo.MatrixSalesQnt, 2)
			//saleCount, _ := strconv.ParseFloat(sale.SalesCount, 2)
			salesDayCount, _ := strconv.ParseFloat(sale.SalesDayCount, 2)
			totalStoreCount, _ := strconv.ParseFloat(sale.TotalStoreCount, 2)
			min, _ := strconv.ParseFloat(m.Min, 2)
			max, _ := strconv.ParseFloat(m.Max, 2)
			price, _ := strconv.ParseFloat(defectInfo.DefectPrice, 2)
			storeSaldoQnt, _ := strconv.ParseFloat(defectInfo.StoreSaldoQnt, 2)

			//Аптека
			f.SetCellValue(defectsSheet, fmt.Sprintf("A%d", i), store.StoreName)

			//Наименование
			f.SetCellValue(defectsSheet, fmt.Sprintf("B%d", i), m.ProductName)

			//код 1С
			f.SetCellValue(defectsSheet, fmt.Sprintf("C%d", i), m.ProductCode)

			// кол-во СКЮ продаваемых за 60 дней по матрице
			f.SetCellValue(defectsSheet, fmt.Sprintf("D%d", i), matrixSales)
			if matrixSales != 0 {
				storeMatrixSales++
			}

			//кол-во СКЮ в дефектуре
			defect := 0.00
			if max != 0 {
				defect = max - totalStoreCount
			} else {
				if salesDayCount == 0 {
					defect = matrixSales*15 + min - totalStoreCount
				} else {
					defect = matrixSales/salesDayCount*15 + min - totalStoreCount
				}
			}
			f.SetCellValue(defectsSheet, fmt.Sprintf("E%d", i), defect)
			if m.ProductCode == "00000076486" {
				log.Println("[start_analyze]")
				log.Printf("[00000076486] %v\n", m)
				log.Printf("[matrixSales] %v\n", matrixSales)
				log.Printf("[salesDayCount] %v\n", salesDayCount)
				log.Printf("[min] %v\n", min)
				log.Printf("[totalStoreCount] %v\n", totalStoreCount)
				log.Printf("[defect] %v\n", defect)
				log.Println("[finish_analyze]")
			}
			if defect > 0 {
				storeDefectQnt++
			}

			//сумма дефектуры
			f.SetCellValue(defectsSheet, fmt.Sprintf("F%d", i), defect*price)
			storeDefectSum += defect * price

			//кол-во СКЮ входящих в АМ аптеки/магазина
			f.SetCellValue(defectsSheet, fmt.Sprintf("H%d", i), 1)
			storeSkuQnt++

			//кол-во СКЮ в дефектуре по АМ ПФ
			hasDefect := 0
			if defect > 0 {
				storeDefectSkuQnt++
				hasDefect = 1
			}
			f.SetCellValue(defectsSheet, fmt.Sprintf("I%d", i), hasDefect)

			//сумма дефектуры
			f.SetCellValue(defectsSheet, fmt.Sprintf("J%d", i), defect*price)
			if hasDefect != 0 {
				storeDefectSum2 += price
			}

			// наличие продукции на складе
			f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), storeSaldoQnt)
			storeStoreSaldoQnt += storeSaldoQnt

			// наличие продукции на складе - в суммарном выражении
			f.SetCellValue(defectsSheet, fmt.Sprintf("N%d", i), storeSaldoQnt*price)
			storeStoreSaldoCount += storeSaldoQnt * price

			// Закупочная цена
			f.SetCellValue(defectsSheet, fmt.Sprintf("O%d", i), price)

			// min
			f.SetCellValue(defectsSheet, fmt.Sprintf("P%d", i), min)

			// max
			f.SetCellValue(defectsSheet, fmt.Sprintf("Q%d", i), max)

			//Остаток
			f.SetCellValue(defectsSheet, fmt.Sprintf("R%d", i), totalStoreCount)

			// Дни продаж
			f.SetCellValue(defectsSheet, fmt.Sprintf("S%d", i), salesDayCount)

			f.SetCellStyle(defectsSheet, fmt.Sprintf("D%d", i), fmt.Sprintf("K%d", i), moneyStyle)
			f.SetCellStyle(defectsSheet, fmt.Sprintf("M%d", i), fmt.Sprintf("S%d", i), moneyStyle)
			i++
		}

		//% дефектуры от факт продаж
		for _, fullInfo := range productFullInfo {
			m := fullInfo.Matrix

			defectInfo := fullInfo.DefectsInfo

			sale := fullInfo.SalesCount

			saleCount, _ := strconv.ParseFloat(sale.SalesCount, 2)
			salesDayCount, _ := strconv.ParseFloat(sale.SalesDayCount, 2)
			totalStoreCount, _ := strconv.ParseFloat(sale.TotalStoreCount, 2)
			min, _ := strconv.ParseFloat(m.Min, 2)
			max, _ := strconv.ParseFloat(m.Max, 2)
			price, _ := strconv.ParseFloat(defectInfo.DefectPrice, 2)

			//кол-во СКЮ в дефектуре
			var defect float64
			if max != 0 {
				defect = max - totalStoreCount
			} else {
				if salesDayCount == 0 {
					salesDayCount++
				}
				defect = saleCount/salesDayCount*15 + min - totalStoreCount
			}

			if storeDefectQnt == 0 {
				storeDefectQnt++
			}
			//% дефектуры от факт продаж
			factSalePercent := fmt.Sprintf("%.2f", defect*price*100/storeDefectQnt)
			styledFactSalePercent, _ := strconv.ParseFloat(factSalePercent, 2)
			f.SetCellValue(defectsSheet, fmt.Sprintf("G%d", j), styledFactSalePercent) //% дефектуры от факт продаж

			if storeDefectSkuQnt == 0 {
				storeDefectSkuQnt++
			}
			defAm := fmt.Sprintf("%.2f", defect*price*100/storeDefectSkuQnt)
			styledDefAm, _ := strconv.ParseFloat(defAm, 2)
			f.SetCellValue(defectsSheet, fmt.Sprintf("K%d", j), styledDefAm) //% дефектуры от АМ

			j++
		}

		f.SetCellValue(defectsSheet, fmt.Sprintf("D%d", storeIndex), storeMatrixSales) // кол-во СКЮ продаваемых за 60 дней по матрице

		f.SetCellValue(defectsSheet, fmt.Sprintf("E%d", storeIndex), storeDefectQnt) //кол-во СКЮ в дефектуре

		f.SetCellValue(defectsSheet, fmt.Sprintf("F%d", storeIndex), storeDefectSum) //сумма дефектуры

		if storeMatrixSales == 0 {
			storeMatrixSales++
		}

		f.SetCellValue(defectsSheet, fmt.Sprintf("G%d", storeIndex), math.Abs(storeDefectQnt/storeMatrixSales*100)) //% дефектуры от факт продаж

		f.SetCellValue(defectsSheet, fmt.Sprintf("H%d", storeIndex), len(matrix)) //кол-во СКЮ входящих в АМ аптеки/магазина

		f.SetCellValue(defectsSheet, fmt.Sprintf("I%d", storeIndex), storeDefectSkuQnt) //кол-во СКЮ в дефектуре по АМ ПФ

		f.SetCellValue(defectsSheet, fmt.Sprintf("J%d", storeIndex), storeDefectSum2) //сумма дефектуры

		if storeSkuQnt == 0 {
			storeSkuQnt++
		}
		f.SetCellValue(defectsSheet, fmt.Sprintf("K%d", storeIndex), math.Abs(storeDefectSkuQnt/storeSkuQnt*100)) //% дефектуры от АМ
		//if storeDefectSkuQnt != 0 {
		//	f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), float64(storeDefectSkuQnt)/float64(storeSkuQnt)) //% дефектуры от АМ
		//} else {
		//	f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), 0) //% дефектуры от АМ
		//}
		f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", storeIndex), storeStoreSaldoQnt) // наличие продукции на складе
		//stream.SetRow(fmt.Sprintf("L%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeStoreSaldoQnt)}})

		f.SetCellValue(defectsSheet, fmt.Sprintf("N%d", storeIndex), storeStoreSaldoCount) // наличие продукции на складе - в суммарном выражении
		//stream.SetRow(fmt.Sprintf("M%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeStoreSaldoCount)}})

		//f.SetCellStyle(defectsSheet, fmt.Sprintf("A%d", storeIndex), fmt.Sprintf("M%d", storeIndex), style)
		f.SetCellStyle(defectsSheet, fmt.Sprintf("A%d", storeIndex), fmt.Sprintf("K%d", storeIndex), style)
		f.SetCellStyle(defectsSheet, fmt.Sprintf("M%d", storeIndex), fmt.Sprintf("S%d", storeIndex), style)
		globalDefectSum1 += storeDefectQnt
		globalDefectSum2 += storeDefectSkuQnt
		//if restriction > 1 {
		//	break
		//}
		//restriction++
	}

	f.SetCellValue(defectsSheet, "E3", globalDefectSum1)
	f.SetCellStyle(defectsSheet, "E3", "E3", moneyMainStyle)
	f.SetCellValue(defectsSheet, "I3", globalDefectSum2)
	f.SetCellStyle(defectsSheet, "I3", "I3", moneyMainStyle)

	f.DeleteSheet("Sheet1")

	fileName = fmt.Sprintf("files/defects/%d.xlsx", id)
	//_ = f.SaveAs("files/defects/res.xlsx")
	_ = f.SaveAs(fileName)
	return fileName, nil
}

func GetFormedDefectsList() (orders []models.DefectOrder, err error) {
	return repository.GetAllFormedDefects()
}

func GetFormedDefectByID(id int) (order models.DefectOrder, err error) {
	return repository.GetFormedDefectByID(id)
}
