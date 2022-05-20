package service

import (
	"admin_panel/models"
	"admin_panel/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	forecast := &models.ForecastSales{}
	dateNow := time.Now()
	for i, f := range history.SalesArr {
		if i > 1 && i < 31 {

			fSale := models.Sale{
				QntTotal: f.QntTotal,
				Date:     dateNow.AddDate(0, 0, i).Format("2006-01-02T15:04:05"),
			}

			forecast.SalesArr = append(forecast.SalesArr, fSale)
		}
	}

	return &models.Forecast{HistorySales: history, ForecastSales: forecast}, err
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
