package service

import (
	"admin_panel/models"
	"encoding/base64"
	"github.com/xuri/excelize/v2"
)

const (
	hyperstocksPath             = "./files/hyperstocks/"
	hyperstocksPharmacyFileName = "hyperstocks_pharmacy.xlsx"
	hyperstocksStockFileName    = "hyperstocks_stock.xlsx"
)

type HyperstocksService interface {
	GetHyperstocksPharmacy(params *models.HyperstocksSearchParameters) (*models.HyperstocksFile, error)
	GetHyperstocksStock(params *models.HyperstocksSearchParameters) (*models.HyperstocksFile, error)
}

type hyperstocksService struct {
}

func NewHyperstocksService() HyperstocksService {
	return &hyperstocksService{}
}

func (s *hyperstocksService) GetHyperstocksPharmacy(params *models.HyperstocksSearchParameters) (*models.HyperstocksFile, error) {
	f, err := excelize.OpenFile(hyperstocksPath + hyperstocksPharmacyFileName)
	if err != nil {
		return nil, err
	}

	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return &models.HyperstocksFile{
		File:     base64.StdEncoding.EncodeToString(buffer.Bytes()),
		FileName: hyperstocksPharmacyFileName,
	}, nil
}

func (s *hyperstocksService) GetHyperstocksStock(params *models.HyperstocksSearchParameters) (*models.HyperstocksFile, error) {
	f, err := excelize.OpenFile(hyperstocksPath + hyperstocksStockFileName)
	if err != nil {
		return nil, err
	}

	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return &models.HyperstocksFile{
		File:     base64.StdEncoding.EncodeToString(buffer.Bytes()),
		FileName: hyperstocksStockFileName,
	}, nil
}
