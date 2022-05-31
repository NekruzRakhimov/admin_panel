package service

import (
	"admin_panel/models"
	"admin_panel/utils"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type ForecastService interface {
	GetForecast(params *models.ForecastSearchParameters) (*models.Forecast, error)
}

type forecastService struct {
}

func NewForecastService() ForecastService {
	return &forecastService{}
}

func (s *forecastService) GetForecast(params *models.ForecastSearchParameters) (*models.Forecast, error) {
	history, err := s.getHistoricalSales(params)
	if err != nil {
		return nil, err
	}
	for i, _ := range history.SalesArr {
		history.SalesArr[i].Category = "Исторические продажи"
	}

	forecast := &models.ForecastSales{}
	dateNow := time.Now()
	for i, f := range history.SalesArr {
		if i > 1 && i < 31 {

			fSale := models.Sale{
				QntTotal: f.QntTotal,
				Date:     dateNow.AddDate(0, 0, i).Format("2006-01-02T15:04:05"),
				Category: "Прогноз",
			}

			forecast.SalesArr = append(forecast.SalesArr, fSale)
		}
	}

	/*csvHistoricalFile, err := s.getCSVFile(*history) //TODO открыть когда будет готово сервис на тестовом сервере
	if err != nil {
		return nil, err
	}

	forecast, err := s.getForecast(csvHistoricalFile)
	if err != nil {
		return nil, err
	}
	for i, _ := range forecast.SalesArr {
		forecast.SalesArr[i].Category = "Прогноз"
	}*/

	sales := make([]models.Sale, 0)
	sales = append(sales, history.SalesArr...)
	sales = append(sales, forecast.SalesArr...)

	return &models.Forecast{Sales: sales}, err
}

func (s *forecastService) getHistoricalSales(params *models.ForecastSearchParameters) (*models.HistoricalSales, error) {
	body, err := json.Marshal(struct {
		DateStart string   `json:"datestart"`
		DateEnd   string   `json:"dateend"`
		Type      string   `json:"type"`
		Sku       []string `json:"sku"`
		Store     []string `json:"store"`
	}{
		DateStart: "20.05.2021 0:00:00",
		DateEnd:   "20.05.2022 23:59:59",
		Type:      "sales_by_day",
		Sku:       []string{*params.ProductCode},
		Store:     []string{*params.PharmacyCode},
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка при формировнии рекуеста для 1с сервиса %s", err.Error()))
	}
	req, err := http.NewRequest(http.MethodPost, utils.AppSettings.Route1c.Host+utils.AppSettings.Route1c.Routes.GetData,
		bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка при формировнии сервиса 1с %s", err.Error()))
	}
	req.SetBasicAuth(utils.AppSettings.Route1c.Login, utils.AppSettings.Route1c.Password)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка при вызове сервиса 1с %s", err.Error()))
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Статус ошибки при запросе продаж %s", resp.Status))
	}

	rawResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка при обработке тело респонса 1с %s", err.Error()))
	}

	defer resp.Body.Close()

	rawResp = bytes.TrimPrefix(rawResp, []byte("\xef\xbb\xbf"))

	response := models.HistoricalSales{}
	err = json.Unmarshal(rawResp, &response)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("Ошибка при обработке респонса 1с %s", err.Error()))
	}

	return &response, nil
}

func (s *forecastService) getForecast(csv *os.File) (*models.ForecastSales, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", "forecast.csv")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, csv)
	if err != nil {
		return nil, err
	}
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, utils.AppSettings.ForecastUrl, bytes.NewReader(body.Bytes()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Request failed with response code: %d", rsp.StatusCode))
	}

	rawResp, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка при обработке тело forecast %s", err.Error()))
	}

	defer rsp.Body.Close()

	rawResp = bytes.TrimPrefix(rawResp, []byte("\xef\xbb\xbf"))

	response := models.ForecastSales{}
	err = json.Unmarshal(rawResp, &response)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("Ошибка при обработке респонса forecast %s", err.Error()))
	}

	return &response, nil
}

func (s *forecastService) getCSVFile(sales models.HistoricalSales) (*os.File, error) {
	// Create a csv file
	f, err := os.Create("forecast.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	for _, obj := range sales.SalesArr {
		var record []string
		record = append(record, obj.Date)
		record = append(record, fmt.Sprintf("%.2f", obj.QntTotal))
		w.Write(record)
	}
	w.Flush()

	return f, nil
}
