package service

import (
	"admin_panel/models"
	"encoding/base64"
	"github.com/xuri/excelize/v2"
)

const (
	defectsPath             = "./files/defects/"
	defectsPharmacyFileName = "defects_pharmacy.xlsx"
	defectsStockFileName    = "defects_stock.xlsx"
)

type DefectsService interface {
	GetDefectsPharmacy(params *models.DefectsSearchParameters) (*models.DefectsFile, error)
	GetDefectsStock(params *models.DefectsSearchParameters) (*models.DefectsFile, error)
}

type defectsService struct {
}

func NewDefectsService() DefectsService {
	return &defectsService{}
}

func (s *defectsService) GetDefectsPharmacy(params *models.DefectsSearchParameters) (*models.DefectsFile, error) {
	f, err := excelize.OpenFile(defectsPath + defectsPharmacyFileName)
	if err != nil {
		return nil, err
	}

	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return &models.DefectsFile{
		File:     base64.StdEncoding.EncodeToString(buffer.Bytes()),
		FileName: defectsPharmacyFileName,
	}, nil
}

func (s *defectsService) GetDefectsStock(params *models.DefectsSearchParameters) (*models.DefectsFile, error) {
	f, err := excelize.OpenFile(defectsPath + defectsStockFileName)
	if err != nil {
		return nil, err
	}

	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return &models.DefectsFile{
		File:     base64.StdEncoding.EncodeToString(buffer.Bytes()),
		FileName: defectsStockFileName,
	}, nil
}
