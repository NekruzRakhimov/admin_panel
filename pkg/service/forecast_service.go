package service

import (
	"admin_panel/models"
	"admin_panel/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

	return &models.Forecast{HistorySales: history}, err
}

func (s *forecastService) getHistoricalSales(params *models.ForecastSearchParameters) (*models.HistoricalSales, error) {
	body, err := json.Marshal(struct {
		DateStart string   `json:"datestart"`
		DateEnd   string   `json:"dateend"`
		Type      string   `json:"type"`
		Sku       []string `json:"sku"`
		Store     []string `json:"store"`
	}{
		DateStart: "01.01.2021 0:00:00",
		DateEnd:   "31.01.2021 23:59:59",
		Type:      "sales_by_day",
		Sku:       []string{*params.ProductCode},
		Store:     []string{*params.PharmacyCode},
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, utils.AppSettings.Route1c.Host+utils.AppSettings.Route1c.Routes.GetData,
		bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(utils.AppSettings.Route1c.Login, utils.AppSettings.Route1c.Password)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Статус ошибки при запросе продаж %s", resp.Status))
	}

	rawResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := models.HistoricalSales{}
	err = json.Unmarshal(rawResp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
