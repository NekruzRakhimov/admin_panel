package service

import (
	"admin_panel/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
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

func GetDefectsPF(req models.DefectsRequest) (filteredDefects []models.DefectsFiltered, err error) {
	//var filteredDefects []models.DefectsFiltered
	req.IsPF = true
	log.Println(time.Now(), " Started Getting Defects from 1C")
	fmt.Println(time.Now(), " Started Getting Defects from 1C")
	now := time.Now()
	req.DaysCount = -60
	req.QueryType = "warehouse_defect"

	defects, err := GetDefectsExt(req)
	if err != nil {
		return nil, err
	}
	log.Println(time.Now(), " Finished Getting Defects from 1C: durance[", time.Now().Sub(now), "]")
	fmt.Println(time.Now(), " Finished Getting Defects from 1C: durance[", time.Now().Sub(now), "]")
	var checkedCodes []string

	log.Println(time.Now(), " Started Filtering data into Stores")
	fmt.Println(time.Now(), " Started Filtering data into Stores")
	now = time.Now()
	for i, defect := range defects {
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
	log.Println(time.Now(), "Finished Filtering data into Stores: durance[", time.Now().Sub(now), "]")
	fmt.Println(time.Now(), "Finished Filtering data into Stores: durance[", time.Now().Sub(now), "]")

	log.Println(time.Now(), "Start Forming excel")
	fmt.Println(time.Now(), "Start Forming excel")
	now = time.Now()
	if err = FormExcelDefectsPF(req, filteredDefects); err != nil {
		return nil, err
	}
	log.Println(time.Now(), "Finished Forming excel: durance[", time.Now().Sub(now), "]")
	fmt.Println(time.Now(), "Finished Forming excel")

	return filteredDefects, nil
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

func FormExcelDefectsPF(req models.DefectsRequest, filteredDefects []models.DefectsFiltered) error {
	f, err := excelize.OpenFile("files/defects/defects_pharmacy_template.xlsx")
	if err != nil {
		return err
	}

	//stream, _ := f.NewStreamWriter(defectsSheet)

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

	f.NewSheet(defectsSheet)
	f.SetCellValue(defectsSheet, "D1", req.Startdate) //Дата
	f.SetCellValue(defectsSheet, "H1", req.Startdate) //Дата

	//stream.SetRow("D1", []interface{}{excelize.Cell{Value: req.Startdate}})
	//stream.SetRow("H1", []interface{}{excelize.Cell{Value: req.Startdate}})

	//stream.SetRow("A1", []interface{}{excelize.Cell{Value: "ПФ"}})

	//stream.SetRow("A2", []interface{}{excelize.Cell{Value: "Аптека"}})
	//stream.SetRow("B2", []interface{}{excelize.Cell{Value: "Наименование"}})
	//stream.SetRow("C2", []interface{}{excelize.Cell{Value: "код 1С"}})
	//stream.SetRow("D2", []interface{}{excelize.Cell{Value: "кол-во СКЮ продаваемых за 60 дней по матрице"}})
	//stream.SetRow("E2", []interface{}{excelize.Cell{Value: "кол-во СКЮ в дефектуре"}})
	//stream.SetRow("F2", []interface{}{excelize.Cell{Value: "сумма дефектуры"}})
	//stream.SetRow("G2", []interface{}{excelize.Cell{Value: "% дефектуры от факт продаж"}})
	//stream.SetRow("H2", []interface{}{excelize.Cell{Value: "кол-во СКЮ входящих в АМ аптеки/магазина"}})
	//stream.SetRow("I2", []interface{}{excelize.Cell{Value: "кол-во СКЮ в дефектуре по АМ ПФ"}})
	//stream.SetRow("J2", []interface{}{excelize.Cell{Value: "сумма дефектуры"}})
	//stream.SetRow("K2", []interface{}{excelize.Cell{Value: "% дефектуры от АМ"}})
	//stream.SetRow("L2", []interface{}{excelize.Cell{Value: "наличие продукции на складе"}})
	//stream.SetRow("M2", []interface{}{excelize.Cell{Value: "наличие продукции на складе - в суммарном выражении"}})
	//stream.SetRow("N2", []interface{}{excelize.Cell{Value: "закупочная цена"}})
	//
	//stream.SetRow("A3", []interface{}{excelize.Cell{Value: "Дивизион/филиал"}})

	var (
		globalDefectSum1 float64
		globalDefectSum2 float64
	)
	var i = 4
	for _, defect := range filteredDefects {
		storeIndex := i
		f.SetCellValue(defectsSheet, fmt.Sprintf("A%d", storeIndex), defect.StoreName) //Аптека
		//stream.SetRow(fmt.Sprintf("A%d", storeIndex), []interface{}{excelize.Cell{Value: defect.StoreName}}) //Аптека
		var (
			storeMatrixSales     float64 // кол-во СКЮ продаваемых за 60 дней по матрице
			storeDefectQnt       float64 //кол-во СКЮ в дефектуре
			storeDefectSum       float64 //сумма дефектуры
			storeFactSaleDefect  float64 //% дефектуры от факт продаж
			storeSkuQnt          float64 //кол-во СКЮ входящих в АМ аптеки/магазина
			storeDefectSkuQnt    float64 //кол-во СКЮ в дефектуре по АМ ПФ
			storeDefectSum2      float64 //сумма дефектуры
			storeDefectAM        float64 //% дефектуры от АМ
			storeStoreSaldoQnt   float64 // наличие продукции на складе
			storeStoreSaldoCount float64 // наличие продукции на складе - в суммарном выражении
		)
		i++
		for _, subDefect := range defect.SubDefects {
			f.SetCellValue(defectsSheet, fmt.Sprintf("A%d", i), defect.StoreName)
			storeSaldoQnt, _ := strconv.ParseFloat(subDefect.StoreSaldoQnt, 2)
			//if err != nil {
			//	return err
			//}

			defectQnt, _ := strconv.ParseFloat(subDefect.DefectQnt, 2)
			//if err != nil {
			//	return err
			//}

			matrixProductQnt, _ := strconv.ParseFloat(subDefect.MatrixProductQnt, 2)
			//if err != nil {
			//	return err
			//}

			price, _ := strconv.ParseFloat(subDefect.DefectPrice, 2)
			//if err != nil {
			//	return err
			//}

			matrixSales, _ := strconv.ParseFloat(subDefect.MatrixSales, 2)
			//if err != nil {
			//	return err
			//}

			f.SetCellValue(defectsSheet, fmt.Sprintf("B%d", i), subDefect.ProductName) //Наименование
			//stream.SetRow(fmt.Sprintf("B%d", i), []interface{}{excelize.Cell{Value: subDefect.ProductName}})

			f.SetCellValue(defectsSheet, fmt.Sprintf("C%d", i), subDefect.ProductCode) //код 1С
			//stream.SetRow(fmt.Sprintf("C%d", i), []interface{}{excelize.Cell{Value: subDefect.ProductCode}})

			f.SetCellValue(defectsSheet, fmt.Sprintf("D%d", i), matrixSales) // кол-во СКЮ продаваемых за 60 дней по матрице
			//stream.SetRow(fmt.Sprintf("D%d", i), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(matrixSales)}})
			storeMatrixSales += matrixSales

			f.SetCellValue(defectsSheet, fmt.Sprintf("E%d", i), defectQnt) //кол-во СКЮ в дефектуре
			//stream.SetRow(fmt.Sprintf("E%d", i), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(defectQnt)}})
			storeDefectQnt += defectQnt

			f.SetCellValue(defectsSheet, fmt.Sprintf("F%d", i), defectQnt*price) //сумма дефектуры
			//stream.SetRow(fmt.Sprintf("F%d", i), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(defectQnt * price)}})
			storeDefectSum += defectQnt * price

			if matrixSales == 0 {
				f.SetCellValue(defectsSheet, fmt.Sprintf("G%d", i), 0) //% дефектуры от факт продаж
				//stream.SetRow(fmt.Sprintf("G%d", i), []interface{}{excelize.Cell{Value: 0}})
			} else {
				f.SetCellValue(defectsSheet, fmt.Sprintf("G%d", i), defectQnt/matrixSales*100) //% дефектуры от факт продаж
				//stream.SetRow(fmt.Sprintf("G%d", i), []interface{}{excelize.Cell{Value: fmt.Sprintf("%s%", utils.FloatToMoneyFormat(defectQnt/matrixSales*100))}})
			}
			storeFactSaleDefect += defectQnt / matrixSales * 100

			f.SetCellValue(defectsSheet, fmt.Sprintf("H%d", i), matrixProductQnt) //кол-во СКЮ входящих в АМ аптеки/магазина
			//stream.SetRow(fmt.Sprintf("H%d", i), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(matrixProductQnt)}})
			storeSkuQnt += matrixProductQnt

			f.SetCellValue(defectsSheet, fmt.Sprintf("I%d", i), len(defect.SubDefects)) //кол-во СКЮ в дефектуре по АМ ПФ
			//stream.SetRow(fmt.Sprintf("I%d", i), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(float64(len(defect.SubDefects)))}})
			storeDefectSkuQnt = float64(len(defect.SubDefects))

			f.SetCellValue(defectsSheet, fmt.Sprintf("J%d", i), float64(len(defect.SubDefects))*price) //сумма дефектуры
			//stream.SetRow(fmt.Sprintf("J%d", i), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(float64(len(defect.SubDefects)) * price)}})
			storeDefectSum2 += float64(len(defect.SubDefects)) * price

			f.SetCellValue(defectsSheet, fmt.Sprintf("K%d", i), float64(len(defect.SubDefects))/matrixProductQnt*100) //% дефектуры от АМ
			//stream.SetRow(fmt.Sprintf("K%d", i), []interface{}{excelize.Cell{Value: fmt.Sprintf("%s%", utils.FloatToMoneyFormat((float64(len(defect.SubDefects)))/matrixProductQnt*100))}})
			storeDefectAM += float64(len(defect.SubDefects)) / matrixProductQnt * 100

			f.SetCellValue(defectsSheet, fmt.Sprintf("L%d", i), storeSaldoQnt) // наличие продукции на складе
			//stream.SetRow(fmt.Sprintf("L%d", i), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeSaldoQnt)}})
			storeStoreSaldoQnt += storeSaldoQnt

			f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), storeSaldoQnt*price) // наличие продукции на складе - в суммарном выражении
			//stream.SetRow(fmt.Sprintf("M%d", i), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeSaldoQnt * price)}})
			storeStoreSaldoCount += storeSaldoQnt * price

			f.SetCellValue(defectsSheet, fmt.Sprintf("N%d", i), price) // Закупочная цена
			//stream.SetRow(fmt.Sprintf("N%d", i), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(price)}})
			f.SetCellStyle(defectsSheet, fmt.Sprintf("D%d", i), fmt.Sprintf("N%d", i), moneyStyle)
			i++
		}

		f.SetCellValue(defectsSheet, fmt.Sprintf("D%d", storeIndex), storeMatrixSales) // кол-во СКЮ продаваемых за 60 дней по матрице
		//stream.SetRow(fmt.Sprintf("D%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeMatrixSales)}})

		f.SetCellValue(defectsSheet, fmt.Sprintf("E%d", storeIndex), storeDefectQnt) //кол-во СКЮ в дефектуре
		//stream.SetRow(fmt.Sprintf("E%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeDefectQnt)}})

		f.SetCellValue(defectsSheet, fmt.Sprintf("F%d", storeIndex), storeDefectSum) //сумма дефектуры
		//stream.SetRow(fmt.Sprintf("F%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeDefectSum)}})

		f.SetCellValue(defectsSheet, fmt.Sprintf("G%d", storeIndex), storeDefectQnt/storeMatrixSales*100) //% дефектуры от факт продаж
		//stream.SetRow(fmt.Sprintf("G%d", storeIndex), []interface{}{excelize.Cell{Value: fmt.Sprintf("%s%", utils.FloatToMoneyFormat(storeDefectQnt/storeMatrixSales))}})

		f.SetCellValue(defectsSheet, fmt.Sprintf("H%d", storeIndex), storeSkuQnt) //кол-во СКЮ входящих в АМ аптеки/магазина
		//stream.SetRow(fmt.Sprintf("H%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeSkuQnt)}})

		f.SetCellValue(defectsSheet, fmt.Sprintf("I%d", storeIndex), storeDefectSkuQnt) //кол-во СКЮ в дефектуре по АМ ПФ
		//stream.SetRow(fmt.Sprintf("I%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeDefectSkuQnt)}})

		f.SetCellValue(defectsSheet, fmt.Sprintf("J%d", storeIndex), storeDefectSum2) //сумма дефектуры
		//stream.SetRow(fmt.Sprintf("J%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeDefectSum2)}})

		f.SetCellValue(defectsSheet, fmt.Sprintf("K%d", storeIndex), storeDefectSkuQnt/storeSkuQnt*100) //% дефектуры от АМ
		//stream.SetRow(fmt.Sprintf("K%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeDefectAM)}})
		//if storeDefectSkuQnt != 0 {
		//	f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), float64(storeDefectSkuQnt)/float64(storeSkuQnt)) //% дефектуры от АМ
		//} else {
		//	f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), 0) //% дефектуры от АМ
		//}
		f.SetCellValue(defectsSheet, fmt.Sprintf("L%d", storeIndex), storeStoreSaldoQnt) // наличие продукции на складе
		//stream.SetRow(fmt.Sprintf("L%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeStoreSaldoQnt)}})

		f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", storeIndex), storeStoreSaldoCount) // наличие продукции на складе - в суммарном выражении
		//stream.SetRow(fmt.Sprintf("M%d", storeIndex), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(storeStoreSaldoCount)}})

		//f.SetCellStyle(defectsSheet, fmt.Sprintf("A%d", storeIndex), fmt.Sprintf("M%d", storeIndex), style)
		f.SetCellStyle(defectsSheet, fmt.Sprintf("A%d", storeIndex), fmt.Sprintf("M%d", storeIndex), style)
		globalDefectSum1 += storeDefectSum
		globalDefectSum2 += storeDefectSum2
	}

	//f.SetCellValue(defectsSheet, "F3", globalDefectSum1)
	//stream.SetRow(fmt.Sprintf("F3"), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(globalDefectSum1)}})

	//f.SetCellValue(defectsSheet, "J3", globalDefectSum2)
	//stream.SetRow(fmt.Sprintf("J3"), []interface{}{excelize.Cell{Value: utils.FloatToMoneyFormat(globalDefectSum2)}})

	f.SetCellValue(defectsSheet, "F3", globalDefectSum1)
	f.SetCellStyle(defectsSheet, "F3", "F3", moneyMainStyle)
	f.SetCellValue(defectsSheet, "J3", globalDefectSum2)
	f.SetCellStyle(defectsSheet, "J3", "J3", moneyMainStyle)

	f.DeleteSheet("Sheet1")
	//if err = stream.Flush(); err != nil {
	//	return err
	//}

	_ = f.SaveAs("files/defects/res.xlsx")
	//f.Close()
	return nil
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
	bodyBin := new(bytes.Buffer)
	err = json.NewEncoder(bodyBin).Encode(&req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Finished marshalling ext_req_body")
	//fmt.Println("BODY", bodyBin)

	fmt.Println("Started sending ext_req")
	client := &http.Client{}
	endpoint := fmt.Sprintf("http://89.218.153.38:8081/AQG_ULAN/hs/integration/salescount")
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
