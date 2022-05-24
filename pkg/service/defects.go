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
)

const defectsSheet = "TDSheet"

func GetDefectsExt(req models.DefectsRequest) (defects []models.Defect, err error) {
	//var binOrganizationAKNIET = "060540001442"

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
	defects, err := GetDefectsExt(req)
	if err != nil {
		return nil, err
	}
	var checkedCodes []string

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

	if err = FormExcelDefects(filteredDefects); err != nil {
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

func FormExcelDefects(filteredDefects []models.DefectsFiltered) error {
	f, err := excelize.OpenFile("files/defects/defects_pharmacy_template.xlsx")
	if err != nil {
		return err
	}

	//f.NewSheet(defectsSheet)
	var i = 2
	for _, defect := range filteredDefects {
		f.SetCellValue(defectsSheet, fmt.Sprintf("C%d", i), defect.StoreName) //Аптека
		i++
		for _, subDefect := range defect.SubDefects {
			storeSaldoQnt, err := strconv.Atoi(subDefect.StoreSaldoQnt)
			if err != nil {
				return err
			}
			defectQnt, err := strconv.Atoi(subDefect.DefectQnt)
			if err != nil {
				return err
			}

			matrixProductQnt, err := strconv.Atoi(subDefect.MatrixProductQnt)
			if err != nil {
				return err
			}

			price, err := strconv.ParseFloat(subDefect.DefectPrice, 2)
			if err != nil {
				return err
			}

			matrixSales, err := strconv.ParseFloat(subDefect.MatrixSales, 2)
			if err != nil {
				return err
			}

			f.SetCellValue(defectsSheet, fmt.Sprintf("D%d", i), subDefect.ProductName) //Наименование
			f.SetCellValue(defectsSheet, fmt.Sprintf("E%d", i), subDefect.ProductCode) //код 1С
			f.SetCellValue(defectsSheet, fmt.Sprintf("F%d", i), subDefect.MatrixSales) // кол-во СКЮ продаваемых за 60 дней по матрице
			f.SetCellValue(defectsSheet, fmt.Sprintf("G%d", i), subDefect.DefectQnt)   //кол-во СКЮ в дефектуре

			f.SetCellValue(defectsSheet, fmt.Sprintf("H%d", i), float64(defectQnt)*price)              //сумма дефектуры
			f.SetCellValue(defectsSheet, fmt.Sprintf("I%d", i), float64(defectQnt)/matrixSales*100)    //% дефектуры от факт продаж
			f.SetCellValue(defectsSheet, fmt.Sprintf("J%d", i), subDefect.MatrixProductQnt)            //кол-во СКЮ входящих в АМ аптеки/магазина
			f.SetCellValue(defectsSheet, fmt.Sprintf("K%d", i), len(defect.SubDefects))                //кол-во СКЮ в дефектуре по АМ ПФ
			f.SetCellValue(defectsSheet, fmt.Sprintf("L%d", i), float64(len(defect.SubDefects))*price) //сумма дефектуры

			f.SetCellValue(defectsSheet, fmt.Sprintf("M%d", i), (float64(len(defect.SubDefects))*price)*float64(matrixProductQnt)) //% дефектуры от АМ
			f.SetCellValue(defectsSheet, fmt.Sprintf("N%d", i), storeSaldoQnt)                                                     // наличие продукции на складе
			f.SetCellValue(defectsSheet, fmt.Sprintf("O%d", i), float64(storeSaldoQnt)*price)                                      // наличие продукции на складе - в суммарном выражении
			f.SetCellValue(defectsSheet, fmt.Sprintf("P%d", i), price)                                                             // Закупочная цена

			i++
		}
	}

	f.DeleteSheet("Sheet1")
	f.SaveAs("files/defects/res.xlsx")
	//f.Close()
	return nil
}

func GetSalesCountExt(req models.SalesCountRequest) (defects []models.SalesCount, err error) {
	//var binOrganizationAKNIET = "060540001442"

	response := struct {
		SalesCountArr []models.SalesCount `json:"sales_count_arr"`
	}{}

	bodyBin := new(bytes.Buffer)
	err = json.NewEncoder(bodyBin).Encode(&req)
	if err != nil {
		return nil, err
	}
	fmt.Println("BODY", bodyBin)

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

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	// ----------> часть Unmarshall json ->
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	fmt.Println("BODY", string(body))
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.SalesCountArr, nil
}
