package repository

import (
	"admin_panel/db"
	"admin_panel/models"
)

func CreateGraphic(graphic models.Graphic) error {
	if err := db.GetDBConn().Table("graphics").Create(&graphic).Error; err != nil {
		return err
	}

	return nil
}

func GetAllGraphics() (graphics []models.Graphic, err error) {
	sqlQuery := "SELECT * FROM graphics"
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&graphics).Error; err != nil {
		return nil, err
	}

	return graphics, nil
}

func GetGraphicByID(id int) (graphic models.Graphic, err error) {
	sqlQuery := "SELECT * FROM graphics WHERE id = ?"
	if err = db.GetDBConn().Raw(sqlQuery, id).Scan(&graphic).Error; err != nil {
		return models.Graphic{}, err
	}

	return graphic, nil
}

func EditGraphic(graphic models.Graphic) error {
	if err := db.GetDBConn().Table("graphics").Create(&graphic).Error; err != nil {
		return err
	}

	return nil
}
