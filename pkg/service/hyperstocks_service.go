package service

import (
	"admin_panel/models"
	"github.com/xuri/excelize/v2"
)

const (
	HyperstocksPath             = "./files/hyperstocks/"
	HyperstocksPathData         = "./files/hyperstocks/data"
	HyperstocksPharmacyFileName = "hyperstocks_pharmacy.xlsx"
	HyperstocksStockFileName    = "hyperstocks_stock.xlsx"
)

type HyperstocksService interface {
	GetHyperstocksPharmacy(params *models.HyperstocksSearchParameters) error
	GetHyperstocksStock(params *models.HyperstocksSearchParameters) error
}

type hyperstocksService struct {
}

func NewHyperstocksService() HyperstocksService {
	return &hyperstocksService{}
}

func (s *hyperstocksService) GetHyperstocksPharmacy(params *models.HyperstocksSearchParameters) error {
	f, err := excelize.OpenFile(HyperstocksPath + HyperstocksPharmacyFileName)
	if err != nil {
		return err
	}

	if err = f.SaveAs(HyperstocksPathData + HyperstocksPharmacyFileName); err != nil {
		return err
	}

	return nil
}

func (s *hyperstocksService) GetHyperstocksStock(params *models.HyperstocksSearchParameters) error {
	f, err := excelize.OpenFile(HyperstocksPath + HyperstocksStockFileName)
	if err != nil {
		return err
	}

	if err = f.SaveAs(HyperstocksPathData + HyperstocksStockFileName); err != nil {
		return err
	}

	return nil
}
