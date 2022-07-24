package repository

import (
	"admin_panel/db"
	"admin_panel/models"
	"fmt"
)

func CreateGraphic(graphic models.Graphic) error {
	if err := db.GetDBConn().Table("graphics").Create(&graphic).Error; err != nil {
		return err
	}

	return nil
}

func GetAllGraphics() (graphics []models.Graphic, err error) {
	sqlQuery := `SELECT id,
					   number,
					   author,
					   supplier_name,
					   supplier_code,
					   region_name,
					   region_code,
					   store_name,
					   store_code,
					   nomenclature_group,
					   execution_period,
					   once_a_month,
					   twice_a_month,
					   is_on,
					   to_char(auto_order_date::date, 'DD.MM.YYYY'),
					   created_at,
					   application_day
				from graphics WHERE is_removed = false`
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&graphics).Error; err != nil {
		return nil, err
	}

	return graphics, nil
}

func GetGraphicByID(id int) (graphic models.Graphic, err error) {
	sqlQuery := `SELECT id,
					   number,
					   author,
					   supplier_name,
					   supplier_code,
					   region_name,
					   region_code,
					   store_name,
					   store_code,
					   nomenclature_group,
					   execution_period,
					   once_a_month,
					   twice_a_month,
					   is_on,
					   to_char(auto_order_date::date, 'DD.MM.YYYY'),
					   created_at,
					   application_day
				from graphics WHERE id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, id).Scan(&graphic).Error; err != nil {
		return models.Graphic{}, err
	}

	return graphic, nil
}

func EditGraphic(graphic models.Graphic) error {
	if err := db.GetDBConn().Table("graphics").Save(&graphic).Error; err != nil {
		return err
	}

	return nil
}

func DeleteGraphic(id int) error {
	fmt.Println("ID", id)
	// change auto_order to graphic
	//update := "UPDATE auto_order SET is_removed = ? WHERE id = ?"
	update := "UPDATE graphics SET is_removed = ? WHERE id = ?"
	err := db.GetDBConn().Exec(update, true, id).Error
	if err != nil {
		return err
	}

	return nil

}
