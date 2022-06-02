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
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ForecastService interface {
	GetForecast(params *models.ForecastSearchParameters) (*models.Forecast, error)
}

type forecastService struct {
}

func NewForecastService() ForecastService {
	return &forecastService{}
}

const forecastFile = "forecast.csv"

func (s *forecastService) GetForecast(params *models.ForecastSearchParameters) (*models.Forecast, error) {
	history, err := s.getHistoricalSales(params)
	if err != nil {
		return nil, err
	}
	for i, _ := range history.SalesArr {
		history.SalesArr[i].Category = "Исторические продажи"
	}

	forecast, err := s.getForecast(*history)
	if err != nil {
		return nil, err
	}
	for i, _ := range forecast.SalesArr {
		forecast.SalesArr[i].Category = "Прогноз"
	}

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

func (s *forecastService) getForecast(sales models.HistoricalSales) (*models.ForecastSales, error) {
	err := generateCSVFile(sales)
	if err != nil {
		return nil, err
	}

	prophet, err := postProphet(utils.AppSettings.ForecastUrl, forecastFile)
	if err != nil {
		return nil, err
	}

	forecast := &models.ForecastSales{}
	for _, f := range prophet.Data {
		fSale := models.Sale{
			QntTotal: f.XGBoost,
			Date:     strings.ReplaceAll(f.Ds, ".000Z", ""),
			Category: "Прогноз",
		}

		forecast.SalesArr = append(forecast.SalesArr, fSale)

	}

	return forecast, nil
}

func generateCSVFile(sales models.HistoricalSales) error {
	// Create a csv file
	f, err := os.Create(forecastFile)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	titleCSV := []string{"dteday", "cnt", "holiday", "workingday"}
	w.Write(titleCSV)
	for i, obj := range sales.SalesArr {
		var record []string
		record = append(record, obj.Date)
		record = append(record, fmt.Sprintf("%.2f", obj.QntTotal))
		if i == 5 {
			record = append(record, "1")
		} else {
			record = append(record, "0")
		}
		record = append(record, "1")
		w.Write(record)
	}
	w.Flush()

	return nil
}

func postProphet(dst, fname string) (*models.ProphetSales, error) {
	u, err := url.Parse(dst)
	if err != nil {
		return nil, fmt.Errorf("failed to parse destination url: %w", err)
	}

	form, err := makeRequestBody(fname)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request body: %w", err)
	}

	hdr := make(http.Header)
	hdr.Set("Content-Type", form.contentType)
	req := http.Request{
		Method:        http.MethodPost,
		URL:           u,
		Header:        hdr,
		Body:          ioutil.NopCloser(form.body),
		ContentLength: int64(form.contentLen),
	}

	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform http request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Request failed with response code in python service: %d", resp.StatusCode))
	}

	rawResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка при обработке тело forecast %s", err.Error()))
	}

	defer resp.Body.Close()

	rawResp = bytes.TrimPrefix(rawResp, []byte("\xef\xbb\xbf"))

	response := models.ProphetSales{}
	err = json.Unmarshal(rawResp, &response)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("Ошибка при обработке респонса forecast %s", err.Error()))
	}

	return &response, nil
}

type form struct {
	body        *bytes.Buffer
	contentType string
	contentLen  int
}

func makeRequestBody(fname string) (form, error) {
	ct, err := getFileContentType(fname)
	if err != nil {
		return form{}, fmt.Errorf(
			`failed to get content type for file "%s": %w`,
			fname, err)
	}

	fd, err := os.Open(fname)
	if err != nil {
		return form{}, fmt.Errorf("failed to open file to upload: %w", err)
	}
	defer fd.Close()

	stat, err := fd.Stat()
	if err != nil {
		return form{}, fmt.Errorf("failed to query file info: %w", err)
	}

	hdr := make(textproto.MIMEHeader)
	cd := mime.FormatMediaType("form-data", map[string]string{
		"name":     "file",
		"filename": fname,
	})
	hdr.Set("Content-Disposition", cd)
	hdr.Set("Content-Type", ct)
	hdr.Set("Content-Length", strconv.FormatInt(stat.Size(), 10))

	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)

	part, err := mw.CreatePart(hdr)
	if err != nil {
		return form{}, fmt.Errorf("failed to create new form part: %w", err)
	}

	n, err := io.Copy(part, fd)
	if err != nil {
		return form{}, fmt.Errorf("failed to write form part: %w", err)
	}

	if int64(n) != stat.Size() {
		return form{}, fmt.Errorf("file size changed while writing: %s", fd.Name())
	}

	err = mw.Close()
	if err != nil {
		return form{}, fmt.Errorf("failed to prepare form: %w", err)
	}

	return form{
		body:        &buf,
		contentType: mw.FormDataContentType(),
		contentLen:  buf.Len(),
	}, nil
}

var fileContentTypes = map[string]string{
	"csv": "text/csv",
}

func getFileContentType(fname string) (string, error) {
	ext := filepath.Ext(fname)
	if ext == "" {
		return "", fmt.Errorf("file name has no extension: %s", fname)
	}

	ext = strings.ToLower(ext[1:])
	ct, found := fileContentTypes[ext]
	if !found {
		return "", fmt.Errorf("unknown file name extension: %s", ext)
	}

	return ct, nil
}
