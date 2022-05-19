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
	var i = 4
	for _, defect := range filteredDefects {
		f.SetCellValue(defectsSheet, fmt.Sprintf("C%d", i), defect.StoreName)
		i++
		for _, subDefect := range defect.SubDefects {
			f.SetCellValue(defectsSheet, fmt.Sprintf("D%d", i), subDefect.ProductName)
			f.SetCellValue(defectsSheet, fmt.Sprintf("E%d", i), subDefect.ProductCode)
			f.SetCellValue(defectsSheet, fmt.Sprintf("F%d", i), subDefect.MatrixTotalSales)
			f.SetCellValue(defectsSheet, fmt.Sprintf("G%d", i), subDefect.DefectQnt)
			f.SetCellValue(defectsSheet, fmt.Sprintf("H%d", i), subDefect.StoreSaldoTotal)
			f.SetCellValue(defectsSheet, fmt.Sprintf("I%d", i), subDefect.StoreSaldoQnt)
			f.SetCellValue(defectsSheet, fmt.Sprintf("J%d", i), subDefect.DefectTotalQnt)
			f.SetCellValue(defectsSheet, fmt.Sprintf("K%d", i), subDefect.DefectTotal)
			f.SetCellValue(defectsSheet, fmt.Sprintf("L%d", i), subDefect.DifPercent)
			i++
		}
	}

	f.DeleteSheet("Sheet1")
	f.SaveAs("files/defects/res.xlsx")
	//f.Close()
	return nil
}
