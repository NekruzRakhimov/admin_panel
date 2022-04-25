package service

import (
	"admin_panel/models"
	"github.com/xuri/excelize/v2"
)

const (
	DefectsPath             = "./files/defects/"
	DefectsPathData         = "./files/defects/data/"
	DefectsPharmacyFileName = "defects_pharmacy.xlsx"
	DefectsStockFileName    = "defects_stock.xlsx"
)

type DefectsService interface {
	GetDefectsPharmacy(params *models.DefectsSearchParameters) error
	GetDefectsStock(params *models.DefectsSearchParameters) error
}

type defectsService struct {
}

func NewDefectsService() DefectsService {
	return &defectsService{}
}

func (s *defectsService) GetDefectsPharmacy(params *models.DefectsSearchParameters) error {
	f, err := excelize.OpenFile(DefectsPath + DefectsPharmacyFileName)
	if err != nil {
		return err
	}

	if err = f.SaveAs(DefectsPathData + DefectsPharmacyFileName); err != nil {
		return err
	}

	return nil
}

func (s *defectsService) GetDefectsStock(params *models.DefectsSearchParameters) error {
	f, err := excelize.OpenFile(DefectsPath + DefectsStockFileName)
	if err != nil {
		return err
	}

	if err = f.SaveAs(DefectsPathData + DefectsStockFileName); err != nil {
		return err
	}

	return nil
}
